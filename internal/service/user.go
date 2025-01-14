package service

import (
	"context"
	"regexp"
	"slices"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/tmaffia/dungeon-time-api/internal/repo"
	"golang.org/x/crypto/bcrypt"
)

// User represents a user account for the application.
// The passwordHash field is not exported to prevent accidental exposure.
// Use GetPasswordHash() to retrieve the password hash.
// Use ValidatePassword() to check if a password matches the hash.
// Use userBuilder to create a new user.
type User struct {
	ID           int32  `json:"id"`
	Username     string `json:"username"`
	Email        string `json:"email"`
	passwordHash string
	Roles        []UserRole    `json:"roles"`
	Timezone     time.Location `json:"timezone"`
	CreatedAt    time.Time     `json:"created_at"`
	UpdatedAt    time.Time     `json:"updated_at"`
}

// Builder object for User struct
type userBuilder struct {
	user *User
}

// BuildUser creates a userBuilder to assemble a valid User.
// Returns an error if the password is invalid according to the
// password rules. Takes all the required fields for a User.
// The Builder has methods to add optional fields, such as Roles and Timezone.
// Call Build() to create the User.
func BuildUser(username, email, password string) (*userBuilder, error) {
	if !isValidPassword(password) {
		return nil, ErrInvalidPassword
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return &userBuilder{
		user: &User{
			Username:     username,
			Email:        email,
			passwordHash: string(hashedPassword),
		},
	}, nil
}

// Roles adds roles to the userBuilder
func (ub *userBuilder) Roles(roles ...UserRole) *userBuilder {
	ub.user.Roles = roles
	return ub
}

// Timezone adds a timezone to the userBuilder
func (ub *userBuilder) Timezone(timezone time.Location) *userBuilder {
	ub.user.Timezone = timezone
	return ub
}

// Build creates a User from the userBuilder
// Call at the end of a typical builder pattern chain
func (ub *userBuilder) Build() *User {
	return ub.user
}

// GetPasswordHash returns the password hash for the user
func (u *User) GetPasswordHash() string {
	return u.passwordHash
}

// ValidatePassword checks if the provided password matches the user's password hash
// Returns true if the password matches, false otherwise
// Use this method to validate a user's password
func (u *User) ValidatePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.passwordHash), []byte(password))
	return err == nil
}

// UserRole represents the role of a user in the application
// This type is used in place of enum as Go does not have an enum type... somehow...
type UserRole string

const (
	RoleLeader = UserRole("Leader")
	RoleMember = UserRole("Member")
	RoleTank   = UserRole("Tank")
	RoleHealer = UserRole("Healer")
	RoleDPS    = UserRole("DPS")
)

// Pre defined user roles, these are the only valid roles for a user.
var userRoles = []UserRole{RoleLeader, RoleMember,
	RoleTank, RoleHealer, RoleDPS}

// UserService is the interface for user-related operations.
type UserService interface {
	RegisterUser(context.Context, *User) (*User, error)
	GetUsers(context.Context) ([]*User, error)
	GetUserByID(context.Context, int32) (*User, error)
	GetUserByEmail(context.Context, string) (*User, error)
	GetUserByUsername(context.Context, string) (*User, error)
}

// userService is the implementation of UserService. It uses a database connection
// pool and a repository to interact with the database. It is responsible for
// user-related operations.
type userService struct {
	dbPool   *pgxpool.Pool
	userRepo repo.Querier
}

// NewUserService creates a new userService with the provided database connection pool.
// It returns a pointer to the userService.
func NewUserService(dbPool *pgxpool.Pool) *userService {
	return &userService{
		dbPool:   dbPool,
		userRepo: repo.New(dbPool),
	}
}

