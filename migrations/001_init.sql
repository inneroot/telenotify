-- Write your migrate up statements here

create table if not exists recipients
(
    id         serial primary key,
    recipient_id bigint,
    created_at timestamptz not null default now(),
    updated_at timestamptz not null default now()
);

create index if not exists recipient_id_idx on recipients (recipient_id);

---- create above / drop below ----

drop table recipients;
