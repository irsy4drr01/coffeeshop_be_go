CREATE TABLE public.products (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    product_name VARCHAR(255) NOT NULL UNIQUE,
    category_id INT NOT NULL REFERENCES public.categories(id),
    price NUMERIC NOT NULL,  -- base price
    description TEXT default '',
    total_sold INT NOT NULL DEFAULT 0,
    total_likes INT NOT NULL DEFAULT 0,
    is_deleted BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMPTZ DEFAULT now(),
    updated_at TIMESTAMPTZ
);