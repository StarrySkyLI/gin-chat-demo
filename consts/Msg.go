package consts

type Msg struct {
	Datebase string `json:"datebase"`
	Id       string `json:"id"`
	Content  string `json:"content"`
	Read     uint   `json:"read"`
	Expire   int64  `json:"expire"`
}
