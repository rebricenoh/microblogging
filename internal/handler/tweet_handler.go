package handler

import (
	"encoding/json"
	"microblogging/internal/service"
	"net/http"
	"strconv"
)

// TweetHandler maneja las solicitudes HTTP para los tweets.
type TweetHandler struct {
	service *service.TweetService
}

// NewTweetHandler crea una nueva instancia de TweetHandler.
func NewTweetHandler(service *service.TweetService) *TweetHandler {
	return &TweetHandler{service: service}
}

// PostTweetHandler maneja la solicitud de publicación de un tweet.
func (h *TweetHandler) PostTweetHandler(w http.ResponseWriter, r *http.Request) {
	userID, _ := strconv.Atoi(r.Header.Get("User-ID"))
	var req struct {
		Content string `json:"content"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	if err := h.service.PostTweet(userID, req.Content); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
