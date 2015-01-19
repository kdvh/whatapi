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
	DebugMode              = true
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

func buildURL(baseURL, path, query string) string {
	u, err := url.Parse(baseURL)
	checkErr(err)
	u.Path = path
	u.RawQuery = query
	if DebugMode {
		fmt.Println(u.String())
	}
	return u.String()
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func checkResponseStatus(status, errorStr string) {
	if status != "success" {
		if errorStr != "" {
			log.Fatal(ErrRequestFailedReason(errorStr))
		}
		log.Fatal(ErrRequestFailed)
	}
}

func NewSite(url string) *Site {
	s := new(Site)
	s.BaseURL = url
	cookieJar, err := cookiejar.New(nil)
	checkErr(err)
	s.Client = &http.Client{Jar: cookieJar}
	return s
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
		checkErr(err)
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		checkErr(err)
		return json.Unmarshal(body, v)
	}
	return ErrRequestFailedLogin
}

func (s *Site) Login(username, password string) {
	params := url.Values{}
	params.Set("username", username)
	params.Set("password", password)
	resp, err := s.Client.PostForm(s.BaseURL+"login.php?", params)
	checkErr(err)
	defer resp.Body.Close()
	if resp.Request.URL.String()[len(s.BaseURL):] != "index.php" {
		log.Fatal(ErrLoginFailed)
	}
	s.LoggedIn = true
	account := s.GetAccount()
	s.Username = account.Username
	s.AuthKey = account.AuthKey
	s.PassKey = account.PassKey
}

func (s *Site) Logout() {
	params := url.Values{"auth": {s.AuthKey}}
	_, err := s.Client.Get(buildURL(s.BaseURL, "logout.php", params.Encode()))
	checkErr(err)
	s.LoggedIn, s.Username, s.AuthKey, s.PassKey = false, "", "", ""
}


func (s *Site) CreateDownloadURL(id int) string {
	params := url.Values{}
	params.Set("action", "download")
	params.Set("id", strconv.Itoa(id))
	params.Set("authkey", s.AuthKey)
	params.Set("torrent_pass", s.PassKey)
	return buildURL(s.BaseURL, "torrents.php", params.Encode())
}

func (s *Site) GetAccount() AccountResponse {
	account := Account{}
	query := buildQuery("index", url.Values{})
	err := s.GetJSON(buildURL(s.BaseURL, "ajax.php", query), &account)
	checkErr(err)
	checkResponseStatus(account.Status, account.Error)
	return account.Response
}

func (s *Site) GetMailbox(params url.Values) MailboxResponse {
	mailbox := Mailbox{}
	query := buildQuery("inbox", params)
	err := s.GetJSON(buildURL(s.BaseURL, "ajax.php", query), &mailbox)
	checkErr(err)
	checkResponseStatus(mailbox.Status, mailbox.Error)
	return mailbox.Response
}

func (s *Site) GetConversation(id int) ConversationResponse {
	conversation := Conversation{}
	params := url.Values{}
	params.Set("type", "viewconv")
	params.Set("id", strconv.Itoa(id))
	query := buildQuery("inbox", params)
	err := s.GetJSON(buildURL(s.BaseURL, "ajax.php", query), &conversation)
	checkErr(err)
	checkResponseStatus(conversation.Status, conversation.Error)
	return conversation.Response
}

func (s *Site) GetNotifications(params url.Values) NotificationsResponse {
	notifications := Notifications{}
	query := buildQuery("notifications", params)
	err := s.GetJSON(buildURL(s.BaseURL, "ajax.php", query), &notifications)
	checkErr(err)
	checkResponseStatus(notifications.Status, notifications.Error)
	return notifications.Response
}

func (s *Site) GetAnnouncements() AnnouncementsResponse {
	params := url.Values{}
	announcements := Announcements{}
	query := buildQuery("announcements", params)
	err := s.GetJSON(buildURL(s.BaseURL, "ajax.php", query), &announcements)
	checkErr(err)
	checkResponseStatus(announcements.Status, announcements.Error)
	return announcements.Response
}

func (s *Site) GetSubscriptions(params url.Values) SubscriptionsResponse {
	subscriptions := Subscriptions{}
	query := buildQuery("subscriptions", params)
	err := s.GetJSON(buildURL(s.BaseURL, "ajax.php", query), &subscriptions)
	checkErr(err)
	checkResponseStatus(subscriptions.Status, subscriptions.Error)
	return subscriptions.Response
}

func (s *Site) GetCategories() CategoriesResponse {
	categories := Categories{}
	params := url.Values{}
	params.Set("type", "main")
	query := buildQuery("forum", params)
	err := s.GetJSON(buildURL(s.BaseURL, "ajax.php", query), &categories)
	checkErr(err)
	checkResponseStatus(categories.Status, categories.Error)
	return categories.Response
}

func (s *Site) GetForum(forumID int, params url.Values) ForumResponse {
	forum := Forum{}
	params.Set("type", "viewforum")
	params.Set("forumid", strconv.Itoa(forumID))
	query := buildQuery("forum", params)
	err := s.GetJSON(buildURL(s.BaseURL, "ajax.php", query), &forum)
	checkErr(err)
	checkResponseStatus(forum.Status, forum.Error)
	return forum.Response
}

func (s *Site) GetThread(threadID int, params url.Values) ThreadResponse {
	thread := Thread{}
	params.Set("type", "viewthread")
	params.Set("threadid", strconv.Itoa(threadID))
	query := buildQuery("forum", params)
	err := s.GetJSON(buildURL(s.BaseURL, "ajax.php", query), &thread)
	checkErr(err)
	checkResponseStatus(thread.Status, thread.Error)
	return thread.Response
}

