package whatapi

type Request struct {
	RequestID       int     `json:"requestId"`
	RequestiorID    int     `json:"requestorId"`
	RequestorName   string  `json:"requestorName"`
	RequestTax      float64 `json:"requestTax"`
	TimeAdded       string  `json:"timeAdded"`
	CanEdit         bool    `json:"canEdit"`
	CanVote         bool    `json:"canVote"`
	MinimumVote     int     `json:"minimumVote"`
	VoteCount       int     `json:"voteCount"`
	LastVote        string  `json:"lastVote"`
	TopContributors []struct {
		UserID   int    `json:"userId"`
		UserName string `json:"userName"`
		Bounty   int64  `json:"bounty"`
	} `json:"topContributers"`
	TotalBounty  int64  `json:"totalBounty"`
	CategoryID   int    `json:"categoryId"`
	CategoryName string `json:"categoryName"`
	Title        string `json:"title"`
	Year         int    `json:"year"`
	Image        string `json:"image"`
	Description  string `json:"description"`
	MusicInfo    struct {
		Composers []string `json:"composers"`
		DJ        []string `json:"dj"`
		Artists   []struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"artists"`
		With []struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"with"`
		Conductor []string `json:"conductor"`
		RemixedBy []string `json:"remixedBy"`
		Producer  []string `json:"producer"`
	} `json"musicInfo"`
	CatalogueNumber string   `json:"catalogueNumber"`
	ReleaseType     int      `json:"releaseType"`
	ReleaseName     string   `json:"releaseName"`
	BitrateList     []string   `json:"bitrateList"`
	FormatList      []string   `json:"formatList"`
	MediaList       []string   `json:"mediaList"`
	LogCue          string   `json:"logCue"`
	IsFilled        bool     `json:"isFilled"`
	FillerID        int      `json:"fillerID"`
	FillerName      string   `json:"fillerName"`
	TorrentID       int      `json:"torrentID"`
	TimeFilled      string   `json:"timeFilled"`
	Tags            []string `json:"tags"`
	Comments        []struct {
		PostID       int    `json:"postId"`
		AuthorID     int    `json:"authorId"`
		Name         string `json:"name"`
		Donor        bool   `json:"donor"`
		Warned       bool   `json:"warned"`
		Enabled      bool   `json:"enabled"`
		Class        string `json:"class"`
		AddedTime    string `json:"addedTime"`
		Avatar       string `json:"avatar"`
		Comment      string `json:"comment"`
		EditUserID   int    `json:"editUserId"`
		EditUsername string `json:"editUsername"`
		EditedTime   string `json:"editedTime"`
	} `json:"comments"`
	CommentPage  int `json:"commentPage"`
	CommentPages int `json:"commentPages"`
}