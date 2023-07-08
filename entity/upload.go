package entity

type Upload struct {
	Time     int64  `json:"time"`
	SelfId   int64  `json:"self_id"`
	PostType string `json:"post_type"`
}
