CREATE TABLE IF NOT EXISTS public.tax (
  id serial PRIMARY KEY NOT NULL UNIQUE,
  tax_value float NOT NULL,
  created_at timestamptz DEFAULT now(),
  updated_at timestamptz
);