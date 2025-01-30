package user

type User struct {
	ID       int64  `json:"id" gorm:"primary_key"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (u *User) HashPassword() error {
	// TODO: 这里是伪代码，实际应用中应该使用 bcrypt 等安全的哈希算法
	u.Password = "hashed:" + u.Password
	return nil
}

func (u *User) Authenticate(password string) bool {
	password = "hashed:" + password
	if password == u.Password {
		return true
	}
	return false
}

// ShowStatus 展示用户状态, 包含用户的功能列表、权限，绑定的设备等
func (u *User) ShowStatus() string {
	return "I am OK."
}
