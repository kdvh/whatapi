package whatapi

type ArtistBookmarks struct {
	Artists []struct {
		ArtistID   int    `json:"artistId"`
		ArtistName string `json:"artistName"`
	} `json:"artists"`
}

type TorrentBookmarks struct {
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
