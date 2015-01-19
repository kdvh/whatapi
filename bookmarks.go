package whatapi

type ArtistBookmarks struct {
	Status   string                  `json:"status"`
	Error    string                  `json:"error"`
	Response ArtistBookmarksResponse `json:"response"`
}

type TorrentBookmarks struct {
	Status   string                   `json:"status"`
	Error    string                   `json:"error"`
	Response TorrentBookmarksResponse `json:"response"`
}

type ArtistBookmarksResponse struct {
	Artists []struct {
		ArtistID   int    `json:"artistId"`
		ArtistName string `json:"artistName"`
	} `json:"artists"`
}

type TorrentBookmarksResponse struct {
	Bookmarks []struct {
		ID              int           `json:"id"`
		Name            string        `json:"name"`
		Year            int           `json:"year"`
		RecordLabel     string        `json:"recordLabel"`
		CatalogueNumber string        `json:"catalogueNumber"`
		TagList         string        `json:"tagList"`
		ReleaseType     string        `json:"releastType"`
		VanityHouse     bool          `json:"vanityHouse"`
		Image           string        `json:"image"`
		Torrents        []TorrentType `json:"torrents"`
	} `json:"bookmarks"`
}
