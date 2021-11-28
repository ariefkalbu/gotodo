package model

import "time"

type Session struct {
	Username  string
	Name      string
	Email     string
	Logindate time.Time
}
