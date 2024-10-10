CREATE TABLE IF NOT EXISTS public.product (
	id serial4 NOT NULL,
	product_name varchar(255) NOT NULL,
	price int4 NOT NULL,
	category varchar(255) NOT NULL,
	description text NULL,
	created_at timestamptz DEFAULT now() NULL,
	updated_at timestamptz NULL,
	"uuid" uuid DEFAULT gen_random_uuid() NULL,
	is_deleted bool DEFAULT false NULL,
	image varchar(255) DEFAULT ''::character varying NOT NULL,
	CONSTRAINT product_name_unique UNIQUE (product_name),
	CONSTRAINT product_pk PRIMARY KEY (id)
);