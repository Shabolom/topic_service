CREATE TABLE IF NOT EXISTS public.like_dizlike_count
(
    id uuid NOT NULL,
    comment_id uuid NOT NULL REFERENCES public.comments(id),
    likes int NOT NULL,
    diz_likes int NOT NULL,
    CONSTRAINT like_dizlike_count_pk PRIMARY KEY (id)
    );

CREATE UNIQUE INDEX idx_like_dizlike_count
    ON public.like_dizlike_count (id);

ALTER TABLE IF EXISTS public.like_dizlike_count
    OWNER to postgres;