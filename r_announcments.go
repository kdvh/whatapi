package whatapi

type Announcements struct {
	Announcements []struct {
		NewsID   int    `json:"newsId"`
		Title    string `json:"title"`
		BbBody   string `json:"bbBody"`
		Body     string `json:"body"`
		NewsTime string `json:"newsTime"`
	} `json:"announcements"`
	BlogPosts []struct {
		BlogID   int    `json:"blogId"`
		Author   string `json:"author"`
		Title    string `json:"title"`
		BbBody   string `json:"bbBody"`
		Body     string `json:"body"`
		BlogTime string `json:"blogTime"`
		ThreadID int    `json:"threadId"`
	} `json:"blogPosts"`
}
