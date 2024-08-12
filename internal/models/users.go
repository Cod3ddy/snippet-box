package models

import (
	"database/sql"
	"time"
)

type User struct {
	ID             int
	Name           string
	Email          string
	HashedPassword []byte
	Created        time.Time
}

type UserModel struct{
	DB *sql.DB
}

// Add new user
func (m *UserModel) Insert(name, email, password string) error{
	return nil
}


func (m *UserModel) Authenicate(email, password string) (int, error){
	return 0, nil
}

func (m *UserModel) Exists (userID int) (bool, error){
	return false, nil
}