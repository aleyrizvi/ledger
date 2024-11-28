package engine

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"strings"

	"github.com/go-playground/validator"
)

type BaseHandler struct {
	validator *validator.Validate
	router    *http.ServeMux
}

func (h *BaseHandler) Init(routes []Route) {
	// init validator
	h.validator = validator.New()
	h.validator.RegisterTagNameFunc(func(f reflect.StructField) string {
		name := strings.SplitN(f.Tag.Get("json"), ",", 2)

		return strings.Join(name, "")
	})

	// init router
	r := http.NewServeMux()

	for _, route := range routes {
		pattern := fmt.Sprintf("%s %s", route.Method, route.Path)
		r.Handle(pattern, applyMiddleware(http.HandlerFunc(route.Handler), route.Middleware...))
	}

	h.router = r
}

func (h *BaseHandler) ParseAndValidate(w http.ResponseWriter, r *http.Request, data interface{}) error {
	if err := json.NewDecoder(r.Body).Decode(data); err != nil {
		h.Error(w, r, APIError{
			Code:     http.StatusBadRequest,
			Messages: nil,
		})

		return err
	}

	if err := h.validator.Struct(data); err != nil {
		fmt.Println("hit the validator error")
		h.Error(w, r, APIError{
			Code:     http.StatusBadRequest,
			Messages: h.FormatErrors(err),
		})

		return err
	}

	return nil
}

func (h *BaseHandler) JSON(w http.ResponseWriter, _ *http.Request, code int, data interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	if data == nil {
		return nil
	}

	return json.NewEncoder(w).Encode(data)
}

func (h *BaseHandler) Error(w http.ResponseWriter, _ *http.Request, err error) error {
	w.Header().Set("Content-Type", "application/json")

	var e APIError
	if !errors.As(err, &e) {
		w.WriteHeader(http.StatusInternalServerError)

		return nil
	}

	w.WriteHeader(e.Code)

	if len(e.Messages) == 0 {
		return nil
	}

	return json.NewEncoder(w).Encode(e)
}

func (h *BaseHandler) FormatErrors(err error) []string {
	var Verr validator.ValidationErrors

	errors.As(err, &Verr)

	var errorResponse []string

	for _, e := range Verr {
		var err string
		switch e.Tag() {
		case "required":
			err = fmt.Sprintf("[%s] is required", e.Field())
		case "email":
			err = fmt.Sprintf("[%s] must be a valid email", e.Field())
		case "gte":
			err = fmt.Sprintf("[%s] must be greater than or equal to %s", e.Field(), e.Param())
		case "oneof":
			err = fmt.Sprintf("[%s] must be one of [%s]", e.Field(), e.Param())
		default:
			err = fmt.Sprintf("[%s] invalid error", e.Field())
		}
		errorResponse = append(errorResponse, err)
	}

	return errorResponse
}

func (h *BaseHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.router.ServeHTTP(w, r)
}
