package whatapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strconv"
)

var (
	ErrLoginFailed         = errors.New("Login failed")
	ErrRequestFailed       = errors.New("Request failed")
	ErrRequestFailedLogin  = errors.New("Request failed: not logged in")
	ErrRequestFailedReason = func(err string) error { return fmt.Errorf("Request failed: %s", err) }
	DebugMode              = false
)

func buildQuery(action string, params url.Values) string {
	query := make(url.Values)
	query.Set("action", action)
	for param, values := range params {
		for _, value := range values {
			query.Set(param, value)
		}
	}
	return query.Encode()
}

func buildURL(baseURL, path, query string) (string, error) {
	u, err := url.Parse(baseURL)
	if err != nil {
		return u.String(), err
	}
	u.Path = path
	u.RawQuery = query
	if DebugMode {
		fmt.Println(u.String())
	}
	return u.String(), nil
}

func checkResponseStatus(status, errorStr string) {
	if status != "success" {
		if errorStr != "" {
			log.Println(ErrRequestFailedReason(errorStr))
		}
		log.Println(ErrRequestFailed)
	}
}

func NewSite(url string) (*Site, error) {
	s := new(Site)
	s.BaseURL = url
	cookieJar, err := cookiejar.New(nil)
	if err != nil {
		return s, err
	}
	s.Client = &http.Client{Jar: cookieJar}
	return s, nil
}

type Site struct {
	BaseURL  string
	Client   *http.Client
	LoggedIn bool
	Username string
	AuthKey  string
	PassKey  string
}

