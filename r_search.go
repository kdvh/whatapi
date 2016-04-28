package whatapi

type RequestsSearch struct {
	CurrentPage int `json:"currentPage"`
	Pages       int `json:"pages"`
	Results     []struct {
		RequestID     int    `json:"requestId"`
		RequestorID   int    `json:"requestorId"`
		ReqyestorName string `json:"requestorName"`
		TimeAdded     string `json:"timeAdded"`
		LastVote      string `json:"lastVote"`
		VoteCount     int    `json:"voteCount"`
		Bounty        int64  `json:"bounty"`
		CategoryID    int    `json:"categoryId"`
		CategoryName  string `json:"categoryName"`
		Artists       [][]struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"artists"`
		Title           string `json:"title"`
		Year            int    `json:"year"`
		Image           string `json:"image"`
		Description     string `json:"description"`
		CatalogueNumber string `json:"catalogueNumber"`
		ReleaseType     string `json:"releaseType"`
		BitrateList     string `json:"bitrateList"`
		FormatList      string `json:"formatList"`
		MediaList       string `json:"mediaList"`
		LogCue          string `json:"logCue"`
		IsFilled        bool   `json:"isFilled"`
		FillerID        int    `json:"fillerId"`
		FillerName      string `json:"fillerName"`
		TorrentID       int    `json:"torrentId"`
		TimeFilled      string `json:"timeFilled"`
	} `json:"results"`
}

type TorrentSearch struct {
	CurrentPage int `json:"currentPage"`
	Pages       int `json:"pages"`
	Results     []struct {
		GroupID       int      `json:"groupId"`
		GroupName     string   `json:"groupName"`
		Artist        string   `json:"artist"`
		Tags          []string `json:"tags"`
		Bookmarked    bool     `json:"bookmarked"`
		VanityHouse   bool     `json:"vanityHouse"`
		GroupYear     int      `json:"groupYear"`
		ReleaseType   string   `json:"releasetType"`
		GroupTime     string   `json:"groupTime"`
		TotalSnatched int      `json:"totalSnatched"`
		TotalSeeders  int      `json:"totalSeeders"`
		TotalLeechers int      `json:"totalLeechers"`
		Torrents      []struct {
			TorrentID int `json:"torrentId"`
			EditionID int `json:"editionId"`
			Artists   []struct {
				ID      int    `json:"id"`
				Name    string `json:"name"`
				AliasID int    `json:"aliasid"`
			} `json:"artists"`
			Remastered              bool   `json:"remastered"`
			RemasterYear            int    `json:"remasterYear"`
			RemasterCatalogueNumber string `json:"remasterCatalogueNumber"`
			RemasterTitle           string `json:"remasterTitle"`
			Media                   string `json:"media"`
			Encoding                string `json:"encoding"`
			Format                  string `json:"format"`
			HasLog                  bool   `json:"hasLog"`
			LogScore                int    `json:"logScore"`
			HasCue                  bool   `json:"hasCue"`
			Scene                   bool   `json:"scene"`
			VanityHouse             bool   `json:"vanityHouse"`
			FileCount               int    `json:"fileCount"`
			Time                    string `json:"time"`
			Size                    int64  `json:"size"`
			Snatches                int    `json:"snatches"`
			Seeders                 int    `json:"seeders"`
			Leechers                int    `json:"leechers"`
			IsFreeleech             bool   `json:"isFreeleech"`
			IsNeutralLeech          bool   `json:"isNeutralLeech"`
			IsPersonalFreeleech     bool   `json:"isPersonalFreeleech"`
			CanUseToken             bool   `json:"canUseToken"`
		} `json:"torrents"`
	} `json:"results"`
}

type UserSearch struct {
	CurrentPage int `json:"currentPage"`
	Pages       int `json:"pages"`
	Results     []struct {
		UserID   int    `json:"userId"`
		Username string `json:"username"`
		Donor    bool   `json:"donor"`
		Warned   bool   `json:"warned"`
		Enabled  bool   `json:"enabled"`
		Class    string `json:"class"`
	} `json:"results"`
}
