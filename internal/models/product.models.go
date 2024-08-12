package models

var schemaProducts = `CREATE TABLE public.product (
	id serial4 NOT NULL,
	product_name varchar(255) NOT NULL,
	price int4 NOT NULL,
	category varchar(255) NOT NULL,
	description text NULL,
	created_at timestamptz DEFAULT now() NULL,
	updated_at timestamptz NULL,
	"uuid" uuid DEFAULT gen_random_uuid() NULL,
	is_deleted bool DEFAULT false NULL,
	CONSTRAINT product_name_unique UNIQUE (product_name),
	CONSTRAINT product_pk PRIMARY KEY (id)
);`

var _ = schemaProducts

type Product struct {
	ID          int     `db:"id" json:"id,omitempty" valid:"-"`
	Uuid        string  `db:"uuid" json:"uuid,omitempty" valid:"uuid~Uuid must be a valid UUID format"`
	ProductName string  `db:"product_name" json:"product_name,omitempty" valid:"stringlength(3|100)~ProductName length must be between 3 and 100 characters"`
	Price       int     `db:"price" json:"price,omitempty" valid:"type(int)~Price must be a interger"`
	Category    string  `db:"category" json:"category,omitempty" valid:"type(string)~Price must be a string"`
	Description *string `db:"description,omitempty" json:"description,omitempty" valid:"stringlength(0|100)~Description length max 100 characters"`
	CreatedAt   string  `db:"created_at,omitempty" json:"created_at,omitempty" valid:"-"`
	UpdatedAt   *string `db:"updated_at,omitempty" json:"updated_at,omitempty" valid:"-"`
	IsDeleted   bool    `db:"is_deleted,omitempty" json:"is_deleted,omitempty" valid:"-"`
}

type Products []Product
