package entities

import (
	//"github.com/go-pg/pg/v10/orm"
)

type User struct {
	Username     string
	PasswordHash string
}

func (u User) toDTO() *UserDTO {
	return &UserDTO{
		Username:     u.Username,
		PasswordHash: u.PasswordHash,
	}
}

type UserDTO struct {
	Username     string
	PasswordHash string
}

func (u UserDTO) toEntity() *User {
	return &User{
		Username:     u.Username,
		PasswordHash: u.PasswordHash,
	}
}

