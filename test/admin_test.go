package test

// import (
// 	"testing"

// 	"magical-crwler/services/admin"

// 	"github.com/DATA-DOG/go-sqlmock"
// 	"github.com/stretchr/testify/assert"
// 	"gorm.io/driver/postgres"
// 	"gorm.io/gorm"
// 	"gorm.io/gorm/logger"
// )

// func setupTestDB(t *testing.T) (*admin.AdminService, sqlmock.Sqlmock) {
// 	db, mock, err := sqlmock.New()
// 	if err != nil {
// 		t.Fatalf("failed to create sqlmock: %v", err)
// 	}

// 	gormDB, err := gorm.Open(postgres.New(postgres.Config{
// 		Conn: db,
// 	}), &gorm.Config{
// 		Logger: logger.Default.LogMode(logger.Silent),
// 	})
// 	if err != nil {
// 		t.Fatalf("failed to create gorm DB: %v", err)
// 	}

// 	adminService := admin.NewAdminService(gormDB)
// 	return adminService, mock
// }

// func TestAddAdmin_UserNotFound(t *testing.T) {
// 	adminService, mock := setupTestDB(t)
// 	defer mock.ExpectClose()

// 	userID := int64(123)

// 	mock.ExpectQuery(`SELECT \* FROM "users" WHERE "users"."id" = \$1 ORDER BY "users"."id" LIMIT \$2`).
// 		WithArgs(userID, 1).
// 		WillReturnError(gorm.ErrRecordNotFound)

// 	err := adminService.AddAdmin(userID)
// 	expectedErrMsg := "user not found: 123"
// 	if err == nil || err.Error() != expectedErrMsg {
// 		t.Errorf("expected error '%s', got %v", expectedErrMsg, err)
// 	}

// 	if err := mock.ExpectationsWereMet(); err != nil {
// 		t.Errorf("expectations not met: %v", err)
// 	}
// }

// func TestAddAdmin_AlreadyAdmin(t *testing.T) {
// 	adminService, mock := setupTestDB(t)
// 	defer mock.ExpectClose()

// 	userID := int64(123)

// 	mock.ExpectQuery(`SELECT \* FROM "users" WHERE "users"."id" = \$1 ORDER BY "users"."id" LIMIT \$2`).
// 		WithArgs(userID, 1).
// 		WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).AddRow(userID, "Admin"))

// 	err := adminService.AddAdmin(userID)
// 	if err == nil {
// 		t.Errorf("expected error indicating user is already an admin, got nil")
// 	}

// 	if err := mock.ExpectationsWereMet(); err != nil {
// 		t.Errorf("expectations not met: %v", err)
// 	}
// }

// func TestAddAdmin_Success(t *testing.T) {
// 	adminService, mock := setupTestDB(t)
// 	defer mock.ExpectClose()

// 	userID := int64(123)

// 	mock.ExpectQuery(`SELECT \* FROM "users" WHERE "users"."id" = \$1 ORDER BY "users"."id" LIMIT \$2`).
// 		WithArgs(userID, 1).
// 		WillReturnRows(sqlmock.NewRows([]string{"id", "role_id"}).AddRow(userID, 2))

// 	mock.ExpectQuery(`SELECT \* FROM "roles" WHERE name = \$1 ORDER BY "roles"."id" LIMIT \$2`).
// 		WithArgs("Admin", 1).
// 		WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "Admin"))

// 	mock.ExpectBegin()

// 	mock.ExpectExec(`UPDATE\s+"users"\s+SET\s+"role_id"\s*=\s*\$1\s+WHERE\s+"id"\s*=\s*\$2`).
// 		WithArgs(1, userID).
// 		WillReturnResult(sqlmock.NewResult(1, 1))

// 	mock.ExpectCommit()

// 	err := adminService.AddAdmin(userID)
// 	if err != nil {
// 		t.Errorf("expected no error, got %v", err)
// 	}

// 	if err := mock.ExpectationsWereMet(); err != nil {
// 		t.Errorf("expectations not met: %v", err)
// 	}
// }

// func TestRemoveAdmin_UserNotFound(t *testing.T) {
// 	adminService, mock := setupTestDB(t)
// 	defer mock.ExpectClose()

// 	userID := int64(123)

// 	mock.ExpectQuery(`SELECT \* FROM "users" WHERE "users"."id" = \$1 ORDER BY "users"."id" LIMIT \$2`).
// 		WithArgs(userID, 1).
// 		WillReturnError(gorm.ErrRecordNotFound)

// 	err := adminService.RemoveAdmin(userID)
// 	expectedErrMsg := "user not found: 123"
// 	if err == nil || err.Error() != expectedErrMsg {
// 		t.Errorf("expected error '%s', got %v", expectedErrMsg, err)
// 	}

// 	if err := mock.ExpectationsWereMet(); err != nil {
// 		t.Errorf("expectations not met: %v", err)
// 	}
// }

// func TestRemoveAdmin_NotAnAdmin(t *testing.T) {
// 	adminService, mock := setupTestDB(t)
// 	defer mock.ExpectClose()

// 	userID := int64(123)

