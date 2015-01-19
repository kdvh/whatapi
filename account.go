package whatapi

type Account struct {
	Status   string          `json:"status"`
	Error    string          `json:"error"`
	Response AccountResponse `json:"response"`
}

type AccountResponse struct {
	Username      string `json:"username"`
	ID            int    `json:"id"`
	AuthKey       string `json:"authKey"`
	PassKey       string `json:"passKey"`
	Notifications struct {
		Messages       int  `json:"messages"`
		Notifications  int  `json:"notifications"`
		NewAnnouncment bool `json:"newAnnouncment"`
		NewBlog        bool `json:"newBlog"`
	} `json:"notifications"`
	UserStats struct {
		Uploaded      int64   `json:"uploaded"`
		Downloaded    int64   `json:"downloaded"`
		Ratio         float64 `json:"ratio"`
		RequiredRatio float64 `json:"requiredRatio"`
		Class         string  `json:"class"`
	} `json:"userstats"`
}
