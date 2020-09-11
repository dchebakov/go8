package http

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/render"

	"eight/internal/grpc/service"
)

func (h *Handlers) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var userRequest service.CreateUserRequest

		err := json.NewDecoder(r.Body).Decode(&userRequest.User)
		if err != nil {
			render.Status(r, http.StatusInternalServerError)
			return
		}

		createUserResponse, err := h.Grpc.UserServiceClient.CreateUser(context.Background(), &userRequest)
		if err != nil {
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, map[string]string{
				"error": err.Error(),
			})
			return
		}

		render.Status(r, http.StatusCreated)
		render.JSON(w, r, map[string]string{
			"token_string": createUserResponse.Token,
		})
		return
	}
}

func (h *Handlers) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var loginRequest service.LoginRequest

		err := json.NewDecoder(r.Body).Decode(&loginRequest)
		if err != nil {
			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, map[string]string{
				"error": err.Error(),
			})
			return
		}

		loginResponse, err := h.Grpc.UserServiceClient.Login(context.Background(), &loginRequest)
		if err != nil {
			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, map[string]string{
				"error": err.Error(),
			})
			return
		}

		render.Status(r, http.StatusOK)
		render.JSON(w, r, map[string]string{
			"token_string": loginResponse.Token,
		})
		return
	}
}

func (h *Handlers) Me() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		info, err := h.Authenticator.Me(r)
		if err != nil {
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, map[string]string{
				"error": "bad request",
			})
			return
		}

		render.Status(r, http.StatusOK)
		render.JSON(w, r, map[string]string{
			"username": info.Username(),
		})
		return
	}
}
