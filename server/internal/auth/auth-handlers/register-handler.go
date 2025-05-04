package auth_handlers

import (
	"encoding/json"
	"net/http"
	"server/internal/auth"
	"server/internal/auth/helper"
)

type registerRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func RegisterHandler(authSvc *auth.AuthService) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		var req registerRequest
		if err := json.NewDecoder(request.Body).Decode(&req); err != nil {
			http.Error(writer, "Invalid JSON body", http.StatusBadRequest)
			return
		}

		if req.Email == "" || req.Password == "" {
			http.Error(writer, "Email and password are required", http.StatusBadRequest)
			return
		}

		if !helper.IsValidEmail(req.Email) {
			http.Error(writer, "Invalid email", http.StatusBadRequest)
			return
		}

		user, err := authSvc.RegisterUser(request.Context(), req.Email, req.Password)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
			return
		}

		writer.WriteHeader(http.StatusCreated)
		json.NewEncoder(writer).Encode(map[string]interface{}{
			"id":    user.ID,
			"email": user.Email,
		})
	}
}
