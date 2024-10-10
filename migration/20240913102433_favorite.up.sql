CREATE TABLE IF NOT EXISTS public.favorite (
    user_id int4 NOT NULL,
    product_id int4 NOT NULL,
    CONSTRAINT favorite_pk PRIMARY KEY (user_id, product_id),
    CONSTRAINT fk_product FOREIGN KEY (product_id) REFERENCES public.product(id) ON DELETE CASCADE,
    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE CASCADE
);
