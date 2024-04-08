CREATE TABLE IF NOT EXISTS public.comments
(
    id uuid NOT NULL,
    when_crated TIMESTAMP WITH time zone,
    when_update TIMESTAMP WITH time zone,
    when_deleted TIMESTAMP WITH time zone,
    user_massage text COLLATE pg_catalog."default",
    user_file_path text COLLATE pg_catalog."default",
    user_id uuid NOT NULL REFERENCES public.users(id),
    topic_id uuid NOT NULL REFERENCES public.topics(id),
    CONSTRAINT comments_pk PRIMARY KEY (id)
    );

CREATE UNIQUE INDEX idx_comments
    ON public.comments (id);

ALTER TABLE IF EXISTS public.comments
    OWNER to postgres;