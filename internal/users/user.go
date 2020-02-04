package users

type User struct {
	ID       int `gorm:"type:serial;"`
	Username string
	Pass     string
	Admin    bool
	Quota    int64
}
