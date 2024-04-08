CREATE TABLE IF NOT EXISTS public.diz_likes
(
    id uuid NOT NULL,
    when_crated TIMESTAMP WITH time zone,
    when_update TIMESTAMP WITH time zone,
    when_deleted TIMESTAMP WITH time zone,
    comment_id uuid NOT NULL REFERENCES public.comments(id),
    user_id uuid NOT NULL REFERENCES public.users(id),
    CONSTRAINT diz_likes_unique UNIQUE (comment_id, user_id),
    CONSTRAINT diz_likes_pk PRIMARY KEY (id)
    );

CREATE UNIQUE INDEX idx_diz_likes
    ON public.diz_likes (id);

ALTER TABLE IF EXISTS public.diz_likes
    OWNER to postgres;