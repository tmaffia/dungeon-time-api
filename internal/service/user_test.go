package service

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/assert"
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

func Test_userBuilder_Roles(t *testing.T) {
	type args struct {
		roles []UserRole
	}
	tests := []struct {
		name string
		ub   *userBuilder
		args args
		want *userBuilder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ub.Roles(tt.args.roles...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userBuilder.Roles() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userBuilder_Timezone(t *testing.T) {
	type args struct {
		timezone time.Location
	}
	tests := []struct {
		name string
		ub   *userBuilder
		args args
		want *userBuilder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ub.Timezone(tt.args.timezone); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userBuilder.Timezone() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userBuilder_Build(t *testing.T) {
	tests := []struct {
		name string
		ub   *userBuilder
		want *User
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ub.Build(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userBuilder.Build() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUser_GetPasswordHash(t *testing.T) {
	tests := []struct {
		name string
		u    *User
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.u.GetPasswordHash(); got != tt.want {
				t.Errorf("User.GetPasswordHash() = %v, want %v", got, tt.want)
			}
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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.u.ValidatePassword(tt.args.password); got != tt.want {
				t.Errorf("User.ValidatePassword() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewUserService(t *testing.T) {
	type args struct {
		dbPool *pgxpool.Pool
	}
	tests := []struct {
		name string
		args args
		want *userService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUserService(tt.args.dbPool); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUserService() = %v, want %v", got, tt.want)
			}
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
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.RegisterUser(tt.args.ctx, tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("userService.RegisterUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userService.RegisterUser() = %v, want %v", got, tt.want)
			}
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
