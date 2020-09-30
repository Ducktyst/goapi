create table if not exists exchange_rate
(
    id            serial primary key,
    currency_from varchar(50),
    currency_to   varchar(50),
    cource        int,
    updated_at    timestamp
);