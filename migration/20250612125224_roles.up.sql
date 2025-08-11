CREATE TABLE public.roles (
    id SERIAL PRIMARY KEY,
    role VARCHAR(50) NOT NULL UNIQUE,
    created_at TIMESTAMPTZ DEFAULT now(),
    updated_at TIMESTAMPTZ
);