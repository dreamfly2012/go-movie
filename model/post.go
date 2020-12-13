package model

import "time"

//Post Model for orm
type Post struct {
	ID      int
	CID     int
	Title   string
	Desc    string
	Content string
	Created time.Time
	Updated time.Time
	Views   int8
	Status  int8
}
