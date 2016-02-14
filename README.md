whatapi
=======

A Go wrapper for the What.CD [JSON API](https://github.com/WhatCD/Gazelle/wiki/JSON-API-Documentation)

Example
-------
```Go
        wcd, _ := whatapi.NewSite("https://what.cd/")
        err := wcd.Login("username", "password")

        if err != nil {
            log.Fatal(err)
        }

        // Get account info
        account, _ := wcd.GetAccount()
        fmt.Println(account.Username)

        // Get PMs
        mailboxParams := url.Values{}
        mailboxParams.Set("type", "sentbox")
        mailbox, _ := wcd.GetMailbox(mailboxParams)
        conversation, _ := wcd.GetConversation(mailbox.Messages[0].ConvID)
        fmt.Println(conversation.Messages[0].Body)

        // Get torrent by ID and make url
        torrentParams := url.Values{}
        torrent := wcd.GetTorrent(31929409, torrentParams)
        fmt.Println(wcd.CreateDownloadURL(torrent.Torrent.ID))

```
