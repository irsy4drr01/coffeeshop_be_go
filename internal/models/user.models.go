package models

import "time"

type User struct {
	Uuid      string  `db:"uuid" json:"uuid" valid:"uuid~Uuid must be a valid UUID format"`
	Role      string  `db:"role" json:"role,omitempty" valid:"matches(^admin$|^user$)~Username length must be admin or user"`
	Username  string  `db:"username" json:"username" valid:"stringlength(5|100)~Username length must be between 5 and 100 characters"`
	Email     string  `db:"email" json:"email" valid:"email~Invalid email format"`
	Password  string  `db:"password,omitempty" json:"password,omitempty" valid:"stringlength(8|100)~Password length must be between 8 and 100 characters"`
	Image     string  `db:"image" json:"image,omitempty" valid:"-"`
	CreatedAt string  `db:"created_at,omitempty" json:"created_at,omitempty" valid:"-"`
	UpdatedAt *string `db:"updated_at,omitempty" json:"updated_at,omitempty" valid:"-"`
	IsDeleted bool    `db:"is_deleted,omitempty" json:"is_deleted,omitempty" valid:"-"`
}

type Users []User

type UserDB struct {
	ID         string     `db:"id"`
	Email      string     `db:"email"`
	Role       string     `db:"role"`
	FullName   string     `db:"fullname"`
	Phone      string     `db:"phone"`
	Address    string     `db:"address"`
	Image      string     `db:"image"`
	IsVerified bool       `db:"is_verified"`
	IsDeleted  bool       `db:"is_deleted"`
	CreatedAt  time.Time  `db:"created_at"`
	UpdatedAt  *time.Time `db:"updated_at"`
}

type UserProfileResponse struct {
	FullName   string `json:"fullname"`
	ProfileImg string `json:"profile_img"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	Address    string `json:"address"`
	CreatedAt  string `json:"created_at"`
}

type UserProfileUpdateDB struct {
	Fullname   string `db:"fullname"`
	ProfileImg string `db:"profile_img"`
	Email      string `db:"email"`
	Phone      string `db:"phone"`
	Address    string `db:"address"`
}

type UserProfileUpdateReq struct {
	Fullname   string `form:"fullname"`
	ProfileImg string `form:"profile_img"`
	Email      string `form:"email"`
	Phone      string `form:"phone"`
	Address    string `form:"address"`
}

type UserPasswordUpdateReq struct {
	Password string `json:"password"`
}
