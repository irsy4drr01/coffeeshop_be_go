CREATE TABLE public.product_stocks (
    product_id UUID NOT NULL REFERENCES public.products(id) ON DELETE CASCADE,
    size_id INT NOT NULL DEFAULT '4' REFERENCES public.sizes(id) ON DELETE CASCADE,
    stock INT NOT NULL DEFAULT 0,
    created_at TIMESTAMPTZ DEFAULT now(),
    updated_at TIMESTAMPTZ,    
    PRIMARY KEY (product_id, size_id)  -- composite PK
);