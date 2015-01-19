package whatapi

type TorrentGroup struct {
	Status   string               `json:"status"`
	Error    string               `json:"error"`
	Response TorrentGroupResponse `json:"response"`
}

type TorrentGroupResponse struct {
	Group    GroupType `json:"group"`
    Torrent []TorrentType `json:"torrents"`
}