func (s *Site) GetJSON(url string, v interface{}) error {
	if s.LoggedIn {
		resp, err := s.Client.Get(url)
		if err != nil {
			return err
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		return json.Unmarshal(body, v)
	}
	return ErrRequestFailedLogin
}

func (s *Site) Login(username, password string) error {
	params := url.Values{}
	params.Set("username", username)
	params.Set("password", password)
	resp, err := s.Client.PostForm(s.BaseURL+"login.php?", params)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.Request.URL.String()[len(s.BaseURL):] != "index.php" {
		log.Println(ErrLoginFailed)
		return err
	}
	s.LoggedIn = true
	account, err := s.GetAccount()
	if err != nil {
		return err
	}
	s.Username = account.Username
	s.AuthKey = account.AuthKey
	s.PassKey = account.PassKey
	return nil
}

func (s *Site) Logout() error {
	params := url.Values{"auth": {s.AuthKey}}
	url, err := buildURL(s.BaseURL, "logout.php", params.Encode())
	if err != nil {
		return err
	}
	_, err = s.Client.Get(url)
	if err != nil {
		return err
	}
	s.LoggedIn, s.Username, s.AuthKey, s.PassKey = false, "", "", ""
	return nil
}

func (s *Site) CreateDownloadURL(id int) (string, error) {
	params := url.Values{}
	params.Set("action", "download")
	params.Set("id", strconv.Itoa(id))
	params.Set("authkey", s.AuthKey)
	params.Set("torrent_pass", s.PassKey)
	url, err := buildURL(s.BaseURL, "torrents.php", params.Encode())
	return url, err
}

func (s *Site) GetAccount() (AccountResponse, error) {
	account := Account{}
	query := buildQuery("index", url.Values{})
	url, err := buildURL(s.BaseURL, "ajax.php", query)
	if err != nil {
		return account.Response, err
	}
	err = s.GetJSON(url, &account)
	if err != nil {
		return account.Response, err
	}
	checkResponseStatus(account.Status, account.Error)
	return account.Response, nil
}

func (s *Site) GetMailbox(params url.Values) (MailboxResponse, error) {
	mailbox := Mailbox{}
	query := buildQuery("inbox", params)
	url, err := buildURL(s.BaseURL, "ajax.php", query)
	if err != nil {
		return mailbox.Response, err
	}
	err = s.GetJSON(url, &mailbox)
	if err != nil {
		return mailbox.Response, err
	}
	checkResponseStatus(mailbox.Status, mailbox.Error)
	return mailbox.Response, nil
}

func (s *Site) GetConversation(id int) (ConversationResponse, error) {
	conversation := Conversation{}
	params := url.Values{}
	params.Set("type", "viewconv")
	params.Set("id", strconv.Itoa(id))
	query := buildQuery("inbox", params)
	url, err := buildURL(s.BaseURL, "ajax.php", query)
	if err != nil {
		return conversation.Response, err
	}
	err = s.GetJSON(url, &conversation)
	if err != nil {
		return conversation.Response, err
	}
	checkResponseStatus(conversation.Status, conversation.Error)
	return conversation.Response, nil
}

func (s *Site) GetNotifications(params url.Values) (NotificationsResponse, error) {
	notifications := Notifications{}
	query := buildQuery("notifications", params)
	url, err := buildURL(s.BaseURL, "ajax.php", query)
	if err != nil {
	}
	err = s.GetJSON(url, &notifications)
	if err != nil {
		return notifications.Response, err
	}
	checkResponseStatus(notifications.Status, notifications.Error)
	return notifications.Response, nil
}

func (s *Site) GetAnnouncements() (AnnouncementsResponse, error) {
	params := url.Values{}
	announcements := Announcements{}
	query := buildQuery("announcements", params)
	url, err := buildURL(s.BaseURL, "ajax.php", query)
	if err != nil {
		return announcements.Response, err
	}
	err = s.GetJSON(url, &announcements)
	if err != nil {
		return announcements.Response, err
	}
	checkResponseStatus(announcements.Status, announcements.Error)
	return announcements.Response, nil
}

func (s *Site) GetSubscriptions(params url.Values) (SubscriptionsResponse, error) {
	subscriptions := Subscriptions{}
	query := buildQuery("subscriptions", params)
	url, err := buildURL(s.BaseURL, "ajax.php", query)
	if err != nil {

	}
	err = s.GetJSON(url, &subscriptions)
	if err != nil {
		return subscriptions.Response, err
	}
	checkResponseStatus(subscriptions.Status, subscriptions.Error)
	return subscriptions.Response, nil
}

func (s *Site) GetCategories() (CategoriesResponse, error) {
	categories := Categories{}
	params := url.Values{}
	params.Set("type", "main")
	query := buildQuery("forum", params)
	url, err := buildURL(s.BaseURL, "ajax.php", query)
	if err != nil {
		return categories.Response, err
	}
	err = s.GetJSON(url, &categories)
	if err != nil {
		return categories.Response, err
	}
	checkResponseStatus(categories.Status, categories.Error)
	return categories.Response, nil
}

func (s *Site) GetForum(forumID int, params url.Values) (ForumResponse, error) {
	forum := Forum{}
	params.Set("type", "viewforum")
	params.Set("forumid", strconv.Itoa(forumID))
	query := buildQuery("forum", params)
	url, err := buildURL(s.BaseURL, "ajax.php", query)
	if err != nil {
		return forum.Response, err
	}
	err = s.GetJSON(url, &forum)
	if err != nil {
		return forum.Response, err
	}
	checkResponseStatus(forum.Status, forum.Error)
	return forum.Response, nil
}

func (s *Site) GetThread(threadID int, params url.Values) (ThreadResponse, error) {
	thread := Thread{}
	params.Set("type", "viewthread")
	params.Set("threadid", strconv.Itoa(threadID))
	query := buildQuery("forum", params)
	url, err := buildURL(s.BaseURL, "ajax.php", query)
	if err != nil {
		return thread.Response, err
	}
	err = s.GetJSON(url, &thread)
	if err != nil {
		return thread.Response, err
	}
	checkResponseStatus(thread.Status, thread.Error)
	return thread.Response, nil
}

func (s *Site) GetArtistBookmarks() (ArtistBookmarksResponse, error) {
	artistBookmarks := ArtistBookmarks{}
	params := url.Values{}
	params.Set("type", "artists")
	query := buildQuery("bookmarks", params)
	url, err := buildURL(s.BaseURL, "ajax.php", query)
	if err != nil {
		return artistBookmarks.Response, err
	}
	err = s.GetJSON(url, &artistBookmarks)
	if err != nil {
		return artistBookmarks.Response, err
	}
	checkResponseStatus(artistBookmarks.Status, artistBookmarks.Error)
	return artistBookmarks.Response, nil
}

func (s *Site) GetTorrentBookmarks() (TorrentBookmarksResponse, error) {
	torrentBookmarks := TorrentBookmarks{}
	params := url.Values{}
	params.Set("type", "torrents")
	query := buildQuery("bookmarks", params)
	url, err := buildURL(s.BaseURL, "ajax.php", query)
	if err != nil {
		return torrentBookmarks.Response, err
	}
	err = s.GetJSON(url, &torrentBookmarks)
	if err != nil {
		return torrentBookmarks.Response, err
	}
	checkResponseStatus(torrentBookmarks.Status, torrentBookmarks.Error)
	return torrentBookmarks.Response, nil
}

func (s *Site) GetArtist(id int, params url.Values) (ArtistResponse, error) {
	artist := Artist{}
	params.Set("id", strconv.Itoa(id))
	query := buildQuery("artist", params)
	url, err := buildURL(s.BaseURL, "ajax.php", query)
	if err != nil {
		return artist.Response, err
	}
	err = s.GetJSON(url, &artist)
	if err != nil {
		return artist.Response, err
	}
	checkResponseStatus(artist.Status, artist.Error)
	return artist.Response, nil
}

func (s *Site) GetRequest(id int, params url.Values) (RequestResponse, error) {
	request := Request{}
	params.Set("id", strconv.Itoa(id))
	query := buildQuery("request", params)
	url, err := buildURL(s.BaseURL, "ajax.php", query)
	if err != nil {
		return request.Response, err
	}
	err = s.GetJSON(url, &request)
	checkResponseStatus(request.Status, request.Error)
	return request.Response, nil
}

func (s *Site) GetTorrent(id int, params url.Values) (TorrentResponse, error) {
	torrent := Torrent{}
	params.Set("id", strconv.Itoa(id))
	query := buildQuery("torrent", params)
	url, err := buildURL(s.BaseURL, "ajax.php", query)
	if err != nil {
		return torrent.Response, err
	}
	err = s.GetJSON(url, &torrent)
	if err != nil {
		return torrent.Response, err
	}
	checkResponseStatus(torrent.Status, torrent.Error)
	return torrent.Response, nil
}

func (s *Site) GetTorrentGroup(id int, params url.Values) (TorrentGroupResponse, error) {
	torrentGroup := TorrentGroup{}
	params.Set("id", strconv.Itoa(id))
	query := buildQuery("torrentgroup", params)
	url, err := buildURL(s.BaseURL, "ajax.php", query)
	if err != nil {
		return torrentGroup.Response, err
	}
	err = s.GetJSON(url, &torrentGroup)
	if err != nil {
		return torrentGroup.Response, err
	}
	checkResponseStatus(torrentGroup.Status, torrentGroup.Error)
	return torrentGroup.Response, nil
}

func (s *Site) SearchTorrents(searchStr string, params url.Values) (TorrentSearchResponse, error) {
	torrentSearch := TorrentSearch{}
	params.Set("searchstr", searchStr)
	query := buildQuery("browse", params)
	url, err := buildURL(s.BaseURL, "ajax.php", query)
	if err != nil {
		return torrentSearch.Response, err
	}
	err = s.GetJSON(url, &torrentSearch)
	if err != nil {
		return torrentSearch.Response, err
	}
	checkResponseStatus(torrentSearch.Status, torrentSearch.Error)
	return torrentSearch.Response, nil
}

func (s *Site) SearchRequests(searchStr string, params url.Values) (RequestsSearchResponse, error) {
	requestsSearch := RequestsSearch{}
	params.Set("search", searchStr)
	query := buildQuery("requests", params)
	url, err := buildURL(s.BaseURL, "ajax.php", query)
	if err != nil {
		return requestsSearch.Response, err
	}
	err = s.GetJSON(url, &requestsSearch)
	if err != nil {
		return requestsSearch.Response, err
	}
	checkResponseStatus(requestsSearch.Status, requestsSearch.Error)
	return requestsSearch.Response, nil
}

func (s *Site) SearchUsers(searchStr string, params url.Values) (UserSearchResponse, error) {
	userSearch := UserSearch{}
	params.Set("search", searchStr)
	query := buildQuery("usersearch", params)
	url, err := buildURL(s.BaseURL, "ajax.php", query)
	if err != nil {
		return userSearch.Response, err
	}
	err = s.GetJSON(url, &userSearch)
	if err != nil {
		return userSearch.Response, err
	}
	checkResponseStatus(userSearch.Status, userSearch.Error)
	return userSearch.Response, nil
}

func (s *Site) GetTopTenTorrents(params url.Values) (TopTenTorrentsResponse, error) {
	topTenTorrents := TopTenTorrents{}
	params.Set("type", "torrents")
	query := buildQuery("top10", params)
	url, err := buildURL(s.BaseURL, "ajax.php", query)
	if err != nil {
		return topTenTorrents.Response, err
	}
	err = s.GetJSON(url, &topTenTorrents)
	if err != nil {
		return topTenTorrents.Response, err
	}
	checkResponseStatus(topTenTorrents.Status, topTenTorrents.Error)
	return topTenTorrents.Response, nil
}

func (s *Site) GetTopTenTags(params url.Values) (TopTenTagsResponse, error) {
	topTenTags := TopTenTags{}
	params.Set("type", "tags")
	query := buildQuery("top10", params)
	url, err := buildURL(s.BaseURL, "ajax.php", query)
	if err != nil {
		return topTenTags.Response, err
	}
	err = s.GetJSON(url, &topTenTags)
	if err != nil {
		return topTenTags.Response, err
	}
	checkResponseStatus(topTenTags.Status, topTenTags.Error)
	return topTenTags.Response, err
}

func (s *Site) GetTopTenUsers(params url.Values) (TopTenUsersResponse, error) {
	topTenUsers := TopTenUsers{}
	params.Set("type", "users")
	query := buildQuery("top10", params)
	url, err := buildURL(s.BaseURL, "ajax.php", query)
	if err != nil {
		return topTenUsers.Response, err
	}
	err = s.GetJSON(url, &topTenUsers)
	if err != nil {
		return topTenUsers.Response, err
	}
	checkResponseStatus(topTenUsers.Status, topTenUsers.Error)
	return topTenUsers.Response, nil
}

func (s *Site) GetSimilarArtists(id, limit int) (SimilarArtists, error) {
	similarArtists := SimilarArtists{}
	params := url.Values{}
	params.Set("id", strconv.Itoa(id))
	params.Set("limit", strconv.Itoa(limit))
	query := buildQuery("similar_artists", params)
	url, err := buildURL(s.BaseURL, "ajax.php", query)
	if err != nil {
		return similarArtists, err
	}
	err = s.GetJSON(url, &similarArtists)
	if err != nil {
		return similarArtists, err
	}
	return similarArtists, err
}
