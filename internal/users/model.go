package users

// User representrs user
type User struct {
	ID       int    `gorm:"type:serial;" json:"id"`
	Username string `gorm:"unique;not null;" json:"username"`
	Pass     string `gorm:"not null;" json:"-"`
	Admin    bool   `gorm:"not null;" json:"admin"`
	Active   bool   `gorm:"not null;" json:"active"`
	Quota    int64  `gorm:"not null;" json:"quota"`
}
