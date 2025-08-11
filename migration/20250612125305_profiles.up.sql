CREATE TABLE public.profiles (
    id UUID PRIMARY KEY REFERENCES public.users(id) ON DELETE CASCADE,
    fullname VARCHAR(255) NOT NULL DEFAULT '',
    phone VARCHAR(20) NOT NULL DEFAULT '',
    address VARCHAR(255) NOT NULL DEFAULT '',
    image TEXT NULL,
    created_at TIMESTAMPTZ DEFAULT now(),
    updated_at TIMESTAMPTZ
);

CREATE UNIQUE INDEX unique_phone_not_empty
ON public.profiles (phone)
WHERE phone <> '';