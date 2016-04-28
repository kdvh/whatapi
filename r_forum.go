package whatapi

type Categories struct {
	Categories []struct {
		CategoryID   int    `json:"categoryId"`
		CategoryName string `json:"categoryName"`
		Forums       []struct {
			ForumID            int      `json:"forumId"`
			ForumName          string   `json:"forumName"`
			ForumDescription   string   `json:"forumDescription"`
			NumTopics          int      `json:"numTopics"`
			NumPosts           int      `json:"numPosts"`
			LastPostID         int      `json:"lastPostId"`
			LastAuthorID       int      `json:"lastAuthorId"`
			LastPostAuthorName string   `json:"lastPostAuthorName"`
			LastTopicID        int      `json:"lastTopicId"`
			LastTime           string   `json:"lastTime"`
			SpecificRules      []string `json:"specificRules"`
			LastTopic          string   `json:"lastTopic"`
			Read               bool     `json:"read"`
			Locked             bool     `json:"locked"`
			Sticky             bool     `json:"sticky"`
		} `json:"forums"`
	} `json:"categories"`
}

type Forum struct {
	ForumName     string   `json:"forumName"`
	SpecificRules []struct{
        ThreadID int `json:"threadID"`
        Thread string `json:"thread"`
    } `json:"specificRules"`
	CurrentPage   int      `json:"currentPage"`
	Pages         int      `json:"pages"`
	Threads       []struct {
		TopicID        int    `json:"topicId"`
		Title          string `json:"title"`
		AuthorID       int    `json:"authorId"`
		AuthorName     string `json:"authorName"`
		Locked         bool   `json:"locked"`
		Sticky         bool   `json:"sticky"`
		PostCount      int    `json:"postCount"`
		LastID         int    `json:"lastID"`
		LastTime       string `json:"lastTime"`
		LastAuthorId   int    `json:"lastAuthorId"`
		LastAuthorName string `json:"lastAuthorNam"`
		LastReadPage   int    `json:"lastReadPage"`
		LastReadPostID int    `json:"lastReadPostId"`
		Read           bool   `json:"read"`
	} `json:"threads"`
}

type Thread struct {
	ForumID     int `json:"forumId"`
	ForumName   string `json:"forumName"`
	ThreadID    int    `json:"threadId"`
	ThreadTitle string `json:"threadTitle"`
	Subscribed  bool   `json:"subscribed"`
	Locked      bool   `json:"locked"`
	Sticky      bool   `json:"sticky"`
	CurrentPage int    `json:"currentPage"`
	Pages       int    `json:"pages"`
	Poll        struct {
		Closed     bool   `json:"closed"`
		Featured   string `json:"featured"`
		Question   string `json:"question"`
		MaxVotes   int    `json:"maxVotes"`
		TotalVotes int    `json:"totalVotes"`
		Voted      bool   `json:"voted"`
		Answers    []struct {
			Answer  string  `json:"answer"`
			Ratio   float64 `json:"ratio"`
			Percent float64 `json:"percent"`
		} `json:"answers"`
	} `json:"poll"`
	Posts []struct {
		PostID         int    `json:"postId"`
		AddedTime      string `json:"addedTime"`
		BbBody         string `json:"bbBody"`
		Body           string `json:"body"`
		EditedUserID   int    `json:"editedUserId"`
		EditedTime     string `json:"editedTime"`
		EditedUsername string `json:"editedUsername"`
		Author         struct {
			AuthorID   int      `json:"authorId"`
			AuthorName string   `json:"authorName"`
			Paranoia   []string `json:"paranoia"`
			Artist     bool     `json:"artist"`
			Donor      bool     `json:"donor"`
			Warned     bool     `json:"warned"`
			Avatar     string   `json:"avatar"`
			Enabled    bool     `json:"enabled"`
			UserTitle  string   `json:"userTitle"`
		} `json:"author"`
	} `json:"posts"`
}

type Subscriptions struct {
	Threads []struct {
		ForumID     int    `json:"forumId"`
		ForumName   string `json:"forumName"`
		ThreadID    int    `json:"threadId"`
		ThreadTitle string `json:"threadTitle"`
		PostID      int    `json:"postId"`
		Locked      bool   `json:"locked"`
		New         bool   `json:"new"`
	} `json:"threads"`
}