// 	mock.ExpectQuery(`SELECT \* FROM "users" WHERE "users"."id" = \$1 ORDER BY "users"."id" LIMIT \$2`).
// 		WithArgs(userID, 1).
// 		WillReturnRows(sqlmock.NewRows([]string{"id", "role_id"}).AddRow(userID, 2)) // Assume Role ID 2 for "User"

// 	mock.ExpectQuery(`SELECT \* FROM "roles" WHERE name = \$1 ORDER BY "roles"."id" LIMIT \$2`).
// 		WithArgs("User", 1).
// 		WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).AddRow(2, "User")) // Role ID 2 for "User"

// 	err := adminService.RemoveAdmin(userID)
// 	if err != nil {
// 		t.Errorf("expected no error, got %v", err)
// 	}

// 	if err := mock.ExpectationsWereMet(); err != nil {
// 		t.Errorf("expectations not met: %v", err)
// 	}
// }

// func TestRemoveAdmin_Success(t *testing.T) {
// 	adminService, mock := setupTestDB(t)
// 	defer mock.ExpectClose()

// 	userID := int64(123)

// 	mock.ExpectQuery(`SELECT \* FROM "users" WHERE "users"."id" = \$1 ORDER BY "users"."id" LIMIT \$2`).
// 		WithArgs(userID, 1).
// 		WillReturnRows(sqlmock.NewRows([]string{"id", "role_id"}).AddRow(userID, 1)) // Role ID 1 assumed for "Admin"

// 	mock.ExpectQuery(`SELECT \* FROM "roles" WHERE name = \$1 ORDER BY "roles"."id" LIMIT \$2`).
// 		WithArgs("User", 1).
// 		WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).AddRow(2, "User"))

// 	mock.ExpectBegin()

// 	mock.ExpectExec(`UPDATE\s+"users"\s+SET\s+"role_id"\s*=\s*\$1\s+WHERE\s+"id"\s*=\s*\$2`).
// 		WithArgs(2, userID).
// 		WillReturnResult(sqlmock.NewResult(1, 1))

// 	mock.ExpectCommit()

// 	err := adminService.RemoveAdmin(userID)
// 	if err != nil {
// 		t.Errorf("expected no error, got %v", err)
// 	}

// 	if err := mock.ExpectationsWereMet(); err != nil {
// 		t.Errorf("expectations not met: %v", err)
// 	}
// }

// func TestListAdmins_NoAdminsFound(t *testing.T) {
// 	adminService, mock := setupTestDB(t)
// 	defer mock.ExpectClose()

// 	// Mock role retrieval for "Admin" role
// 	adminRoleID := int64(1)
// 	mock.ExpectQuery(`SELECT \* FROM "roles" WHERE name = \$1 ORDER BY "roles"."id" LIMIT \$2`).
// 		WithArgs("Admin", 1).
// 		WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).AddRow(adminRoleID, "Admin"))

// 	// Mock an empty response for admin users
// 	mock.ExpectQuery(`SELECT \* FROM "users" WHERE role_id = \$1`).
// 		WithArgs(adminRoleID).
// 		WillReturnRows(sqlmock.NewRows([]string{"id", "first_name", "last_name", "role_id"}))

// 	// Call the ListAdmins function
// 	admins, err := adminService.ListAdmins()
// 	assert.NoError(t, err)
// 	assert.Len(t, admins, 0)

// 	// Ensure all expectations were met
// 	if err := mock.ExpectationsWereMet(); err != nil {
// 		t.Errorf("expectations not met: %v", err)
// 	}
// }

// func TestListAdmins_Success(t *testing.T) {
// 	adminService, mock := setupTestDB(t)
// 	defer mock.ExpectClose()

// 	userID1 := int64(123)
// 	userID2 := int64(124)
// 	adminRoleID := int64(1)

// 	mock.ExpectQuery(`SELECT \* FROM "roles" WHERE name = \$1 ORDER BY "roles"."id" LIMIT \$2`).
// 		WithArgs("Admin", 1).
// 		WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).AddRow(adminRoleID, "Admin"))

// 	mock.ExpectQuery(`SELECT \* FROM "users" WHERE role_id = \$1`).
// 		WithArgs(adminRoleID).
// 		WillReturnRows(sqlmock.NewRows([]string{"id", "first_name", "last_name", "role_id"}).
// 			AddRow(userID1, "Alice", "Admin", adminRoleID).
// 			AddRow(userID2, "Bob", "Admin", adminRoleID))

// 	admins, err := adminService.ListAdmins()
// 	assert.NoError(t, err)
// 	assert.Len(t, admins, 2)

// 	assert.Equal(t, uint(userID1), admins[0].ID)
// 	assert.Equal(t, "Alice", admins[0].FirstName)
// 	assert.Equal(t, "Admin", admins[0].LastName)

// 	assert.Equal(t, uint(userID2), admins[1].ID)
// 	assert.Equal(t, "Bob", admins[1].FirstName)
// 	assert.Equal(t, "Admin", admins[1].LastName)

// 	if err := mock.ExpectationsWereMet(); err != nil {
// 		t.Errorf("expectations not met: %v", err)
// 	}
// }
