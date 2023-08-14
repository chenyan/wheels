package dbx

import "time"

// BaseModel base model
type BaseModel struct {
	ID     int64     `db:"id" json:"id"`
	Status int       `db:"status" json:"-"`
	Ctime  time.Time `db:"ctime" json:"ctime"`
	Mtime  time.Time `db:"mtime" json:"-"`
}
