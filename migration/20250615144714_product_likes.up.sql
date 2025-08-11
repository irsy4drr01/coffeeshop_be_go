CREATE TABLE IF NOT EXISTS public.product_likes (
  user_id uuid NOT NULL,
  product_id uuid NOT NULL,
  created_at timestamptz DEFAULT now() NOT NULL,
  updated_at timestamptz,
  CONSTRAINT product_like_pk PRIMARY KEY (user_id, product_id),
  CONSTRAINT fk_product_like_product FOREIGN KEY (product_id) REFERENCES products (id),
  CONSTRAINT fk_product_like_user FOREIGN KEY (user_id) REFERENCES users (id)
);