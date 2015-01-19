package whatapi

type Mailbox struct {
	Status   string          `json:"status"`
	Error    string          `json:"error"`
	Response MailboxResponse `json:"response"`
}

type Conversation struct {
	Status   string               `json:"status"`
	Error    string               `json:"error"`
	Response ConversationResponse `json:"response"`
}

type MailboxResponse struct {
	CurrentPage int `json:"currentPage"`
	Pages       int `json:"pages"`
	Messages    []struct {
		ConvID        int    `json:"convId"`
		Subject       string `json:"subject"`
		Unread        bool   `json:"unread"`
		Sticky        bool   `json:"sticky"`
		ForwardedID   int    `json:"forwardedID"`
		ForwardedName string `json:"forwardedName"`
		SenderID      int    `json:"senderId"`
		Username      string `json:"username"`
		Donor         bool   `json:"donor"`
		Warned        bool   `json:"warned"`
		Enabled       bool   `json:"enabled"`
		Date          string `json:"date"`
	} `json:"messages"`
}

type ConversationResponse struct {
	ConvID   int    `json:"convId"`
	Subject  string `json:"subject"`
	Sticky   bool   `json:"sticky"`
	Messages []struct {
		MessageID  int    `json:"messageId"`
		SenderID   int    `json:"senderId"`
		SenderName string `json:"senderName"`
		SentDate   string `json:"sentDate"`
		BbBody     string `json:"bbBody"`
		Body       string `json:"body"`
	} `json:"messages"`
}
