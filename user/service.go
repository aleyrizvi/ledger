package user

import (
	"context"
	"database/sql"
	"math"
)

type state string

var (
	StateWin  state = "win"
	StateLose state = "lose"
)

type UserService interface {
	GetUser(ctx context.Context, id uint32) (*User, error)
	UpdateTransaction(ctx context.Context, TransactionOpts *UpdateTransactionOptions) error
}

type UpdateTransactionOptions struct {
	TxID   string
	UserID uint32
	State  state
	Amount float64
}

type Service struct {
	repo Repository
}

func (s *Service) UpdateTransaction(ctx context.Context, opts *UpdateTransactionOptions) error {
	_, err := s.GetUser(ctx, opts.UserID)
	if err != nil {
		// TODO: User not found probably. Should return appropriate error
		return err
	}

	// Check if transaction already exists
	_, err = s.repo.GetTransaction(ctx, opts.TxID)
	if err != nil && err != sql.ErrNoRows {
		return err
	}

	// transaction already exists
	if err == nil {
		return nil
	}

	amount := int32(math.Round(opts.Amount * 100))
	if opts.State == StateLose {
		amount = -amount
	}
	return s.repo.UpdateTransaction(ctx, opts.TxID, opts.UserID, amount)
}

func (s *Service) GetUser(ctx context.Context, id uint32) (*User, error) {
	return s.repo.GetUser(ctx, id)
}

func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}
