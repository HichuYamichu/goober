package users

type loginPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type safeUser struct {
	Username string `json:"username"`
	Quota    int64  `json:"quota"`
	Admin    bool   `json:"admin"`
}

type loginResponce struct {
	Token string    `json:"token"`
	User  *safeUser `json:"user"`
}

type UserConfig struct {
	Quota int64 `json:"quota"`
}
