package db

import "time"

type Parts struct {
	ID        int    `migu:"autoincrement"`
	NAME      string `migu:"size:32"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type PartsImages struct {
	ID        int `migu:"autoincrement"`
	PartsID   int
	NAME      string `migu:"size:32"`
	CreatedAt time.Time
	UpdatedAt time.Time
}