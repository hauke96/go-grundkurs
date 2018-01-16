// full spec at https://api.stackexchange.com/docs/info
type Item struct {
	Users     int `json:"total_users"`
	Questions int `json:"total_questions"`
	Answers   int `json:"total_answers"`
	Comments  int `json:"total_comments"`
}

type Info struct {
	Site  string
	Items []Item
}
