CREATE TABLE IF NOT EXISTS public.categories (
    id serial4 NOT NULL,
    category_name varchar(255) NOT NULL,
    created_at timestamptz DEFAULT now() NOT NULL,
    updated_at timestamptz NULL,
    CONSTRAINT category_name_unique UNIQUE (category_name),
    CONSTRAINT categories_pk PRIMARY KEY (id)
);