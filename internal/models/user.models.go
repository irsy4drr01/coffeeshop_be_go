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
	Uuid      string  `db:"uuid" json:"uuid" valid:"uuid~Uuid must be a valid UUID format"`
	Username  string  `db:"username" json:"username" valid:"stringlength(5|100)~Username length must be between 5 and 100 characters"`
	Email     string  `db:"email" json:"email" valid:"email~Invalid email format"`
	Password  string  `db:"password,omitempty" json:"password,omitempty" valid:"stringlength(8|100)~Password length must be between 8 and 100 characters"`
	Image     string  `db:"image" json:"image" valid:"-"`
	CreatedAt string  `db:"created_at,omitempty" json:"created_at,omitempty" valid:"-"`
	UpdatedAt *string `db:"updated_at,omitempty" json:"updated_at,omitempty" valid:"-"`
	IsDeleted bool    `db:"is_deleted,omitempty" json:"is_deleted,omitempty" valid:"-"`
}

type Users []User
