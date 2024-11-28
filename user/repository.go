package user

import (
	"context"
	"database/sql"
)

type Repository interface {
	GetUser(ctx context.Context, id uint32) (*User, error)
	GetTransaction(ctx context.Context, txID string) (*Transaction, error)
	UpdateTransaction(ctx context.Context, txID string, id uint32, amount int32) error
}

type Repo struct {
	*sql.DB
	dbrw *sql.DB
}

func (r *Repo) GetUser(_ context.Context, id uint32) (*User, error) {
	var u User

	err := r.QueryRow("SELECT id, balance from users where id=$1", id).Scan(&u.ID, &u.Balance)

	return &u, err
}

func (r *Repo) GetTransaction(ctx context.Context, txID string) (*Transaction, error) {
	var t Transaction
	err := r.QueryRowContext(ctx, `SELECT  
		transaction_id, 
		user_id, 
		amount, 
		created_at
		FROM transactions WHERE transaction_id = $1`, txID).
		Scan(&t.ID, &t.UserID, &t.Amount, &t.CreatedAt)

	return &t, err
}

func (r *Repo) UpdateTransaction(ctx context.Context, txID string, id uint32, amount int32) error {
	tx, err := r.dbrw.Begin()
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, `
		INSERT INTO transactions (transaction_id, user_id, amount)
		VALUES ($1, $2, $3)
		`, txID, id, amount)
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.ExecContext(ctx, `
		UPDATE users SET balance = balance + $1 WHERE id = $2
		`, amount, id)
	if err != nil {
		tx.Rollback()
		return nil
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func NewRepository(dbrw, dbro *sql.DB) *Repo {
	return &Repo{
		dbrw: dbrw,
		DB:   dbro,
	}
}
