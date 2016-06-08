package whatapi

type AccountResponse struct {
	Status   string  `json:"status"`
	Error    string  `json:"error"`
	Response Account `json:"response"`
}

type AnnouncementsResponse struct {
	Status   string        `json:"status"`
	Error    string        `json:"error"`
	Response Announcements `json:"response"`
}

type ArtistResponse struct {
	Status   string `json:"status"`
	Error    string `json:"error"`
	Response Artist `json:"response"`
}

type ArtistBookmarksResponse struct {
	Status   string          `json:"status"`
	Error    string          `json:"error"`
	Response ArtistBookmarks `json:"response"`
}

type CategoriesResponse struct {
	Status   string     `json:"status"`
	Error    string     `json:"error"`
	Response Categories `json:"response"`
}

type ConversationResponse struct {
	Status   string       `json:"status"`
	Error    string       `json:"error"`
	Response Conversation `json:"response"`
}

type ForumResponse struct {
	Status   string `json:"status"`
	Error    string `json:"error"`
	Response Forum  `json:"response"`
}

type MailboxResponse struct {
	Status   string  `json:"status"`
	Error    string  `json:"error"`
	Response Mailbox `json:"response"`
}

type NotificationsResponse struct {
	Status   string        `json:"status"`
	Error    string        `json:"error"`
	Response Notifications `json:"response"`
}

type RequestResponse struct {
	Status   string  `json:"status"`
	Error    string  `json:"error"`
	Response Request `json:"response"`
}

type RequestsSearchResponse struct {
	Status   string         `json:"status"`
	Error    string         `json:"error"`
	Response RequestsSearch `json:"response"`
}

type SubscriptionsResponse struct {
	Status   string        `json:"status"`
	Error    string        `json:"error"`
	Response Subscriptions `json:"response"`
}

type ThreadResponse struct {
	Status   string `json:"status"`
	Error    string `json:"error"`
	Response Thread `json:"response"`
}

type TopTenTagsResponse struct {
	Status   string     `json:"status"`
	Error    string     `json:"error"`
	Response TopTenTags `json:"response"`
}

type TopTenTorrentsResponse struct {
	Status   string         `json:"status"`
	Error    string         `json:"error"`
	Response TopTenTorrents `json:"response"`
}

type TopTenUsersResponse struct {
	Status   string      `json:"status"`
	Error    string      `json:"error"`
	Response TopTenUsers `json:"response"`
}

type TorrentResponse struct {
	Status   string  `json:"status"`
	Error    string  `json:"error"`
	Response Torrent `json:"response"`
}

type TorrentBookmarksResponse struct {
	Status   string           `json:"status"`
	Error    string           `json:"error"`
	Response TorrentBookmarks `json:"response"`
}

type TorrentGroupResponse struct {
	Status   string       `json:"status"`
	Error    string       `json:"error"`
	Response TorrentGroup `json:"response"`
}

type TorrentSearchResponse struct {
	Status   string        `json:"status"`
	Error    string        `json:"error"`
	Response TorrentSearch `json:"response"`
}

type UserResponse struct {
	Status   string `json:"status"`
	Error    string `json:"error"`
	Response User   `json:"response"`
}

type UserSearchResponse struct {
	Status   string     `json:"status"`
	Error    string     `json:"error"`
	Response UserSearch `json:"response"`
}
