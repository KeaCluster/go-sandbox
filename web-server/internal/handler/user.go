package handler

import (
	"encoding/json"
	"net/http"
	"web-server/internal/model"
	"web-server/util"
)

func UserHandler(w http.ResponseWriter, r *http.Request) {
	// create user
	user := model.User{
		ID:    1,
		Name:  "Nacho Libre",
		Email: "nacho_l@gmail.com",
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(user); err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, "Failed to encode user")
	}
}
