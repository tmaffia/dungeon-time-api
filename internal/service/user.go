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

func NewUserService(dbPool *pgxpool.Pool) *userService {
	return &userService{
		dbPool:   dbPool,
		userRepo: repo.New(dbPool),
	}
}

func (s *userService) GetUserByID(ctx context.Context, id int32) (*User, error) {
	s.userRepo.GetUserByID(ctx, id)
	return nil, nil
}

func (s *userService) GetUsers(ctx context.Context) ([]*User, error) {
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

func (s *userService) RegisterUser(ctx context.Context, username, email string, opts ...func(*User)) (*User, error) {
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

// func WithPassword(password string) func(*repo.User) {
// 	return func(u *repo.User) {
// 		u.Password = password
// 	}
// }

// // NewUserService creates a new instance of userService.
// func NewUserService(userRepo repo.Queries) *userService {
// 	return &userService{userRepo: &userRepo}
// }

// // RegisterUser registers a new user.
// func (s *userService) RegisterUser(ctx context.Context, username, email string) (*repository.User, error) {
// 	// 1. Validate input
// 	if !isValidUsername(username) {
// 		return nil, errors.New("invalid username format")
// 	}
// 	if !isValidEmail(email) {
// 		return nil, errors.New("invalid email format")
// 	}

// 	// 2. Check if the user already exists (optional, depending on your requirements)
// 	existingUser, err := s.userRepo.GetUserByEmail(ctx, email)     // Assume you have GetUserByEmail in your repo
// 	if err != nil && !errors.Is(err, repository.ErrUserNotFound) { // Handle database errors gracefully
// 		return nil, err
// 	}
// 	if existingUser != nil {
// 		return nil, errors.New("user with this email already exists")
// 	}

// 	// 3. Create the user (consider hashing passwords here if you have them)
// 	newUser := &repository.User{
// 		Username: username,
// 		Email:    email,
// 	}
// 	err = s.userRepo.CreateUser(ctx, newUser)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// 4. Return the created user (you might fetch it again to get the generated ID)
// 	return s.userRepo.GetUserByID(ctx, newUser.ID)
// }

// // GetUserProfile retrieves a user's profile.
// func (s *userService) GetUserProfile(ctx context.Context, userID int64) (*repository.User, error) {
// 	user, err := s.userRepo.GetUserByID(ctx, userID)
// 	if err != nil {
// 		return nil, err // Could wrap the error with more context: fmt.Errorf("get user profile: %w", err)
// 	}
// 	if user == nil {
// 		return nil, errors.New("user not found") // You might use a custom error here
// 	}
// 	return user, nil
// }

// // UpdateUserEmail updates a user's email address.
// func (s *userService) UpdateUserEmail(ctx context.Context, userID int64, newEmail string) error {
// 	// 1. Validate the new email
// 	if !isValidEmail(newEmail) {
// 		return errors.New("invalid email format")
// 	}
// 	// 2. Check if the user exist
// 	user, err := s.userRepo.GetUserByID(ctx, userID)
// 	if err != nil {
// 		return err
// 	}
// 	if user == nil {
// 		return errors.New("user not found")
// 	}
// 	// 3. Check if email is already in used.
// 	existingUser, err := s.userRepo.GetUserByEmail(ctx, newEmail) // Assume you have GetUserByEmail in your repo
// 	if err != nil && !errors.Is(err, repository.ErrUserNotFound) {
// 		return err
// 	}
// 	if existingUser != nil {
// 		return errors.New("email is already in used")
// 	}

// 	// 4. Update email
// 	return s.userRepo.UpdateUserEmail(ctx, userID, newEmail) // Assume you have this method in your repo
// }

// // Helper functions for validation (could be in a separate utility package)
// func isValidUsername(username string) bool {
// 	// Implement your username validation logic (e.g., length, allowed characters)
// 	return len(username) >= 3 && len(username) <= 20
// }

// var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

// func isValidEmail(email string) bool {
// 	return emailRegex.MatchString(email)
// }
