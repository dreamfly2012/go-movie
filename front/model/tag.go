package model

//Tag 类型
type Tag struct {
	ID   int
	Name string `form:"name" json:"name"`
	URL  string `form:"url" json:"url"`
}
