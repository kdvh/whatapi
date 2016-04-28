package whatapi

type User struct {
	Username    string `json:"username"`
	Avatar      string `json:"avatar"`
	IsFriend    bool   `json:"isFriend"`
	ProfileText string `json:"profileText"`
	Stats       struct {
		JoinedDate    string  `json:"joinedDate"`
		LastAccess    string  `json:"lastAccess"`
		Uploaded      int64   `json:"uploaded"`
		Downloaded    int64   `json:"downloaded"`
		Ratio         string  `json:"ratio"`
		RequiredRatio float64 `json:"requiredRatio"`
	} `json:"stats"`
	Ranks struct {
		Uploaded   int `json:"uploaded"`
		Downloaded int `json:"downloaded"`
		Uploads    int `json:"uploads"`
		Requests   int `json:"requests"`
		Bounty     int `json:"bounty"`
		Posts      int `json:"posts"`
		Artists    int `json:"artists"`
		Overall    int `json:"overall"`
	} `json:"ranks"`
	Personal struct {
		Class        string `json:"class"`
		Paranoia     int    `json:"paranoia"`
		ParanoiaText string `json:"paranoiaText"`
		Donor        bool   `json:"donor"`
		Warned       bool   `json:"warned"`
		Enabled      bool   `json:"enabled"`
		PassKey      string `json:"passKey"`
	} `json:"personal"`
	Community struct {
		Posts           int `json:"posts"`
		TorrentComments int `json:"torrentComments"`
		CollagesStarted int `json:"collagesStarted"`
		CollagesContrib int `json:"collagesContrib"`
		RequestsFilled  int `json:"requestsFilled"`
		RequestsVoted   int `json:"requestsVoted"`
		PerfectFlacs    int `json:"perfectFlacs"`
		Uploaded        int `json:"uploaded"`
		Groups          int `json:"groups"`
		Seeding         int `json:"seeding"`
		Leeching        int `json:"leeching"`
		Snatched        int `json:"snatched"`
		Invited         int `json:"invited"`
	} `json:"community"`
}
