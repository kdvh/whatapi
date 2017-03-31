//Package whatapi is a wrapper for the What.CD JSON API (https://github.com/WhatCD/Gazelle/wiki/JSON-API-Documentation).
package whatapi

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strconv"
	"strings"
)

//NewWhatAPI creates a new client for the What.CD API using the provided URL.
func NewWhatAPI(url, agent string) (*WhatAPI, error) {
	w := new(WhatAPI)
	w.baseURL = url
	w.userAgent = agent
	cookieJar, err := cookiejar.New(nil)
	if err != nil {
		return w, err
	}
	w.client = &http.Client{Jar: cookieJar}
	return w, err
}

//WhatAPI represents a client for the What.CD API.
type WhatAPI struct {
	baseURL   string
	userAgent string
	client    *http.Client
	authkey   string
	passkey   string
	loggedIn  bool
}

//GetJSON sends a HTTP GET request to the API and decodes the JSON response into responseObj.
func (w *WhatAPI) GetJSON(requestURL string, responseObj interface{}) error {
	if !w.loggedIn {
		return errRequestFailedLogin
	}

	req, err := http.NewRequest("GET", requestURL, nil)
	req.Header.Set("User-Agent", w.userAgent)
	if err != nil {
		return err
	}
	resp, err := w.client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return errRequestFailedReason("Status Code " + resp.Status)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var st GenericResponse
	if err := json.Unmarshal(body, &st); err != nil {
		return err
	}

	if err := checkResponseStatus(st.Status, st.Error); err != nil {
		return err
	}
	return json.Unmarshal(body, responseObj)
}

type GenericResponse struct {
	Status string `json:"status"`
	Error  string `json:"error"`
}

func (w *WhatAPI) Do(action string, params url.Values, result interface{}) error {
	requestURL, err := buildURL(w.baseURL, "ajax.php", action, params)
	if err != nil {
		return err
	}
	return w.GetJSON(requestURL, result)
}

//CreateDownloadURL constructs a download URL using the provided torrent id.
func (w *WhatAPI) CreateDownloadURL(id int) (string, error) {
	if !w.loggedIn {
		return "", errRequestFailedLogin
	}

	params := url.Values{}
	params.Set("action", "download")
	params.Set("id", strconv.Itoa(id))
	params.Set("authkey", w.authkey)
	params.Set("torrent_pass", w.passkey)
	downloadURL, err := buildURL(w.baseURL, "torrents.php", "", params)
	if err != nil {
		return "", err
	}
	return downloadURL, nil
}

//Login logs in to the API using the provided credentials.
func (w *WhatAPI) Login(username, password string) error {
	params := url.Values{}
	params.Set("username", username)
	params.Set("password", password)

	reqBody := strings.NewReader(params.Encode())
	req, err := http.NewRequest("POST", w.baseURL+"login.php", reqBody)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("User-Agent", w.userAgent)
	resp, err := w.client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	if resp.Request.URL.String()[len(w.baseURL):] != "index.php" {
		return errLoginFailed
	}
	w.loggedIn = true
	account, err := w.GetAccount()
	if err != nil {
		return err
	}
	w.authkey, w.passkey = account.AuthKey, account.PassKey
	return nil
}

//Logout logs out of the API, ending the current session.
func (w *WhatAPI) Logout() error {
	params := url.Values{"auth": {w.authkey}}
	requestURL, err := buildURL(w.baseURL, "logout.php", "", params)
	if err != nil {
		return err
	}
	_, err = w.client.Get(requestURL)
	if err != nil {
		return err
	}
	w.loggedIn, w.authkey, w.passkey = false, "", ""
	return nil
}

//GetAccount retrieves account information for the current user.
func (w *WhatAPI) GetAccount() (Account, error) {
	account := AccountResponse{}
	requestURL, err := buildURL(w.baseURL, "ajax.php", "index", url.Values{})
	if err != nil {
		return account.Response, err
	}
	err = w.GetJSON(requestURL, &account)
	if err != nil {
		return account.Response, err
	}
	return account.Response, checkResponseStatus(account.Status, account.Error)
}

