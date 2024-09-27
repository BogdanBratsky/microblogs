package models

import "time"

type User struct {
	UserId         int       `json:"id"`
	UserName       string    `json:"name"`
	UserEmail      string    `json:"email"`
	PasswordHash   string    `json:"password"`
	AvatarUrl      string    `json:"avatar"`
	CreatedAt      time.Time `json:"createdAt"`
	IsAdmin        bool      `json:"admin"`
	IsActive       bool      `json:"active"`
	FollowingCount int       `json:"following"`
	FollowersCount int       `json:"followers"`
}
