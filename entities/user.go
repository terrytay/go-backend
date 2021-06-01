package entities

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (u User) ToDTO() *UserDTO {
	return &UserDTO{
		Username: u.Username,
		Password: u.Password,
	}
}

type UserDTO struct {
	tableName struct{} `pg:"users"`

	Username string
	Password string
}

func (u UserDTO) ToEntity() *User {
	return &User{
		Username: u.Username,
		Password: u.Password,
	}
}
