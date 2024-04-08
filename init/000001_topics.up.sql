CREATE TABLE IF NOT EXISTS public.topics
(
    id uuid NOT NULL,
    when_crated TIMESTAMP WITH time zone,
    when_update TIMESTAMP WITH time zone,
    when_deleted TIMESTAMP WITH time zone,
    topic_info text COLLATE pg_catalog."default",
    topic_name text COLLATE pg_catalog."default",
    topic_file text COLLATE pg_catalog."default",
    CONSTRAINT topics_pk PRIMARY KEY (id)
    );

CREATE UNIQUE INDEX idx_topics
    ON public.topics (id);

ALTER TABLE IF EXISTS public.topics
    OWNER to postgres;