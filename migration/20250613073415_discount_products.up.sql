CREATE TABLE IF NOT EXISTS public.discount_products (
  product_id UUID NOT NULL,
  discount_id INT NOT NULL,  
  CONSTRAINT discount_product_pk PRIMARY KEY (product_id, discount_id),
  CONSTRAINT fk_discount_product_product FOREIGN KEY (product_id) REFERENCES products (id),
  CONSTRAINT fk_discount_product_discount FOREIGN KEY (discount_id) REFERENCES discounts (id)
);