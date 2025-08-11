CREATE TABLE public.product_images (
    id SERIAL PRIMARY KEY,
    product_id UUID NOT NULL REFERENCES public.products(id) ON DELETE CASCADE,
    slot_number INT NOT NULL CHECK (slot_number >= 1 AND slot_number <= 3),
    image_file TEXT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT now(),
    CONSTRAINT unique_product_slot UNIQUE (product_id, slot_number)
);