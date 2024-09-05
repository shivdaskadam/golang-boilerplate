// repository/repository_test.go
package repository

import (
	"testing"

	"github.com/shivdaskadam/golang-boilerplate/schemas"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// setupTestDB initializes a new in-memory SQLite database for testing
func setupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to connect to in-memory database: %v", err)
	}

	// AutoMigrate the schemas.Employee model
	err = db.AutoMigrate(&schemas.Employee{})
	if err != nil {
		t.Fatalf("Failed to migrate database schema: %v", err)
	}

	return db
}

// TestCreateEmployee tests the CreateEmployee method
func TestCreateEmployee(t *testing.T) {
	db := setupTestDB(t)
	repo := NewEmployeeRepository(&repository{db: db})

	employee := &schemas.Employee{Name: "John Doe"}

	err := repo.CreateEmployee(employee)
	assert.NoError(t, err)
	assert.NotZero(t, employee.Id)
}

// TestGetEmployeeByID tests the GetEmployeeByID method
func TestGetEmployeeByID(t *testing.T) {
	db := setupTestDB(t)
	repo := NewEmployeeRepository(&repository{db: db})

	employee := &schemas.Employee{Name: "John Doe"}
	_ = repo.CreateEmployee(employee)

	result, err := repo.GetEmployeeByID(employee.Id)
	assert.NoError(t, err)
	assert.Equal(t, employee.Name, result.Name)
}

// TestGetAllEmployees tests the GetAllEmployees method
func TestGetAllEmployees(t *testing.T) {
	db := setupTestDB(t)
	repo := NewEmployeeRepository(&repository{db: db})

	employees := []schemas.Employee{
		{Name: "John Doe"},
		{Name: "Jane Doe"},
	}
	for _, employee := range employees {
		_ = repo.CreateEmployee(&employee)
	}

	result, err := repo.GetAllEmployees()
	assert.NoError(t, err)
	assert.Len(t, result, len(employees))
}

// TestUpdateEmployee tests the UpdateEmployee method
func TestUpdateEmployee(t *testing.T) {
	db := setupTestDB(t)
	repo := NewEmployeeRepository(&repository{db: db})

	employee := &schemas.Employee{Name: "John Doe"}
	_ = repo.CreateEmployee(employee)

	employee.Name = "John Smith"
	err := repo.UpdateEmployee(employee)
	assert.NoError(t, err)

	result, _ := repo.GetEmployeeByID(employee.Id)
	assert.Equal(t, "John Smith", result.Name)
}

// TestDeleteEmployee tests the DeleteEmployee method
func TestDeleteEmployee(t *testing.T) {
	db := setupTestDB(t)
	repo := NewEmployeeRepository(&repository{db: db})

	employee := &schemas.Employee{Name: "John Doe"}
	_ = repo.CreateEmployee(employee)

	err := repo.DeleteEmployee(employee.Id)
	assert.NoError(t, err)

	result, err := repo.GetEmployeeByID(employee.Id)
	assert.Error(t, err)
	assert.Nil(t, result)
}