//GetMailbox retrieves mailbox information for the current user using the provided parameters.
func (w *WhatAPI) GetMailbox(params url.Values) (Mailbox, error) {
	mailbox := MailboxResponse{}
	requestURL, err := buildURL(w.baseURL, "ajax.php", "inbox", params)
	if err != nil {
		return mailbox.Response, err
	}
	err = w.GetJSON(requestURL, &mailbox)
	if err != nil {
		return mailbox.Response, err
	}
	return mailbox.Response, checkResponseStatus(mailbox.Status, mailbox.Error)
}

//GetConversation retrieves conversation information for the current user using the provided conversation id and parameters.
func (w *WhatAPI) GetConversation(id int) (Conversation, error) {
	conversation := ConversationResponse{}
	params := url.Values{}
	params.Set("type", "viewconv")
	params.Set("id", strconv.Itoa(id))
	requestURL, err := buildURL(w.baseURL, "ajax.php", "inbox", params)
	if err != nil {
		return conversation.Response, err
	}
	err = w.GetJSON(requestURL, &conversation)
	if err != nil {
		return conversation.Response, err
	}
	return conversation.Response, checkResponseStatus(conversation.Status, conversation.Error)
}

//GetNotifications retrieves notification information using the specifed parameters.
func (w *WhatAPI) GetNotifications(params url.Values) (Notifications, error) {
	notifications := NotificationsResponse{}
	requestURL, err := buildURL(w.baseURL, "ajax.php", "notifications", params)
	if err != nil {
		return notifications.Response, err
	}
	err = w.GetJSON(requestURL, &notifications)
	if err != nil {
		return notifications.Response, err
	}
	return notifications.Response, checkResponseStatus(notifications.Status, notifications.Error)
}

//GetAnnouncements retrieves announcement information.
func (w *WhatAPI) GetAnnouncements() (Announcements, error) {
	params := url.Values{}
	announcements := AnnouncementsResponse{}
	requestURL, err := buildURL(w.baseURL, "ajax.php", "announcements", params)
	if err != nil {
		return announcements.Response, err
	}
	err = w.GetJSON(requestURL, &announcements)
	if err != nil {
		return announcements.Response, err
	}
	return announcements.Response, checkResponseStatus(announcements.Status, announcements.Error)
}

//GetSubscriptions retrieves forum subscription information for the current user using the provided parameters.
func (w *WhatAPI) GetSubscriptions(params url.Values) (Subscriptions, error) {
	subscriptions := SubscriptionsResponse{}
	requestURL, err := buildURL(w.baseURL, "ajax.php", "subscriptions", params)
	if err != nil {
		return subscriptions.Response, err
	}
	err = w.GetJSON(requestURL, &subscriptions)
	if err != nil {
		return subscriptions.Response, err
	}
	return subscriptions.Response, checkResponseStatus(subscriptions.Status, subscriptions.Error)
}

//GetCategories retrieves forum category information.
func (w *WhatAPI) GetCategories() (Categories, error) {
	categories := CategoriesResponse{}
	params := url.Values{}
	params.Set("type", "main")
	requestURL, err := buildURL(w.baseURL, "ajax.php", "forum", params)
	if err != nil {
		return categories.Response, err
	}
	err = w.GetJSON(requestURL, &categories)
	if err != nil {
		return categories.Response, err
	}
	return categories.Response, checkResponseStatus(categories.Status, categories.Error)
}

//GetForum retrieves forum information using the provided forum id and parameters.
func (w *WhatAPI) GetForum(id int, params url.Values) (Forum, error) {
	forum := ForumResponse{}
	params.Set("type", "viewforum")
	params.Set("forumid", strconv.Itoa(id))
	requestURL, err := buildURL(w.baseURL, "ajax.php", "forum", params)
	if err != nil {
		return forum.Response, err
	}
	err = w.GetJSON(requestURL, &forum)
	if err != nil {
		return forum.Response, err
	}
	return forum.Response, checkResponseStatus(forum.Status, forum.Error)
}

//GetThread retrieves forum thread information using the provided thread id and parameters.
func (w *WhatAPI) GetThread(id int, params url.Values) (Thread, error) {
	thread := ThreadResponse{}
	params.Set("type", "viewthread")
	params.Set("threadid", strconv.Itoa(id))
	requestURL, err := buildURL(w.baseURL, "ajax.php", "forum", params)
	if err != nil {
		return thread.Response, err
	}
	err = w.GetJSON(requestURL, &thread)
	if err != nil {
		return thread.Response, err
	}
	return thread.Response, checkResponseStatus(thread.Status, thread.Error)
}

