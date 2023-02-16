package entity

import (
	"time"

	"vikishptra/domain_pensfess/model/vo"
)

type User struct {
	ID               vo.UserID `bson:"_id" json:"id"`
	Username         string    `bson:"username" json:"username"`
	Password         string    `bson:"password" json:"password"`
	ActiveUser       bool      `bson:"active_user" json:"active_user"`
	Email            string    `bson:"email" json:"email"`
	VerificationCode string
	Token            string
	CreatedAt        time.Time `bson:"created" json:"created"`
	UpdatedAt        time.Time `bson:"updated" json:"updated"`
}

type UserCreateRequest struct {
	RandomString    string    `json:"id"`
	Now             time.Time `json:"time"`
	Name            string    `bson:"name" json:"name" form:"name"`
	Email           string    `bson:"email" json:"email" form:"email"`
	Password        string    `form:"password" json:"password"`
	ConfirmPassword string    `form:"confirm_password" json:"confirm_password"`
}

func NewUser(req UserCreateRequest) (*User, error) {

	id, err := vo.NewUserID(req.RandomString, req.Now)
	if err != nil {
		return nil, err
	}

	var obj User
	obj.ID = id
	obj.Email = req.Email
	obj.Password = req.Password
	obj.CreatedAt = req.Now
	obj.UpdatedAt = req.Now

	return &obj, nil
}

type UserUpdateRequest struct {
}

func (r *User) Update(req UserUpdateRequest) error {

	return nil
}
