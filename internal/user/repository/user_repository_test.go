package repository

import (
	"fmt"
	"regexp"
	"testing"
	"time"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/loak155/techbranch-backend/internal/user/domain"
	"github.com/loak155/techbranch-backend/mock"
)

func testUser() *domain.User {
	return &domain.User{
		// ID:       1,
		Username: "test_user",
		Email:    "test@example.com",
		Password: "password",
	}
}

func testUser2() *domain.User {
	return &domain.User{
		Username: "test_user2",
		Email:    "test2@example.com",
		Password: "password",
	}
}

func TestCreateUser(t *testing.T) {
	testUser := testUser()

	db, mock, err := mock.NewDBMock()
	if err != nil {
		t.Errorf("Failed to initialize mock DB: %v", err)
	}

	rows := sqlmock.NewRows([]string{"id"}).AddRow(1)

	mock.ExpectBegin()
	mock.ExpectQuery(regexp.QuoteMeta(
		`INSERT INTO "users" ("username","email","password","created_at","updated_at","deleted_at") VALUES ($1,$2,$3,$4,$5,$6) RETURNING "id"`)).
		// WithArgs(user.Username, user.Email, user.Password).
		WillReturnRows(rows)
	mock.ExpectCommit()

	ur := NewUserRepository(db)
	err = ur.CreateUser(testUser)
	// assert.Equal(t, nil, err)
	if err != nil {
		t.Fatal(err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Test Create User: %v", err)
	}
}

func TestGetUser(t *testing.T) {
	testUser := testUser()

	db, mock, err := mock.NewDBMock()
	if err != nil {
		t.Errorf("Failed to initialize mock DB: %v", err)
	}

	rows := sqlmock.NewRows([]string{"id", "username", "email", "password", "created_at", "updated_at", "deleted_at"}).
		AddRow(1, testUser.Username, testUser.Email, testUser.Password, time.Now(), time.Now(), nil)

	mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "users" WHERE "users"."id" = $1 AND "users"."deleted_at" IS NULL ORDER BY "users"."id" LIMIT $2`)).
		WithArgs(1, 1).
		WillReturnRows(rows)

	ur := NewUserRepository(db)
	user, err := ur.GetUser(1)
	fmt.Println(user)
	// assert.NoError(t, err)
	// assert.Equal(t, user.Username, u.Username)
	// assert.Equal(t, user.Email, u.Email)
	// assert.Equal(t, user.Password, u.Password)
	if err != nil {
		t.Fatalf("failed to get user: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Test Find User: %v", err)
	}
}

func TestGetUserByEmail(t *testing.T) {
	testUser := testUser()

	db, mock, err := mock.NewDBMock()
	if err != nil {
		t.Errorf("Failed to initialize mock DB: %v", err)
	}

	rows := sqlmock.NewRows([]string{"id", "username", "email", "password", "created_at", "updated_at", "deleted_at"}).
		AddRow(1, testUser.Username, testUser.Email, testUser.Password, time.Now(), time.Now(), nil)

	mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "users" WHERE email=$1 AND "users"."deleted_at" IS NULL ORDER BY "users"."id" LIMIT $2`)).
		WithArgs(testUser.Email, 1).
		WillReturnRows(rows)

	ur := NewUserRepository(db)
	user, err := ur.GetUserByEmail(testUser.Email)
	fmt.Println(user)
	// assert.NoError(t, err)
	// assert.Equal(t, user.Username, u.Username)
	// assert.Equal(t, user.Email, u.Email)
	// assert.Equal(t, user.Password, u.Password)
	if err != nil {
		t.Fatalf("failed to get user by email: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Test Find User: %v", err)
	}
}

func TestListUsers(t *testing.T) {
	testUser1 := testUser()
	testUser2 := testUser2()

	db, mock, err := mock.NewDBMock()
	if err != nil {
		t.Errorf("Failed to initialize mock DB: %v", err)
	}

	rows := sqlmock.NewRows([]string{"id", "username", "email", "password", "created_at", "updated_at", "deleted_at"}).
		AddRow(1, testUser1.Username, testUser1.Email, testUser1.Password, time.Now(), time.Now(), nil).
		AddRow(2, testUser2.Username, testUser2.Email, testUser2.Password, time.Now(), time.Now(), nil)

	mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "users" WHERE "users"."deleted_at" IS NULL`)).
		WillReturnRows(rows)

	ur := NewUserRepository(db)
	users, err := ur.ListUsers()
	fmt.Println(users)
	// assert.NoError(t, err)
	// assert.Equal(t, user.Username, u.Username)
	// assert.Equal(t, user.Email, u.Email)
	// assert.Equal(t, user.Password, u.Password)
	if err != nil {
		t.Fatalf("failed to list user: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Test Find User: %v", err)
	}
}

func TestUpdateUser(t *testing.T) {
	testUser := testUser()

	db, mock, err := mock.NewDBMock()
	if err != nil {
		t.Errorf("Failed to initialize mock DB: %v", err)
	}

	rows := sqlmock.NewRows([]string{"id"}).AddRow(1)

	mock.ExpectBegin()
	mock.ExpectQuery(regexp.QuoteMeta(
		`INSERT INTO "users" ("username","email","password","created_at","updated_at","deleted_at") VALUES ($1,$2,$3,$4,$5,$6) RETURNING "id"`)).
		WillReturnRows(rows)
	mock.ExpectCommit()

	ur := NewUserRepository(db)
	err = ur.UpdateUser(testUser)
	// fmt.Println(users)
	// assert.NoError(t, err)
	// assert.Equal(t, user.Username, u.Username)
	// assert.Equal(t, user.Email, u.Email)
	// assert.Equal(t, user.Password, u.Password)
	if err != nil {
		t.Fatalf("failed to list user: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Test Find User: %v", err)
	}
}

func TestDeleteUser(t *testing.T) {
	db, mock, err := mock.NewDBMock()
	if err != nil {
		t.Errorf("Failed to initialize mock DB: %v", err)
	}

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(
		`UPDATE "users" SET "deleted_at"=$1 WHERE "users"."id" = $2 AND "users"."deleted_at" IS NULL`)).
		WithArgs(sqlmock.AnyArg(), 1).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	ur := NewUserRepository(db)
	err = ur.DeleteUser(1)
	if err != nil {
		t.Fatalf("failed to list user: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Test Find User: %v", err)
	}
}
