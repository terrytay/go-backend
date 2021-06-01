package entities

type User struct {
	Username string
	Password string
}

func (u User) toDTO() *UserDTO {
	return &UserDTO{
		Username: u.Username,
		Password: u.Password,
	}
}

type UserDTO struct {
	Username string
	Password string
}

func (u UserDTO) toEntity() *User {
	return &User{
		Username: u.Username,
		Password: u.Password,
	}
}
