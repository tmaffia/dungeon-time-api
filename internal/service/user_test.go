package service

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/tmaffia/dungeon-time-api/internal/repo"
)

func TestBuildUser(t *testing.T) {
	type args struct {
		username string
		email    string
		password string
	}
	tests := []struct {
		name     string
		args     args
		want     *userBuilder
		password string
		wantErr  error
	}{
		{"TestBuildUser Success",
			args{"testusername", "test@gmail.com", "test12345!"},
			&userBuilder{&User{Username: "testusername", Email: "test@gmail.com"}},
			"test12345!",
			nil,
		},
		{"TestBuildUser Invalid Password",
			args{"testusername", "test@gmail.com", "test"},
			nil,
			"test",
			ErrInvalidPassword,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := BuildUser(tt.args.username, tt.args.email, tt.args.password)
			if err != nil {
				assert.ErrorIs(t, err, tt.wantErr)
				return
			}

			assert.Equal(t, tt.want.user.Email, got.user.Email)
			assert.Equal(t, tt.want.user.Username, got.user.Username)
			assert.True(t, got.user.ValidatePassword(tt.password))
		})
	}
}

func TestUser_ValidatePassword(t *testing.T) {
	type args struct {
		password string
	}
	tests := []struct {
		name string
		u    *User
		args args
		want bool
	}{
		{
			"TestValidatePassword Success",
			&User{passwordHash: "$2a$10$Hur1mzq5JZbbXAYBvwgH0uAOlc5dOPn0EswvqVmY6PTBdquTBiXs."},
			args{"test12345!"},
			true,
		},
		{
			"TestValidatePassword Fail",
			&User{passwordHash: "$2a$10$Hur1mzq5JZbbXAYBvwgH0uAOlc5dOPn0EswvqVmY6PTBdquTBiXs."},
			args{"test"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.u.ValidatePassword(tt.args.password))
		})
	}
}

func Test_userService_RegisterUser(t *testing.T) {
	type args struct {
		ctx  context.Context
		user *User
	}
	tests := []struct {
		name    string
		s       *userService
		args    args
		want    *User
		wantErr error
	}{
		{
			"TestRegisterUser Success",
			&userService{},
			args{context.Background(), &User{
				Username:     "testusername",
				Email:        "example@example.com",
				passwordHash: "$2a$10$Hur1mzq5JZbbXAYBvwgH0uAOlc5dOPn0EswvqVmY6PTBdquTBiXs.",
			}},
			&User{
				ID:           0,
				Username:     "testusername",
				Email:        "example@example.com",
				passwordHash: "$2a$10$Hur1mzq5JZbbXAYBvwgH0uAOlc5dOPn0EswvqVmY6PTBdquTBiXs.",
			},
			nil,
		},
		// {
		// 	"TestRegisterUser Invalid User",
		// 	&userService{},
		// 	args{context.Background(), &User{
		// 		Username:     "testusername",
		// 		Email:        "example@example.com",
		// 		passwordHash: "$2a$10$Hur1mzq5JZbbXAYBvwgH0uAOlc5dOPn0EswvqVmY6PTBdquTBiXs.",
		// 	}},
		// 	&User{
		// 		ID:           0,
		// 		Username:     "testusername",
		// 		Email:        "example@example.com",
		// 		passwordHash: "$2a$10$Hur1mzq5JZbbXAYBvwgH0uAOlc5dOPn0EswvqVmY6PTBdquTBiXs.",
		// 	},
		// 	ErrInvalidUser,
		// },
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Mock the userRepo
			mockq := repo.NewMockQuerier(t)
			mockq.EXPECT().CreateUser(tt.args.ctx, repo.CreateUserParams{
				Username:     tt.args.user.Username,
				Email:        tt.args.user.Email,
				PasswordHash: tt.args.user.passwordHash,
				Roles:        []string{},
			}).Return(repo.User{
				ID:       0,
				Username: tt.args.user.Username,
				Email:    tt.args.user.Email,
				Roles:    []string{},
				Timezone: "UTC",
			}, nil)
			tt.s = &userService{userRepo: mockq}

			u, err := tt.s.RegisterUser(tt.args.ctx, tt.args.user)
			if !assert.ErrorIs(t, err, tt.wantErr) {
				return
			}
			assert.Equal(t, tt.want, u)
		})
	}
}

func Test_userService_GetUsers(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		s       *userService
		args    args
		want    []*User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.GetUsers(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("userService.GetUsers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userService.GetUsers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userService_GetUserByID(t *testing.T) {
	type args struct {
		ctx context.Context
		id  int32
	}
	tests := []struct {
		name    string
		s       *userService
		args    args
		want    *User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.GetUserByID(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("userService.GetUserByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userService.GetUserByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userService_GetUserByEmail(t *testing.T) {
	type args struct {
		ctx   context.Context
		email string
	}
	tests := []struct {
		name    string
		s       *userService
		args    args
		want    *User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.GetUserByEmail(tt.args.ctx, tt.args.email)
			if (err != nil) != tt.wantErr {
				t.Errorf("userService.GetUserByEmail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userService.GetUserByEmail() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userService_GetUserByUsername(t *testing.T) {
	type args struct {
		ctx      context.Context
		username string
	}
	tests := []struct {
		name    string
		s       *userService
		args    args
		want    *User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.GetUserByUsername(tt.args.ctx, tt.args.username)
			if (err != nil) != tt.wantErr {
				t.Errorf("userService.GetUserByUsername() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userService.GetUserByUsername() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mapRoles(t *testing.T) {
	type args struct {
		roleStrings []string
	}
	tests := []struct {
		name    string
		args    args
		want    []UserRole
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := mapRoles(tt.args.roleStrings)
			if (err != nil) != tt.wantErr {
				t.Errorf("mapRoles() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mapRoles() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mapTimezone(t *testing.T) {
	type args struct {
		timezone string
	}
	tests := []struct {
		name    string
		args    args
		want    time.Location
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := mapTimezone(tt.args.timezone)
			if (err != nil) != tt.wantErr {
				t.Errorf("mapTimezone() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mapTimezone() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isValidUser(t *testing.T) {
	type args struct {
		user *User
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := isValidUser(tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("isValidUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_isValidUsername(t *testing.T) {
	type args struct {
		username string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidUsername(tt.args.username); got != tt.want {
				t.Errorf("isValidUsername() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isValidEmail(t *testing.T) {
	type args struct {
		email string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidEmail(tt.args.email); got != tt.want {
				t.Errorf("isValidEmail() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isValidPassword(t *testing.T) {
	type args struct {
		password string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidPassword(tt.args.password); got != tt.want {
				t.Errorf("isValidPassword() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isValidRoles(t *testing.T) {
	type args struct {
		roles []UserRole
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidRoles(tt.args.roles...); got != tt.want {
				t.Errorf("isValidRoles() = %v, want %v", got, tt.want)
			}
		})
	}
}
