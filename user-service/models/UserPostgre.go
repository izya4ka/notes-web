package models

type UserPostgres struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Notes     []Note `json:"notes" gorm:"foreignKey:UserID;references:ID"`
	CreatedAt int64  `gorm:"autoCreateTime:milli" json:"created_at"`
	UpdatedAt int64  `gorm:"autoUpdateTime:milli" json:"updated_at"`
}
