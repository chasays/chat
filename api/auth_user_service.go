package main

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"github.com/swuecho/chatgpt_backend/auth"
	"github.com/swuecho/chatgpt_backend/sqlc_queries"
)

type AuthUserService struct {
	q *sqlc_queries.Queries
}

// NewAuthUserService creates a new AuthUserService.
func NewAuthUserService(q *sqlc_queries.Queries) *AuthUserService {
	return &AuthUserService{q: q}
}

// CreateAuthUser creates a new authentication user record.
func (s *AuthUserService) CreateAuthUser(ctx context.Context, auth_user_params sqlc_queries.CreateAuthUserParams) (sqlc_queries.AuthUser, error) {
	auth_user, err := s.q.CreateAuthUser(ctx, auth_user_params)
	if err != nil {
		return sqlc_queries.AuthUser{}, err
	}
	return auth_user, nil
}

// GetAuthUserByID returns an authentication user record by ID.
func (s *AuthUserService) GetAuthUserByID(ctx context.Context, id int32) (sqlc_queries.AuthUser, error) {
	auth_user, err := s.q.GetAuthUserByID(ctx, id)
	if err != nil {
		return sqlc_queries.AuthUser{}, errors.New("failed to retrieve authentication user")
	}
	return auth_user, nil
}

// UpdateAuthUser updates an existing authentication user record.
func (s *AuthUserService) UpdateAuthUser(ctx context.Context, auth_user_params sqlc_queries.UpdateAuthUserParams) (sqlc_queries.AuthUser, error) {
	auth_user_u, err := s.q.UpdateAuthUser(ctx, auth_user_params)
	if err != nil {
		return sqlc_queries.AuthUser{}, errors.New("failed to update authentication user")
	}
	return auth_user_u, nil
}

// GetAllAuthUsers returns all authentication user records.
func (s *AuthUserService) GetAllAuthUsers(ctx context.Context) ([]sqlc_queries.AuthUser, error) {
	auth_users, err := s.q.GetAllAuthUsers(ctx)
	if err != nil {
		return nil, errors.New("failed to retrieve authentication users")
	}
	return auth_users, nil
}

func (s *AuthUserService) Authenticate(ctx context.Context, email, password string) (sqlc_queries.AuthUser, error) {
	user, err := s.q.GetUserByEmail(ctx, email)
	println(email, password)
	if err != nil {
		return sqlc_queries.AuthUser{}, err
	}
	println("x|" + user.Password + "|xxx")
	if !auth.ValidatePassword(password, user.Password) {
		return sqlc_queries.AuthUser{}, ErrInvalidCredentials
	}
	return user, nil
}

func (s *AuthUserService) Logout(tokenString string) (*http.Cookie, error) {
	userID, err := auth.ValidateToken(tokenString, appConfig.JWT.SECRET)
	if err != nil {
		return nil, err
	}
	// Implement a mechanism to track invalidated tokens for the given user ID
	// auth.AddInvalidToken(userID, "insert-invalidated-token-here")
	cookie := auth.GetExpireSecureCookie(strconv.Itoa(int(userID)), false)
	return cookie, nil
}
