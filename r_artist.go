package whatapi

type Artist struct {
	ID                   int    `json:"id"`
	Name                 string `json:"name"`
	NotificationsEnabled bool   `json:"notificationsEnabled"`
	HasBookmarked        bool   `json:"hasBookmarked"`
	Image                string `json:"image"`
	Body                 string `json:"body"`
	VanityHouse          bool   `json:"vanityHouse"`
	Tags                 []struct {
		Name  string `json:"name"`
		Count int    `json:"count"`
	} `json:"tags"`
	SimilarArtists []struct {
		ArtistID  int    `json:"artistId"`
		Name      string `json:"name"`
		Score     int    `json:"score"`
		SimilarID int    `json:"similarId"`
	} `json:"similarArtists"`
	Statistics struct {
		NumGroups   int `json:"numGroups"`
		NumTorrents int `json:"numTorrents"`
		NumSeeders  int `json:"numSeeders"`
		NumLeechers int `json:"numLeechers"`
		NumSnatches int `json:"numSnatches"`
	} `json:"statistics"`
	TorrentGroup []struct {
		GroupID              int           `json:"groupId"`
		GroupYear            int           `json:"groupYear"`
		GroupRecordLabel     string        `json:"groupRecordLabel"`
		GroupCatalogueNumber string        `json:"groupCatalogueNumber"`
		Tags                 []string      `json:"tags"`
		ReleaseType          int           `json:"releaseType"`
		GroupVanityHouse     bool          `json:"groupVanityHouse"`
		HasBookmarked        bool          `json:"hasBookmarked"`
		Torrent              []TorrentType `json:"torrent"`
	} `json:"torrentgroup"`
	Requests []struct {
		RequestID  int    `json:"requestId"`
		CategoryID int    `json:"categoryId"`
		Title      string `json:"title"`
		Year       int    `json:"year"`
		TimeAdded  string `json:"timeAdded"`
		Votes      int    `json:"votes"`
		Bounty     int64  `json:"bounty"`
	} `json:"requests"`
}
