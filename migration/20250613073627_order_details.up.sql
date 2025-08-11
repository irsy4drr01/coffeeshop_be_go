CREATE TABLE IF NOT EXISTS public.order_details (
  order_id VARCHAR NOT NULL,
  product_id uuid NOT NULL,
  qty int NOT NULL,
  size_id INT NOT NULL DEFAULT 4,
  is_iced BOOLEAN DEFAULT FALSE,  
  is_discount BOOLEAN DEFAULT FALSE,
  discount_name VARCHAR NOT NULL DEFAULT '',  
  base_price NUMERIC NOT NULL,
  final_price NUMERIC NOT NULL,
  sub_total NUMERIC NOT NULL,
  CONSTRAINT products_orders_pk PRIMARY KEY (order_id, product_id, size_id, is_iced),
  CONSTRAINT fk_products_orders_order FOREIGN KEY (order_id) REFERENCES orders (id),
  CONSTRAINT fk_products_orders_product FOREIGN KEY (product_id) REFERENCES products (id),
  CONSTRAINT fk_products_orders_size FOREIGN KEY (size_id) REFERENCES sizes (id)
);