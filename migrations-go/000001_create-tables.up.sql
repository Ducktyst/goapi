create table if not exists exchange_rate
(
    id            serial primary key,
    currency_from varchar(50),
    currency_to   varchar(50),
    cource        numeric(4, 2) default -1,
    updated_at    timestamp with time zone default now(),
    unique (currency_from, currency_to)
);

insert into exchange_rate (currency_from, currency_to, cource, updated_at)
values ('USD', 'RUB', 78.31, NOW());
insert into exchange_rate (currency_from, currency_to, cource, updated_at)
values ('RUB', 'USD', 0.013, NOW());