// RegisterUser creates a new user in the database. It takes an already created User
// and performs validation on that user. Users should always be created using the BuildUser() function
// to ensure that the password is hashed correctly. Returns the created user if successful.
// Returns an error specific to the validation problem if the user is invalid.
func (s *userService) RegisterUser(ctx context.Context, user *User) (*User, error) {
	err := isValidUser(user)
	if err != nil {
		return nil, err
	}

	u, err := s.userRepo.CreateUser(ctx, repo.CreateUserParams{
		Username:     user.Username,
		Email:        user.Email,
		PasswordHash: user.passwordHash,
		Roles:        []string{},
		Timezone:     user.Timezone.String(),
	})

	if err != nil {
		return nil, err
	}

	user.ID = u.ID
	user.CreatedAt = u.CreatedAt.Time
	user.UpdatedAt = u.UpdatedAt.Time

	return user, nil
}

// GetUsers returns all registered users. Only the ID, Username, Email, Roles, and Timezone
// fields are returned for each user.
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

		tz, err := mapTimezone(user.Timezone)
		if err != nil {
			return nil, err
		}

		users = append(users, &User{
			ID:       user.ID,
			Username: user.Username,
			Email:    user.Email,
			Roles:    roles,
			Timezone: tz,
		})
	}
	return users, nil
}

// GetUserByID returns a user by ID. Only the ID, Username, Email, Roles, and Timezone
// fields are returned for the user.
func (s *userService) GetUserByID(ctx context.Context, id int32) (*User, error) {
	u, err := s.userRepo.GetUserByID(ctx, id)
	if err != nil {
		return nil, err
	}

	roles, err := mapRoles(u.Roles)
	if err != nil {
		return nil, err
	}

	tz, err := mapTimezone(u.Timezone)
	if err != nil {
		return nil, err
	}

	return &User{
		ID:       u.ID,
		Username: u.Username,
		Email:    u.Email,
		Roles:    roles,
		Timezone: tz,
	}, nil
}

// GetUserByEmail returns a user by email. Only the ID, Username, Email, Roles, and Timezone
// fields are returned for the user.
func (s *userService) GetUserByEmail(ctx context.Context, email string) (*User, error) {
	u, err := s.userRepo.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	roles, err := mapRoles(u.Roles)
	if err != nil {
		return nil, err
	}

	tz, err := mapTimezone(u.Timezone)
	if err != nil {
		return nil, err
	}

	return &User{
		ID:       u.ID,
		Username: u.Username,
		Email:    u.Email,
		Roles:    roles,
		Timezone: tz,
	}, nil
}

// GetUserByUsername returns a user by username. Only the ID, Username, Email, Roles, and Timezone
// fields are returned for the user.
func (s *userService) GetUserByUsername(ctx context.Context, username string) (*User, error) {
	u, err := s.userRepo.GetUserByUsername(ctx, username)
	if err != nil {
		return nil, err
	}

	roles, err := mapRoles(u.Roles)
	if err != nil {
		return nil, err
	}

	tz, err := mapTimezone(u.Timezone)
	if err != nil {
		return nil, err
	}

	return &User{
		ID:       u.ID,
		Username: u.Username,
		Email:    u.Email,
		Roles:    roles,
		Timezone: tz,
	}, nil
}

func mapRoles(roleStrings []string) ([]UserRole, error) {
	var roles []UserRole
	for _, role := range roleStrings {
		r := UserRole(role)
		roles = append(roles, r)
	}
	if !isValidRoles(roles...) {
		return nil, ErrInvalidRole
	}
	return roles, nil
}

func mapTimezone(timezone string) (time.Location, error) {
	loc, err := time.LoadLocation(timezone)
	if err != nil {
		return time.Location{}, err
	}
	return *loc, nil
}

func isValidUser(user *User) error {
	if !isValidUsername(user.Username) {
		return ErrInvalidUsername
	}

	if !isValidEmail(user.Email) {
		return ErrInvalidEmail
	}

	return nil
}

func isValidUsername(username string) bool {
	return len(username) >= 3 && len(username) <= 20 &&
		regexp.MustCompile(`^[a-zA-Z]+$`).MatchString(username)
}

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

func isValidEmail(email string) bool {
	return emailRegex.MatchString(email)
}

func isValidPassword(password string) bool {
	return len(password) >= 8
}

func isValidRoles(roles ...UserRole) bool {
	for _, r := range roles {
		if !slices.Contains(userRoles, r) {
			return false
		}
	}
	return true
}
