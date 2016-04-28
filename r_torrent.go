package whatapi

type Torrent struct {
	Group   GroupType `json:"group"`
	Torrent TorrentType `json:"torrent"`
}

type GroupType struct {
	WikiBody        string `json:"wikiBody"`
	WikiImage       string `json:"wikiImage"`
	ID              int    `json:"id"`
	Name            string `json:"name"`
	Year            int    `json:"year"`
	RecordLabel     string `json:"recordLabel"`
	CatalogueNumber string `json:"catalogueNumber"`
	ReleaseType     int    `json:"releaseType"`
	CategoryID      int    `json:"caregoryId"`
	CategoryName    string `json:"categoryName"`
	Time            string `json:"time"`
	VanityHouse     bool   `json:"vanityHouse"`
	MusicInfo       struct {
		Composers []string `json:"composers"`
		DJ        []string `json:"dj"`
		Artists   []struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		}
		With []struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"with"`
		Conductor []string `json:"conductor"`
		RemixedBy []string `json:"remixedBy"`
		Producer  []string `json:"producer"`
	} `json"musicInfo"`
	Tags []string `json"tags"`
}

type TorrentType struct {
	ID                      int    `json:"id"`
	Media                   string `json:"media"`
	Format                  string `json:"format"`
	Encoding                string `json:"encoding"`
	Remastered              bool   `json:"remastered"`
	RemasterYear            int    `json:"remasterYear"`
	RemasterTitle           string `json:"remasterTitle"`
	RemasterRecordLabel     string `json:"remasterRecordLabel"`
	RemasterCatalogueNumber string `json:"remasterCatalogueNumber"`
	Scene                   bool   `json:"scene"`
	HasLog                  bool   `json:"hasLog"`
	HasCue                  bool   `json:"hasCue"`
	LogScore                int    `json:"logScore"`
	FileCount               int    `json:"fileCount"`
	Size                    int    `json:"size"`
	Seeders                 int    `json:"seeders"`
	Leechers                int    `json:"leechers"`
	Snatched                int    `json:"snatched"`
	FreeTorrent             bool   `json:"freeTorrent"`
	Time                    string `json:"time"`
	Description             string `json:"description"`
	FileList                string `json:"fileList"`
	FilePath                string `json:"filePath"`
	UserID                  int    `json:"userID"`
	Username                string `json:"username"`
}