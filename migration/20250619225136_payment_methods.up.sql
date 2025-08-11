CREATE TABLE public.payment_methods (
	id serial4 NOT NULL,
	name varchar NOT NULL,
	created_at timestamptz DEFAULT now() NULL,
	updated_at timestamptz NULL,
	CONSTRAINT payment_methods_name_key UNIQUE (name),
	CONSTRAINT payment_methods_pkey PRIMARY KEY (id)
);