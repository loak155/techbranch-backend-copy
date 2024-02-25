package usecase

import (
	"context"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/loak155/techbranch-backend/internal/user/domain"
	"github.com/loak155/techbranch-backend/internal/user/repository"
	"github.com/loak155/techbranch-backend/mock"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func TestNewUserUsecase(t *testing.T) {
	type args struct {
		ur repository.IUserRepository
	}
	tests := []struct {
		name string
		args args
		want IUserUsecase
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUserUsecase(tt.args.ur); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUserUsecase() = %v, want %v", got, tt.want)
			}
		})
	}
}

// func testUser() *domain.User {
// 	return &domain.User{
// 		Username: "test_user",
// 		Email:    "test@example.com",
// 		Password: "password",
// 	}
// }

func TestCreateUser(t *testing.T) {
	req := domain.User{
		Username: "test_user",
		Email:    "test@example.com",
		Password: "password",
	}

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	ur := mock.NewMockIUserRepository(mockCtrl)
	ur.EXPECT().CreateUser(gomock.Any()).Return(nil)

	user, err := NewUserUsecase(ur).CreateUser(context.Background(), req)
	assert.NoError(t, err)
	assert.Equal(t, req.Username, user.Username)
	assert.Equal(t, req.Email, user.Email)
	assert.NotEqual(t, req.Password, user.Password)
	assert.NoError(t, bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)))
	assert.NotNil(t, user.CreatedAt)
	assert.NotNil(t, user.UpdatedAt)
	assert.Equal(t, user.DeletedAt, gorm.DeletedAt{})
}

func TestCreateUser2(t *testing.T) {
	type args struct {
		ctx  context.Context
		user domain.User
	}

	testUser := domain.User{
		Username: "test_user",
		Email:    "test@example.com",
		Password: "password",
	}

	testCases := []struct {
		name          string
		args          args
		buildStubs    func(ur *mock.MockIUserRepository)
		checkResponse func(t *testing.T, user domain.User, err error)
	}{
		{
			name: "OK",
			args: args{
				ctx: context.Background(),
				user: domain.User{
					Username: testUser.Username,
					Email:    testUser.Email,
					Password: testUser.Password,
				},
			},
			buildStubs: func(ur *mock.MockIUserRepository) {
				arg := domain.User{
					Username: testUser.Username,
					Email:    testUser.Email,
					Password: testUser.Password,
				}
				ur.EXPECT().CreateUser(arg).Return(nil)
			},
			checkResponse: func(t *testing.T, user domain.User, err error) {
				assert.NoError(t, err)
				assert.Equal(t, testUser.Username, user.Username)
				assert.Equal(t, testUser.Email, user.Email)
				assert.NotEqual(t, testUser.Password, user.Password)
				assert.NoError(t, bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(testUser.Password)))
				assert.NotNil(t, user.CreatedAt)
				assert.NotNil(t, user.UpdatedAt)
				assert.Equal(t, user.DeletedAt, gorm.DeletedAt{})
			},
		},
		{
			name: "OK2",
			args: args{
				ctx: context.Background(),
				user: domain.User{
					Username: testUser.Username,
					Email:    testUser.Email,
					Password: testUser.Password,
				},
			},
			buildStubs: func(ur *mock.MockIUserRepository) {
				ur.EXPECT().CreateUser(gomock.Any()).Return(nil)
			},
			checkResponse: func(t *testing.T, user domain.User, err error) {
				assert.NoError(t, err)
				assert.Equal(t, testUser.Username, user.Username)
				assert.Equal(t, testUser.Email, user.Email)
				assert.NotEqual(t, testUser.Password, user.Password)
				assert.NoError(t, bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(testUser.Password)))
				assert.NotNil(t, user.CreatedAt)
				assert.NotNil(t, user.UpdatedAt)
				assert.Equal(t, user.DeletedAt, gorm.DeletedAt{})
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			ur := mock.NewMockIUserRepository(mockCtrl)
			tc.buildStubs(ur)

			uu := NewUserUsecase(ur)
			user, err := uu.CreateUser(tc.args.ctx, tc.args.user)
			tc.checkResponse(t, user, err)
		})
	}

}

func Test_userUsecase_CreateUser(t *testing.T) {
	type fields struct {
		ur repository.IUserRepository
	}
	type args struct {
		ctx  context.Context
		user domain.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    domain.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uu := &userUsecase{
				ur: tt.fields.ur,
			}
			got, err := uu.CreateUser(tt.args.ctx, tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("userUsecase.CreateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userUsecase.CreateUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userUsecase_GetUser(t *testing.T) {
	type fields struct {
		ur repository.IUserRepository
	}
	type args struct {
		ctx context.Context
		id  int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    domain.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uu := &userUsecase{
				ur: tt.fields.ur,
			}
			got, err := uu.GetUser(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("userUsecase.GetUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userUsecase.GetUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userUsecase_GetUserByEmail(t *testing.T) {
	type fields struct {
		ur repository.IUserRepository
	}
	type args struct {
		ctx   context.Context
		email string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    domain.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uu := &userUsecase{
				ur: tt.fields.ur,
			}
			got, err := uu.GetUserByEmail(tt.args.ctx, tt.args.email)
			if (err != nil) != tt.wantErr {
				t.Errorf("userUsecase.GetUserByEmail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userUsecase.GetUserByEmail() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userUsecase_ListUsers(t *testing.T) {
	type fields struct {
		ur repository.IUserRepository
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []domain.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uu := &userUsecase{
				ur: tt.fields.ur,
			}
			got, err := uu.ListUsers(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("userUsecase.ListUsers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userUsecase.ListUsers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userUsecase_UpdateUser(t *testing.T) {
	type fields struct {
		ur repository.IUserRepository
	}
	type args struct {
		ctx  context.Context
		user domain.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uu := &userUsecase{
				ur: tt.fields.ur,
			}
			got, err := uu.UpdateUser(tt.args.ctx, tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("userUsecase.UpdateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("userUsecase.UpdateUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userUsecase_DeleteUser(t *testing.T) {
	type fields struct {
		ur repository.IUserRepository
	}
	type args struct {
		ctx context.Context
		id  int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uu := &userUsecase{
				ur: tt.fields.ur,
			}
			got, err := uu.DeleteUser(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("userUsecase.DeleteUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("userUsecase.DeleteUser() = %v, want %v", got, tt.want)
			}
		})
	}
}
