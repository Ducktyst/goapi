package database

import (
	"context"
	"time"

	"github.com/ducktyst/goapi/internal/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Database interface {
	GetExchangeRates(ctx context.Context) (result []ExchangeRate, err error)
	GetExchangeRate(ctx context.Context, from, to string) (result []ExchangeRate, err error)
	CreateExchangeRate(ctx context.Context, from, to string) error
	UpdateExchangeRate(ctx context.Context, er ExchangeRate) error
	Close()
}

type DB struct {
	Conn *sqlx.DB
}

func New(cfg config.DatabaseConfig) (*DB, error) {
	conn, err := sqlx.Connect("postgres", cfg.ConnectionString)
	if err != nil {
		return nil, err
	}
	return &DB{
		Conn: conn,
	}, nil
}

type ExchangeRate struct {
	tableName    struct{}  `sql:"exchange_rate"`
	ID           int       `db:"id"`
	CurrencyFrom string    `db:"currency_from"`
	CurrencyTo   string    `db:"currency_to"`
	Cource       float64   `db:"cource"` // todo: decimal
	UpdatedAt    time.Time `db:"updated_at"`
}

func (d *DB) GetExchangeRates(ctx context.Context) (result []ExchangeRate, err error) {
	q := "SELECT id, currency_from, currency_to, cource, updated_at FROM exchange_rate;"
	if err = d.Conn.SelectContext(ctx, &result, q); err != nil {
		return nil, err
	}
	return result, err
}

func (d *DB) GetExchangeRate(ctx context.Context, from, to string) (result []ExchangeRate, err error) {
	q := "SELECT id, currency_from, currency_to, cource, updated_at FROM exchange_rate WHERE currency_from = $1 AND currency_to = $2;"
	if err = d.Conn.SelectContext(ctx, &result, q, from, to); err != nil {
		return nil, err
	}
	return result, err
}

func (d *DB) CreateExchangeRate(ctx context.Context, from, to string) error {
	q := "INSERT INTO exchange_rate (currency_from, currency_to) VALUES ($1, $2);"
	_, err := d.Conn.ExecContext(ctx, q, from, to)
	return err
}

func (d *DB) UpdateExchangeRate(ctx context.Context, er ExchangeRate) error {
	q := "UPDATE exchange_rate SET currency_from = $1, currency_to = $2, cource = $3, updated_at = $4 WHERE id = $5;"
	_, err := d.Conn.ExecContext(ctx, q, er.CurrencyFrom, er.CurrencyTo, er.Cource, time.Now(), er.ID)
	return err
}

func (d *DB) Close() {
	d.Conn.Close()
}
