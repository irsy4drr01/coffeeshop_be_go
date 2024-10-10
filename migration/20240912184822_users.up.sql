CREATE TABLE IF NOT EXISTS public.users (
	id serial4 NOT NULL,
	username varchar(255) NOT NULL,
	email varchar(255) NOT NULL,
	"password" varchar(255) NOT NULL,
	created_at timestamptz DEFAULT now() NULL,
	updated_at timestamptz NULL,
	"uuid" uuid DEFAULT gen_random_uuid() NULL,
	is_deleted bool DEFAULT false NULL,
	"role" varchar(255) DEFAULT 'user'::character varying NOT NULL,
	image varchar(255) DEFAULT ''::character varying NOT NULL,
	CONSTRAINT unique_email UNIQUE (email),
	CONSTRAINT users_pk PRIMARY KEY (id)
);