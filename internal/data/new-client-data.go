package data

type SendLogin struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	Token      string `json:"token"`
	RememberMe bool   `json:"rememberMe"`
}

type ReceiveLogin struct {
	Success         bool     `json:"success"`
	Permissions     []string `json:"permissions"`
	SessionID       string   `json:"sessionID"`
	RememberMeToken string   `json:"rememberMeToken"`
	UserInfo        struct {
		ID                 string `json:"ID"`
		Username           string `json:"Username"`
		EmailAddress       string `json:"EmailAddress"`
		IsTwoFactorEnabled bool   `json:"IsTwoFactorEnabled"`
		Disabled           bool   `json:"Disabled"`
		LastLogin          string `json:"LastLogin"`
		GravatarHash       string `json:"GravatarHash"`
		IsLDAPUser         bool   `json:"IsLDAPUser"`
	} `json:"userInfo"`
	Result int `json:"result"`
}
