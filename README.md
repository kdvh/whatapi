whatapi
=======

A Go wrapper for the What.CD [AJAX API](https://github.com/WhatCD/Gazelle/wiki/JSON-API-Documentation)

Example
-------
```Go
        wcd := whatapi.NewSite("https://what.cd/")
        wcd.Login("username", "password")

        account := wcd.GetAccount()
        fmt.Println(account.Username)

        user := wcd.GetUser(100)
        fmt.Println(user.Username)

        mailboxParams := url.Values{}
        mailboxParams.Set("type", "sentbox")
        mailbox := wcd.GetMailbox(mailboxParams)
        conversation := wcd.GetConversation(mailbox.Messages[0].ConvID)
        fmt.Println(conversation.Messages[0].Body)

        torrentParams := url.Values{}
        torrent := wcd.GetTorrent(31929409, torrentParams)
        fmt.Println(wcd.CreateDownloadURL(torrent.Torrent.ID))

```
