whatapi
=======

A Go wrapper for the What.CD [JSON API](https://github.com/WhatCD/Gazelle/wiki/JSON-API-Documentation)


Install
-------

```
go get "github.com/kdvh/whatapi"
```

Example
-------
```Go
	wcd, err := whatapi.NewWhatAPI("https://what.cd/")
	if err != nil {
		log.Fatal(err)
	}
	
	err = wcd.Login("username", "password")
	if err != nil {
		log.Fatal(err)
	}
	
	mailboxParams := url.Values{}
	mailboxParams.Set("type", "sentbox")
	mailbox, err := wcd.GetMailbox(mailboxParams)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(mailbox)

	conversation, err := wcd.GetConversation(mailbox.Messages[0].ConvID)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(conversation.Messages[0].Body)

	torrentSearchParams := url.Values{}
	torrentSearch, err := wcd.SearchTorrents("Tool", torrentSearchParams)
	if err != nil {
		log.Fatal(err)
	}

	downloadURL, err := wcd.CreateDownloadURL(torrentSearch.Results[0].Torrents[0].TorrentID)
	if err != nil {
		log.Fatal(downloadURL)
	}
	log.Println(downloadURL)
```
