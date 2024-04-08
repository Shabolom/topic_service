CREATE TABLE IF NOT EXISTS public.users_topics
(
    id uuid NOT NULL,
    when_crated TIMESTAMP WITH time zone,
    when_deleted TIMESTAMP WITH time zone,
    user_id uuid NOT NULL REFERENCES  public.users(id),
    topic_id uuid NOT NULL REFERENCES  public.topics(id),
    CONSTRAINT users_topics_unique UNIQUE (user_id, topic_id),
    CONSTRAINT users_topics_pk PRIMARY KEY (id)
    );

CREATE UNIQUE INDEX idx_users_topics
    ON public.users_topics (id);

ALTER TABLE IF EXISTS public.users_topics
    OWNER to postgres;