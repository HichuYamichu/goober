package users

// User representrs user
type User struct {
	ID       int    `gorm:"type:serial;" json:"id"`
	Username string `json:"username"`
	Pass     string `json:"-"`
	Admin    bool   `json:"admin"`
	Active   bool   `json:"active"`
	Quota    int64  `json:"quota"`
}
