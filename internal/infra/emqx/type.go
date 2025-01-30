package emqx

type CommonModel struct {
	TID       string `json:"tid"`
	BID       string `json:"bid"`
	Method    string `json:"method"`
	Timestamp int64  `json:"timestamp"`
}
