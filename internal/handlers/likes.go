package handlers

import (
	"forum/internal/middleware"
	"forum/internal/services"
)

type LikeHandler struct {
	LikeService   *services.LikeService
	AuthService   *services.AuthService
	AuthMidlaware *middleware.AuthMiddleware
}
