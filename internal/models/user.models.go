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
	Password  string  `db:"password,omitempty" json:"password,omitempty"`
	CreatedAt string  `db:"created_at,omitempty" json:"created_at,omitempty"`
	UpdatedAt *string `db:"updated_at,omitempty" json:"updated_at,omitempty"`
	IsDeleted bool    `db:"is_deleted,omitempty" json:"is_deleted,omitempty"`
}

type Users []User
