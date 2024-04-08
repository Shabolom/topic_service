CREATE TABLE IF NOT EXISTS public.likes
(
    id uuid NOT NULL,
    when_crated TIMESTAMP WITH time zone,
    when_update TIMESTAMP WITH time zone,
    when_deleted TIMESTAMP WITH time zone,
    comment_id uuid NOT NULL UNIQUE REFERENCES public.comments(id),
    user_id uuid NOT NULL REFERENCES public.users(id),
    CONSTRAINT likes_unique UNIQUE (comment_id, user_id),
    CONSTRAINT likes_pk PRIMARY KEY (id)
    );

CREATE UNIQUE INDEX idx_likes
    ON public.likes (id);

ALTER TABLE IF EXISTS public.likes
    OWNER to postgres;