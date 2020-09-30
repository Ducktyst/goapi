create table exchange_rate (
	id int serial primary key,
	currency_from varchar(50),
	currency_to varchar(50),
	cource int,
	update_at timestamp
);