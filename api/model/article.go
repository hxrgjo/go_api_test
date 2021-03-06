package model

import "time"

type Article struct {
	ID        int64 `xorm:"'id' pk autoincr"`
	UserID    int64 `xorm:"'user_id' pk autoincr"`
	Name      string
	Content   string
	LikeCount int
	CreatedAt time.Time `xorm:"created"`
	UpdatedAt time.Time `xorm:"updated"`
}

func (Article) TableName() string {
	return "articles"
}
