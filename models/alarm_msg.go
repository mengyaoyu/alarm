package models

import "time"

type AlarmMsg struct {
	Id          int64
	Msg         string
	Status      int
	AccessToken string
	CreateTime  time.Time
	UpdateTime  time.Time
}