//GetArtistBookmarks retrieves artist bookmark information for the current user.
func (w *WhatAPI) GetArtistBookmarks() (ArtistBookmarks, error) {
	artistBookmarks := ArtistBookmarksResponse{}
	params := url.Values{}
	params.Set("type", "artists")
	requestURL, err := buildURL(w.baseURL, "ajax.php", "bookmarks", params)
	if err != nil {
		return artistBookmarks.Response, err
	}
	err = w.GetJSON(requestURL, &artistBookmarks)
	if err != nil {
		return artistBookmarks.Response, err
	}
	return artistBookmarks.Response, checkResponseStatus(artistBookmarks.Status, artistBookmarks.Error)
}

//GetTorrentBookmarks retrieves torrent bookmark information for the current user.
func (w *WhatAPI) GetTorrentBookmarks() (TorrentBookmarks, error) {
	torrentBookmarks := TorrentBookmarksResponse{}
	params := url.Values{}
	params.Set("type", "torrents")
	requestURL, err := buildURL(w.baseURL, "ajax.php", "bookmarks", params)
	if err != nil {
		return torrentBookmarks.Response, err
	}
	err = w.GetJSON(requestURL, &torrentBookmarks)
	if err != nil {
		return torrentBookmarks.Response, err
	}
	return torrentBookmarks.Response, checkResponseStatus(torrentBookmarks.Status, torrentBookmarks.Error)
}

//GetArtist retrieves artist information using the provided artist id and parameters.
func (w *WhatAPI) GetArtist(id int, params url.Values) (Artist, error) {
	artist := ArtistResponse{}
	params.Set("id", strconv.Itoa(id))
	requestURL, err := buildURL(w.baseURL, "ajax.php", "artist", params)
	if err != nil {
		return artist.Response, err
	}
	err = w.GetJSON(requestURL, &artist)
	if err != nil {
		return artist.Response, err
	}
	return artist.Response, checkResponseStatus(artist.Status, artist.Error)
}

//GetRequest retrieves request information using the provided request id and parameters.
func (w *WhatAPI) GetRequest(id int, params url.Values) (Request, error) {
	request := RequestResponse{}
	params.Set("id", strconv.Itoa(id))
	requestURL, err := buildURL(w.baseURL, "ajax.php", "request", params)
	if err != nil {
		return request.Response, err
	}
	err = w.GetJSON(requestURL, &request)
	if err != nil {
		return request.Response, err
	}
	return request.Response, checkResponseStatus(request.Status, request.Error)
}

//GetTorrent retrieves torrent information using the provided torrent id and parameters.
func (w *WhatAPI) GetTorrent(id int, params url.Values) (Torrent, error) {
	torrent := TorrentResponse{}
	params.Set("id", strconv.Itoa(id))
	requestURL, err := buildURL(w.baseURL, "ajax.php", "torrent", params)
	if err != nil {
		return torrent.Response, err
	}
	err = w.GetJSON(requestURL, &torrent)
	if err != nil {
		return torrent.Response, err
	}
	return torrent.Response, checkResponseStatus(torrent.Status, torrent.Error)
}

//GetTorrentGroup retrieves torrent group information using the provided torrent group id and parameters.
func (w *WhatAPI) GetTorrentGroup(id int, params url.Values) (TorrentGroup, error) {
	torrentGroup := TorrentGroupResponse{}
	params.Set("id", strconv.Itoa(id))
	requestURL, err := buildURL(w.baseURL, "ajax.php", "torrentgroup", params)
	if err != nil {
		return torrentGroup.Response, err
	}
	err = w.GetJSON(requestURL, &torrentGroup)
	if err != nil {
		return torrentGroup.Response, err
	}
	return torrentGroup.Response, checkResponseStatus(torrentGroup.Status, torrentGroup.Error)
}

