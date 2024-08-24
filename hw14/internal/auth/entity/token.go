package entity

import "time"

type Token struct {
	UserId      int
	Token       string
	ExpiresIn   time.Time
	Fingerprint string
}
