package whatapi

type Notifications struct {
	CurrentPages int `json:"currentPages"`
	Pages        int `json:"pages"`
	NumNew       int `json:"numNew"`
	Results      []struct {
		TorrentID        int    `json:"torrentId"`
		GroupID          int    `json:"groupId"`
		GroupName        string `json:"groupName"`
		GroupCategoryID  int    `json:"groupCategoryId"`
		TorrentTags      string `json:"torrentTags"`
		Size             int64  `json:"size"`
		FileCount        int    `json:"filecount"`
		Format           string `json:"format"`
		Encoding         string `json:"encoding"`
		Media            string `json:"mdia"`
		Scene            bool   `json:"scene"`
		GroupYear        int    `json:"groupYear"`
		RemasterYear     int    `json:"remasterYear"`
		RemasterTitle    string `json:"remasterTitle"`
		Snatched         int    `json:"snatched"`
		Seeders          int    `json:"seeders"`
		Leechers         int    `json:"leechers"`
		NotificationTime string `json:"notificationTime"`
		HasLog           bool   `json:"hasLog"`
		HasCue           bool   `json:"hasCue"`
		LogScore         int    `json:"logScore"`
		FreeTorrent      bool   `json:"freeTorrent"`
		LogInDB          bool   `json:"logInDb"`
		Unread           bool   `json:"unread"`
	} `json:"results"`
}