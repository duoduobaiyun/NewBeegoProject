package models

//解析结构体
type Info struct {
	User   string  `json:"user"`
	Password string `json:"password"`
	School string `json:"school"`
	Class  string `json:"class"`
	Nick   string `json:"nick"`
	Sex    string `json:"sex"`
	Hobby  string `json:"hobby"`
	Color  string `json:"color"`
}
