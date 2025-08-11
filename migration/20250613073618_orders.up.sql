CREATE table public.orders (
	id VARCHAR PRIMARY KEY UNIQUE,
	user_id uuid NOT NULL,
	fullname VARCHAR NOT NULL,
	address VARCHAR NOT NULL,
	phone VARCHAR(20) NOT NULL,
	payment_method VARCHAR NOT NULL,
	total_purchase NUMERIC NOT NULL,
	delivery_method VARCHAR NOT NULL,
	delivery_method_fee NUMERIC NOT NULL,
	tax NUMERIC NOT NULL,
	status_id INT NOT NULL,
	total_amount NUMERIC NOT NULL,
	created_at timestamptz DEFAULT now(),
	updated_at timestamptz,
	CONSTRAINT fk_order_statuses FOREIGN KEY (status_id) REFERENCES statuses (id)
);