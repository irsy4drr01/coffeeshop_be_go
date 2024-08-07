package models

var schemaUsers = `CREATE TABLE public.users (
	id serial4 NOT NULL,
	username varchar(255) NOT NULL,
	email varchar(255) NOT NULL,
	"password" varchar(255) NOT NULL,
	created_at timestamptz DEFAULT now() NULL,
	updated_at timestamptz NULL,
	"uuid" uuid DEFAULT gen_random_uuid() NULL,
	is_deleted bool DEFAULT false NULL,
	CONSTRAINT unique_email UNIQUE (email),
	CONSTRAINT users_pk PRIMARY KEY (id)
);`

var _ = schemaUsers

type User struct {
	Uuid      string  `db:"uuid" json:"uuid"`
	Username  string  `db:"username" json:"username"`
	Email     string  `db:"email" json:"email"`
	Password  string  `db:"password" json:"password"`
	CreatedAt string  `db:"created_at" json:"created_at"`
	UpdatedAt *string `db:"updated_at" json:"updated_at"`
	IsDeleted bool    `db:"is_deleted" json:"is_deleted"`
}

type UserDetail struct {
	Uuid      string `db:"uuid" json:"uuid"`
	Username  string `db:"username" json:"username"`
	Email     string `db:"email" json:"email"`
	Password  string `db:"password" json:"password"`
	CreatedAt string `db:"created_at" json:"created_at"`
}

type CreateUserResponse struct {
	Uuid      string `db:"uuid" json:"uuid"`
	Username  string `db:"username" json:"username"`
	Email     string `db:"email" json:"email"`
	Password  string `db:"password" json:"password"`
	CreatedAt string `db:"created_at" json:"created_at"`
}

type UpdateUserResponse struct {
	Uuid      string  `db:"uuid" json:"uuid"`
	Username  string  `db:"username" json:"username"`
	Email     string  `db:"email" json:"email"`
	Password  string  `db:"password" json:"password"`
	UpdatedAt *string `db:"updated_at" json:"updated_at"`
}

type DeleteUserResponse struct {
	Uuid      string `db:"uuid" json:"uuid"`
	Username  string `db:"username" json:"username"`
	Email     string `db:"email" json:"email"`
	IsDeleted bool   `db:"is_deleted" json:"is_deleted"`
}

type Users []User
