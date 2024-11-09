package test

import (
	"magical-crwler/services/admin"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func setupTestDB(t *testing.T) (*admin.AdminService, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create sqlmock: %v", err)
	}

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		t.Fatalf("failed to create gorm DB: %v", err)
	}

	adminService := admin.NewAdminService(gormDB)
	return adminService, mock
}

func TestAddAdmin_UserNotFound(t *testing.T) {
	adminService, mock := setupTestDB(t)
	defer mock.ExpectClose()

	userID := int64(123)

	mock.ExpectQuery(`SELECT \* FROM "users" WHERE "users"."id" = \$1 ORDER BY "users"."id" LIMIT \$2`).
		WithArgs(userID, 1).
		WillReturnError(gorm.ErrRecordNotFound)

	err := adminService.AddAdmin(userID)
	if err != gorm.ErrRecordNotFound {
		t.Errorf("expected ErrRecordNotFound, got %v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("expectations not met: %v", err)
	}
}

func TestAddAdmin_AlreadyAdmin(t *testing.T) {
	adminService, mock := setupTestDB(t)
	defer mock.ExpectClose()

	userID := int64(123)

	mock.ExpectQuery(`SELECT \* FROM "users" WHERE "users"."id" = \$1 ORDER BY "users"."id" LIMIT \$2`).
		WithArgs(userID, 1).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).AddRow(userID, "Admin"))

	err := adminService.AddAdmin(userID)
	if err == nil {
		t.Errorf("expected error indicating user is already an admin, got nil")
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("expectations not met: %v", err)
	}
}

func TestAddAdmin_Success(t *testing.T) {
	adminService, mock := setupTestDB(t)
	defer mock.ExpectClose()

	userID := int64(123)

	mock.ExpectQuery(`SELECT \* FROM "users" WHERE "users"."id" = \$1 ORDER BY "users"."id" LIMIT \$2`).
		WithArgs(userID, 1).
		WillReturnRows(sqlmock.NewRows([]string{"id", "role_name"}).AddRow(userID, "user"))

	mock.ExpectExec(`UPDATE "users" SET "role_id" = \$1 WHERE "id" = \$2`).
		WithArgs(1, userID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := adminService.AddAdmin(userID)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("expectations not met: %v", err)
	}
}

func TestRemoveAdmin_UserNotFound(t *testing.T) {
	adminService, mock := setupTestDB(t)
	defer mock.ExpectClose()

	userID := int64(123)

	mock.ExpectQuery(`SELECT \* FROM "users" WHERE "users"."id" = \$1 ORDER BY "users"."id" LIMIT \$2`).
		WithArgs(userID, 1).
		WillReturnError(gorm.ErrRecordNotFound)

	err := adminService.RemoveAdmin(userID)
	if err != gorm.ErrRecordNotFound {
		t.Errorf("expected ErrRecordNotFound, got %v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("expectations not met: %v", err)
	}
}

func TestRemoveAdmin_NotAnAdmin(t *testing.T) {
	adminService, mock := setupTestDB(t)
	defer mock.ExpectClose()

	userID := int64(123)

	mock.ExpectQuery(`SELECT \* FROM "users" WHERE "users"."id" = \$1 ORDER BY "users"."id" LIMIT \$2`).
		WithArgs(userID, 1).
		WillReturnRows(sqlmock.NewRows([]string{"id", "role_name"}).AddRow(userID, "user"))

	err := adminService.RemoveAdmin(userID)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("expectations not met: %v", err)
	}
}

func TestRemoveAdmin_Success(t *testing.T) {
	adminService, mock := setupTestDB(t)
	defer mock.ExpectClose()

	userID := int64(123)

	mock.ExpectQuery(`SELECT \* FROM "users" WHERE "users"."id" = \$1 ORDER BY "users"."id" LIMIT \$2`).
		WithArgs(userID, 1).
		WillReturnRows(sqlmock.NewRows([]string{"id", "role_name"}).AddRow(userID, "admin"))

	mock.ExpectExec(`UPDATE "users" SET "role_id" = \$1 WHERE "id" = \$2`).
		WithArgs(1, userID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := adminService.RemoveAdmin(userID)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("expectations not met: %v", err)
	}
}
