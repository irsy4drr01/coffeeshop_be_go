CREATE TABLE IF NOT EXISTS public.delivery_methods (
  id serial PRIMARY KEY,
  name varchar UNIQUE NOT NULL,
  fee INT NOT NULL,
  created_at timestamptz DEFAULT now(),
  updated_at timestamptz
);