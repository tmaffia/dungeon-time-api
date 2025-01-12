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

type userBuilder struct {
	user *User
}

func BuildUser(username, email, password string) (*userBuilder, error) {
	if !isValidPassword(password) {
		return nil, ErrorInvalidPassword
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

func (ub *userBuilder) Roles(roles ...UserRole) *userBuilder {
	ub.user.Roles = roles
	return ub
}

func (ub *userBuilder) Timezone(timezone time.Location) *userBuilder {
	ub.user.Timezone = timezone
	return ub
}

func (ub *userBuilder) Build() *User {
	return ub.user
}

func (u *User) GetPasswordHash() string {
	return u.passwordHash
}

func (u *User) ValidatePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.passwordHash), []byte(password))
	return err == nil
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

type UserService interface {
	RegisterUser(context.Context, *User) (*User, error)
	GetUsers(context.Context) ([]*User, error)
	GetUserByID(context.Context, int32) (*User, error)
	GetUserByEmail(context.Context, string) (*User, error)
	GetUserByUsername(context.Context, string) (*User, error)
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
		return nil, ErrorInvalidRole
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
		return ErrorInvalidUsername
	}

	if !isValidEmail(user.Email) {
		return ErrorInvalidEmail
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
