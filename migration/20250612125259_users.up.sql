CREATE TABLE IF NOT EXISTS public.users (
  id UUID DEFAULT gen_random_uuid() NOT NULL,
  email VARCHAR(255) NOT NULL UNIQUE,
  password VARCHAR(255) NOT NULL,
  role_id INT DEFAULT 3 NOT NULL,
  is_verified BOOLEAN DEFAULT FALSE,
  is_deleted BOOLEAN DEFAULT FALSE,
  created_at TIMESTAMPTZ DEFAULT now(),
  updated_at TIMESTAMPTZ,
  CONSTRAINT users_pk PRIMARY KEY (id)
);