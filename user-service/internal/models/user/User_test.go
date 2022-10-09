package user

import (
	"testing"

	"gorm.io/gorm"
)

func TestUser_Prepare(t *testing.T) {
	tests := []struct {
		name     string
		user     *User
		wantUser *User
	}{
		{
			name:     "Trim text",
			user:     &User{Email: "  test@example.com "},
			wantUser: &User{Email: "test@example.com"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &User{
				ID:          tt.user.ID,
				Email:       tt.user.Email,
				Password:    tt.user.Password,
				IsVerified:  tt.user.IsVerified,
				CalendarID:  tt.user.CalendarID,
				IV_Calendar: tt.user.IV_Calendar,
				CreatedAt:   tt.user.CreatedAt,
				UpdatedAt:   tt.user.UpdatedAt,
			}
			u.Prepare()
			if u.Email != tt.wantUser.Email {
				t.Errorf("User.Prepare() want T '%v', but get T '%v'", tt.wantUser.Email, u.Email)
			}
			if u.CreatedAt.IsZero() || u.UpdatedAt.IsZero() {
				t.Errorf("User.Prepare() want CreatedAt and UpdatedAt not zero")
			}
			if u.CalendarID == "" {
				t.Errorf("User.Prepare() want CalendarID to be not empty")
			}
		})
	}
}

func TestUser_Validate(t *testing.T) {
	tests := []struct {
		name    string
		user    *User
		wantErr bool
	}{
		{
			name:    "fail if email is empty",
			user:    &User{Email: "", Password: "foopassword"},
			wantErr: true,
		},
		{
			name:    "fail if is not email",
			user:    &User{Email: "bonk12345", Password: "foopassword"},
			wantErr: true,
		},
		{
			name:    "fail if password is empty",
			user:    &User{Email: "test@example.com", Password: ""},
			wantErr: true,
		},
		{
			name:    "fail if password is too short",
			user:    &User{Email: "test@example.com", Password: "bonk"},
			wantErr: true,
		},
		{
			name: "fail if password is too long",
			user: &User{
				Email:    "test@example.com",
				Password: "168601827c50b37054f4e565dbf4050a6bd854bc91650280539cca45bae1fb2f1",
			},
			wantErr: true,
		},
		{
			name:    "should pass",
			user:    &User{Email: "test@example.com", Password: "foopassword"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &User{
				ID:          tt.user.ID,
				Email:       tt.user.Email,
				Password:    tt.user.Password,
				IsVerified:  tt.user.IsVerified,
				CalendarID:  tt.user.CalendarID,
				IV_Calendar: tt.user.IV_Calendar,
				CreatedAt:   tt.user.CreatedAt,
				UpdatedAt:   tt.user.UpdatedAt,
			}
			if err := u.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("User.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestHash(t *testing.T) {
	type args struct {
		password string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "should pass",
			args:    args{password: "foopassword"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Hash(tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("Hash() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got) == 0 && tt.wantErr {
				t.Errorf("Hash is empty")
			}
		})
	}
}

func TestVerifyPassword(t *testing.T) {
	type args struct {
		hashedPassword string
		password       string
	}

	testPass := "foopassword"
	hashedPass, _ := Hash(testPass)

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "should pass",
			args:    args{hashedPassword: string(hashedPass), password: testPass},
			wantErr: false,
		},
		{
			name:    "should fail",
			args:    args{hashedPassword: string(hashedPass), password: "wrongpassword"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := VerifyPassword(tt.args.hashedPassword, tt.args.password); (err != nil) != tt.wantErr {
				t.Errorf("VerifyPassword() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUser_BeforeCreate(t *testing.T) {
	type args struct {
		tx *gorm.DB
	}

	a := args{
		tx: &gorm.DB{},
	}

	tests := []struct {
		name    string
		user    *User
		args    args
		wantErr bool
	}{
		{
			name:    "should pass",
			user:    &User{Email: "test@example.com", Password: "password"},
			args:    a,
			wantErr: false,
		},
		{
			name:    "should fail on empty password",
			user:    &User{Email: "test@example.com", Password: ""},
			args:    a,
			wantErr: true,
		},
		{
			name:    "should fail on empty email",
			user:    &User{Email: "", Password: "password"},
			args:    a,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &User{
				ID:          tt.user.ID,
				Email:       tt.user.Email,
				Password:    tt.user.Password,
				IsVerified:  tt.user.IsVerified,
				CalendarID:  tt.user.CalendarID,
				IV_Calendar: tt.user.IV_Calendar,
				CreatedAt:   tt.user.CreatedAt,
				UpdatedAt:   tt.user.UpdatedAt,
			}
			if err := u.BeforeCreate(tt.args.tx); (err != nil) != tt.wantErr {
				t.Errorf("User.BeforeCreate() error = %v, wantErr %v", err, tt.wantErr)
			}

			if u.ID.String() == "" {
				t.Error("User.ID is empty")
			}
			if u.CreatedAt.IsZero() {
				t.Error("User.CreatedAt is zero")
			}
			if u.UpdatedAt.IsZero() {
				t.Error("User.UpdatedAt is zero")
			}
		})
	}
}

func TestUser_BeforeSave(t *testing.T) {
	type args struct {
		tx *gorm.DB
	}

	a := args{
		tx: &gorm.DB{},
	}

	tests := []struct {
		name    string
		user    *User
		args    args
		wantErr bool
	}{
		{
			name:    "should pass",
			user:    &User{Email: "test@example.com", Password: "password"},
			args:    a,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &User{
				ID:          tt.user.ID,
				Email:       tt.user.Email,
				Password:    tt.user.Password,
				IsVerified:  tt.user.IsVerified,
				CalendarID:  tt.user.CalendarID,
				IV_Calendar: tt.user.IV_Calendar,
				CreatedAt:   tt.user.CreatedAt,
				UpdatedAt:   tt.user.UpdatedAt,
			}
			_ = u.BeforeCreate(tt.args.tx)
			if err := u.BeforeSave(tt.args.tx); (err != nil) != tt.wantErr {
				t.Errorf("User.BeforeSave() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.user.Password == u.Password || u.Password == "" {
				t.Error("User.Password is not hashed")
			}
		})
	}
}
