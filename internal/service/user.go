package service

import (
	"context"
	"errors"
	"slices"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/tmaffia/dungeon-time-api/internal/repo"
)

type User struct {
	ID           int32      `json:"id"`
	Username     string     `json:"username"`
	Email        string     `json:"email"`
	PasswordHash string     `json:"password"`
	Roles        []UserRole `json:"roles"`
	Timezone     string     `json:"timezone"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
}

type UserRole string

const (
	RoleLeader = UserRole("Leader")
	RoleMember = UserRole("Member")
	RoleTank   = UserRole("Tank")
	RoleHealer = UserRole("Healer")
	RoleDPS    = UserRole("DPS")
)

var userRoles = []UserRole{RoleLeader, RoleMember,
	RoleTank, RoleHealer, RoleDPS}

func ValidateUserRole(s string) (UserRole, error) {
	r := UserRole(s)
	if slices.Contains(userRoles, r) {
		return r, nil
	}
	return "", errors.New("invalid user role: " + s)
}

type UserService interface {
	RegisterUser(context.Context, string, string, ...func(*User)) (*User, error)
	GetUsers(context.Context) ([]*User, error)
	GetUserByID(context.Context, int32) (*User, error)
}

type userService struct {
	dbPool   *pgxpool.Pool
	userRepo *repo.Queries
}

func NewUserService(dbPool *pgxpool.Pool) userService {
	return userService{
		dbPool:   dbPool,
		userRepo: repo.New(dbPool),
	}
}

func (s userService) GetUserByID(ctx context.Context, id int32) (*User, error) {
	s.userRepo.GetUserByID(ctx, id)
	return nil, nil
}

func (s userService) GetUsers(ctx context.Context) ([]*User, error) {
	var users []*User
	u, err := s.userRepo.GetUsers(ctx)
	if err != nil {
		return nil, err
	}

	for _, user := range u {
		roles, err := mapRoles(user.Roles)
		if err != nil {
			return nil, err
		}

		users = append(users, &User{
			ID:           user.ID,
			Username:     user.Username,
			Email:        user.Email,
			PasswordHash: user.PasswordHash,
			Roles:        roles,
			Timezone:     user.Timezone,
			CreatedAt:    user.CreatedAt.Time,
			UpdatedAt:    user.UpdatedAt.Time,
		})
	}
	return users, nil

}

func (s userService) RegisterUser(ctx context.Context, username, email string, opts ...func(*User)) (*User, error) {
	return nil, nil
}

func mapRoles(roleStrings []string) ([]UserRole, error) {
	var roles []UserRole
	for _, role := range roleStrings {
		r, err := ValidateUserRole(role)
		if err != nil {
			return nil, err
		}

		roles = append(roles, r)
	}
	return roles, nil
}

// // Helper functions for validation (could be in a separate utility package)
// func isValidUsername(username string) bool {
// 	// Implement your username validation logic (e.g., length, allowed characters)
// 	return len(username) >= 3 && len(username) <= 20
// }

// var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

// func isValidEmail(email string) bool {
// 	return emailRegex.MatchString(email)
// }
