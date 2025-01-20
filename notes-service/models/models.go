package models

type Note struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	CreatedAt   int64  `gorm:"autoCreateTime:milli" json:"created_at"`
	UpdatedAt   int64  `gorm:"autoUpdateTime:milli" json:"updated_at"`
}

type UserPostgre struct {
	ID        uint `gorm:"primaryKey" json:"id"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Notes     []Note `json:"notes"`
	CreatedAt int64 `gorm:"autoCreateTime:milli" json:"created_at"`
	UpdatedAt int64 `gorm:"autoUpdateTime:milli" json:"updated_at"`
}

type UserRedis struct {
	Username string
	Session  string
}