func (s *Site) GetArtistBookmarks() ArtistBookmarksResponse {
	artistBookmarks := ArtistBookmarks{}
	params := url.Values{}
	params.Set("type", "artists")
	query := buildQuery("bookmarks", params)
	err := s.GetJSON(buildURL(s.BaseURL, "ajax.php", query), &artistBookmarks)
	checkErr(err)
	checkResponseStatus(artistBookmarks.Status, artistBookmarks.Error)
	return artistBookmarks.Response
}

func (s *Site) GetTorrentBookmarks() TorrentBookmarksResponse {
	torrentBookmarks := TorrentBookmarks{}
	params := url.Values{}
	params.Set("type", "torrents")
	query := buildQuery("bookmarks", params)
	err := s.GetJSON(buildURL(s.BaseURL, "ajax.php", query), &torrentBookmarks)
	checkErr(err)
	checkResponseStatus(torrentBookmarks.Status, torrentBookmarks.Error)
	return torrentBookmarks.Response
}

func (s *Site) GetArtist(id int, params url.Values) ArtistResponse {
	artist := Artist{}
	params.Set("id", strconv.Itoa(id))
	query := buildQuery("artist", params)
	err := s.GetJSON(buildURL(s.BaseURL, "ajax.php", query), &artist)
	checkErr(err)
	checkResponseStatus(artist.Status, artist.Error)
	return artist.Response
}

func (s *Site) GetRequest(id int, params url.Values) RequestResponse {
	request := Request{}
	params.Set("id", strconv.Itoa(id))
	query := buildQuery("request", params)
	err := s.GetJSON(buildURL(s.BaseURL, "ajax.php", query), &request)
	checkErr(err)
	checkResponseStatus(request.Status, request.Error)
	return request.Response
}

func (s *Site) GetTorrent(id int, params url.Values) TorrentResponse {
	torrent := Torrent{}
	params.Set("id", strconv.Itoa(id))
	query := buildQuery("torrent", params)
	err := s.GetJSON(buildURL(s.BaseURL, "ajax.php", query), &torrent)
	checkErr(err)
	checkResponseStatus(torrent.Status, torrent.Error)
	return torrent.Response
}

func (s *Site) GetTorrentGroup(id int, params url.Values) TorrentGroupResponse {
	torrentGroup := TorrentGroup{}
	params.Set("id", strconv.Itoa(id))
	query := buildQuery("torrentgroup", params)
	err := s.GetJSON(buildURL(s.BaseURL, "ajax.php", query), &torrentGroup)
	checkErr(err)
	checkResponseStatus(torrentGroup.Status, torrentGroup.Error)
	return torrentGroup.Response
}

func (s *Site) SearchTorrents(searchStr string, params url.Values) TorrentSearchResponse {
	torrentSearch := TorrentSearch{}
	params.Set("searchstr", searchStr)
	query := buildQuery("browse", params)
	err := s.GetJSON(buildURL(s.BaseURL, "ajax.php", query), &torrentSearch)
	checkErr(err)
	checkResponseStatus(torrentSearch.Status, torrentSearch.Error)
	return torrentSearch.Response
}

func (s *Site) SearchRequests(searchStr string, params url.Values) RequestsSearchResponse {
	requestsSearch := RequestsSearch{}
	params.Set("search", searchStr)
	query := buildQuery("requests", params)
	err := s.GetJSON(buildURL(s.BaseURL, "ajax.php", query), &requestsSearch)
	checkErr(err)
	checkResponseStatus(requestsSearch.Status, requestsSearch.Error)
	return requestsSearch.Response
}

func (s *Site) SearchUsers(searchStr string, params url.Values) UserSearchResponse {
	userSearch := UserSearch{}
	params.Set("search", searchStr)
	query := buildQuery("usersearch", params)
	err := s.GetJSON(buildURL(s.BaseURL, "ajax.php", query), &userSearch)
	checkErr(err)
	checkResponseStatus(userSearch.Status, userSearch.Error)
	return userSearch.Response
}

func (s *Site) GetTopTenTorrents(params url.Values) TopTenTorrentsResponse {
	topTenTorrents := TopTenTorrents{}
	params.Set("type", "torrents")
	query := buildQuery("top10", params)
	err := s.GetJSON(buildURL(s.BaseURL, "ajax.php", query), &topTenTorrents)
	checkErr(err)
	checkResponseStatus(topTenTorrents.Status, topTenTorrents.Error)
	return topTenTorrents.Response
}

func (s *Site) GetTopTenTags(params url.Values) TopTenTagsResponse {
	topTenTags := TopTenTags{}
	params.Set("type", "tags")
	query := buildQuery("top10", params)
	err := s.GetJSON(buildURL(s.BaseURL, "ajax.php", query), &topTenTags)
	checkErr(err)
	checkResponseStatus(topTenTags.Status, topTenTags.Error)
	return topTenTags.Response
}

func (s *Site) GetTopTenUsers(params url.Values) TopTenUsersResponse {
	topTenUsers := TopTenUsers{}
	params.Set("type", "users")
	query := buildQuery("top10", params)
	err := s.GetJSON(buildURL(s.BaseURL, "ajax.php", query), &topTenUsers)
	checkErr(err)
	checkResponseStatus(topTenUsers.Status, topTenUsers.Error)
	return topTenUsers.Response
}

func (s *Site) GetSimilarArtists(id, limit int) SimilarArtists {
	similarArtists := SimilarArtists{}
	params := url.Values{}
	params.Set("id", strconv.Itoa(id))
	params.Set("limit", strconv.Itoa(limit))
	query := buildQuery("similar_artists", params)
	err := s.GetJSON(buildURL(s.BaseURL, "ajax.php", query), &similarArtists)
	checkErr(err)
	return similarArtists
}

