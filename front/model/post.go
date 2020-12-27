package model

import "time"

//Post Model for orm
type Post struct {
	ID      int       `form:"id" json:"id"`
	Cid     int       `form:"cid" json:"cid"`
	Title   string    `form:"title" json:"title"`
	Desc    string    `form:"desc" json:"desc"`
	Content string    `form:"content" json:"content"`
	Created time.Time `form:"created" json:"created"`
	Updated time.Time `form:"updated" json:"updated"`
	Views   int8      `form:"views" json:"views"`
	Status  int8      `form:"status" json:"status"`
}
