CREATE TABLE public.sizes (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL UNIQUE,
    additional_price FLOAT NOT NULL,  -- 1.0, 1.25, 1.5
    created_at TIMESTAMPTZ DEFAULT now(),
    updated_at TIMESTAMPTZ
);