//SearchTorrents retrieves torrent search results using the provided search string and parameters.
func (w *WhatAPI) SearchTorrents(searchStr string, params url.Values) (TorrentSearch, error) {
	torrentSearch := TorrentSearchResponse{}
	params.Set("searchstr", searchStr)
	requestURL, err := buildURL(w.baseURL, "ajax.php", "browse", params)
	if err != nil {
		return torrentSearch.Response, err
	}
	err = w.GetJSON(requestURL, &torrentSearch)
	if err != nil {
		return torrentSearch.Response, err
	}
	return torrentSearch.Response, checkResponseStatus(torrentSearch.Status, torrentSearch.Error)
}

//SearchRequests retrieves request search results using the provided search string and parameters.
func (w *WhatAPI) SearchRequests(searchStr string, params url.Values) (RequestsSearch, error) {
	requestsSearch := RequestsSearchResponse{}
	params.Set("search", searchStr)
	requestURL, err := buildURL(w.baseURL, "ajax.php", "requests", params)
	if err != nil {
		return requestsSearch.Response, err
	}
	err = w.GetJSON(requestURL, &requestsSearch)
	if err != nil {
		return requestsSearch.Response, err
	}
	return requestsSearch.Response, checkResponseStatus(requestsSearch.Status, requestsSearch.Error)
}

//SearchUsers retrieves user search results using the provided search string and parameters.
func (w *WhatAPI) SearchUsers(searchStr string, params url.Values) (UserSearch, error) {
	userSearch := UserSearchResponse{}
	params.Set("search", searchStr)
	requestURL, err := buildURL(w.baseURL, "ajax.php", "usersearch", params)
	if err != nil {
		return userSearch.Response, err
	}
	err = w.GetJSON(requestURL, &userSearch)
	if err != nil {
		return userSearch.Response, err
	}
	return userSearch.Response, checkResponseStatus(userSearch.Status, userSearch.Error)
}

//GetTopTenTorrents retrieves "top ten torrents" information using the provided parameters.
func (w *WhatAPI) GetTopTenTorrents(params url.Values) (TopTenTorrents, error) {
	topTenTorrents := TopTenTorrentsResponse{}
	params.Set("type", "torrents")
	requestURL, err := buildURL(w.baseURL, "ajax.php", "top10", params)
	if err != nil {
		return topTenTorrents.Response, err
	}
	err = w.GetJSON(requestURL, &topTenTorrents)
	if err != nil {
		return topTenTorrents.Response, err
	}
	return topTenTorrents.Response, checkResponseStatus(topTenTorrents.Status, topTenTorrents.Error)
}

//GetTopTenTags retrieves "top ten tags" information using the provided parameters.
func (w *WhatAPI) GetTopTenTags(params url.Values) (TopTenTags, error) {
	topTenTags := TopTenTagsResponse{}
	params.Set("type", "tags")
	requestURL, err := buildURL(w.baseURL, "ajax.php", "top10", params)
	if err != nil {
		return topTenTags.Response, err
	}
	err = w.GetJSON(requestURL, &topTenTags)
	if err != nil {
		return topTenTags.Response, err
	}
	return topTenTags.Response, checkResponseStatus(topTenTags.Status, topTenTags.Error)
}

//GetTopTenUsers retrieves "top tem users" information using the provided parameters.
func (w *WhatAPI) GetTopTenUsers(params url.Values) (TopTenUsers, error) {
	topTenUsers := TopTenUsersResponse{}
	params.Set("type", "users")
	requestURL, err := buildURL(w.baseURL, "ajax.php", "top10", params)
	if err != nil {
		return topTenUsers.Response, err
	}
	err = w.GetJSON(requestURL, &topTenUsers)
	if err != nil {
		return topTenUsers.Response, err
	}
	return topTenUsers.Response, checkResponseStatus(topTenUsers.Status, topTenUsers.Error)
}

//GetSimilarArtists retrieves similar artist information using the provided artist id and limit.
func (w *WhatAPI) GetSimilarArtists(id, limit int) (SimilarArtists, error) {
	similarArtists := SimilarArtists{}
	params := url.Values{}
	params.Set("id", strconv.Itoa(id))
	params.Set("limit", strconv.Itoa(limit))
	requestURL, err := buildURL(w.baseURL, "ajax.php", "similar_artists", params)
	if err != nil {
		return similarArtists, err
	}
	err = w.GetJSON(requestURL, &similarArtists)
	if err != nil {
		return similarArtists, err
	}
	return similarArtists, nil
}
