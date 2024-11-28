package user

import (
	"net/http"
	"strconv"

	"github.com/aleyrizvi/ledger/engine"
)

type Handler struct {
	service UserService
	engine.BaseHandler
}

func (u *Handler) handleGetBalance(w http.ResponseWriter, r *http.Request) {
	type UserBalanceResponse struct {
		UserID  uint32  `json:"userId"`
		Balance float64 `json:"balance"`
	}

	ID, err := fetchUserID(r)
	if err != nil {
		u.Error(w, r, engine.ErrBadRequest([]string{"userID - must be a valid numeric id"}))
		return
	}

	user, err := u.service.GetUser(r.Context(), ID)
	if err != nil {
		// TODO: Can check error type here.
		u.Error(w, r, engine.ErrNotFound())
		return
	}

	u.JSON(w, r, http.StatusOK, &UserBalanceResponse{
		UserID:  user.ID,
		Balance: user.Balance.FromCents(),
	})
}

func (u *Handler) handlePostTransaction(w http.ResponseWriter, r *http.Request) {
	type TransactionRequest struct {
		State  state   `json:"state" validate:"required,oneof=win lose"`
		Amount float64 `json:"amount,string" validate:"required"`
		ID     string  `json:"transactionId" validate:"required"`
	}

	var req TransactionRequest

	err := u.ParseAndValidate(w, r, &req)
	if err != nil {
		return
	}

	ID, err := fetchUserID(r)
	if err != nil {
		u.Error(w, r, engine.ErrBadRequest([]string{"userID - must be a valid numeric id"}))
		return
	}

	err = u.service.UpdateTransaction(r.Context(), &UpdateTransactionOptions{
		TxID:   req.ID,
		UserID: ID,
		State:  req.State,
		Amount: req.Amount,
	})
	if err != nil {
		u.Error(w, r, engine.ErrBadRequest([]string{err.Error()}))
		return
	}

	u.JSON(w, r, http.StatusOK, nil)
}

func NewHandler(s UserService) *Handler {
	u := &Handler{
		service: s,
	}

	u.Init([]engine.Route{
		{
			Method:  http.MethodGet,
			Handler: u.handleGetBalance,
			Path:    "/{userID}/balance",
		},
		{
			Method:  http.MethodPost,
			Handler: u.handlePostTransaction,
			Path:    "/{userID}/transaction",
		},
	})

	return u
}

func fetchUserID(r *http.Request) (uint32, error) {
	u64ID, err := strconv.ParseUint(r.PathValue("userID"), 10, 32)
	if err != nil {
		return 0, engine.ErrBadRequest([]string{"userID - must be a valid numeric id"})
	}
	return uint32(u64ID), nil
}
