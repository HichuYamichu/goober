package users

// User representrs user
type User struct {
	ID       int    `gorm:"type:serial;" json:"id"`
	Username string `json:"username"`
	Pass     string `json:"password"`
	Admin    bool   `json:"admin"`
	Quota    int64  `json:"quota"`
}
