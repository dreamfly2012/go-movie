package model

//Category 类型
type Category struct {
	ID   int
	Name string `form:"name" json:"name"`
	URL  string `form:"url" json:"url"`
}
