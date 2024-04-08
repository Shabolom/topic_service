CREATE TABLE IF NOT EXISTS public.users
(
    id uuid NOT NULL,
    when_created TIMESTAMP WITH time zone,
    when_update TIMESTAMP WITH time zone,
    when_deleted TIMESTAMP WITH time zone,
    user_name text COLLATE pg_catalog."default",
    CONSTRAINT user_pk PRIMARY KEY (id)
    );

CREATE UNIQUE INDEX idx_users
    ON public.users (id);

ALTER TABLE IF EXISTS public.users
    OWNER to postgres;