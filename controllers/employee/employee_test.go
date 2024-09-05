package employee

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/shivdaskadam/golang-boilerplate/iface"
	"github.com/shivdaskadam/golang-boilerplate/request_response/healthCheck"
	"github.com/shivdaskadam/golang-boilerplate/request_response/user"
	"github.com/shivdaskadam/golang-boilerplate/schemas"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockService is a mock implementation of the Service interface
type MockService struct {
	mock.Mock
}

// GetUser implements iface.Service.
func (m *MockService) GetUser(ctx context.Context) (res user.GetUserResponse, err error) {
	panic("unimplemented")
}

// HealthCheck implements iface.Service.
func (m *MockService) HealthCheck(ctx context.Context) (res healthCheck.Response, err error) {
	panic("unimplemented")
}

func (m *MockService) CreateEmployee(ctx context.Context, employee *schemas.Employee) error {
	args := m.Called(ctx, employee)
	return args.Error(0)
}

func (m *MockService) GetEmployee(ctx context.Context, id int) (*schemas.Employee, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*schemas.Employee), args.Error(1)
}

func (m *MockService) GetEmployees(ctx context.Context) ([]schemas.Employee, error) {
	args := m.Called(ctx)
	return args.Get(0).([]schemas.Employee), args.Error(1)
}

func (m *MockService) UpdateEmployee(ctx context.Context, employee *schemas.Employee) error {
	args := m.Called(ctx, employee)
	return args.Error(0)
}

func (m *MockService) DeleteEmployee(ctx context.Context, id int) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

// Helper function to create a new Fiber app with the employee routes
func setupApp(svc iface.Service) *fiber.App {
	app := fiber.New()
	app.Post("/employees", CreateEmployeeHandler(svc))
	app.Get("/employees/:id", GetEmployeeHandler(svc))
	app.Get("/employees", GetEmployeesHandler(svc))
	app.Put("/employees/:id", UpdateEmployeeHandler(svc))
	app.Delete("/employees/:id", DeleteEmployeeHandler(svc))
	return app
}

func TestCreateEmployeeHandler(t *testing.T) {
	mockSvc := new(MockService)
	app := setupApp(mockSvc)

	employee := &schemas.Employee{
		Id:   1,
		Name: "John Doe",
	}
	mockSvc.On("CreateEmployee", mock.Anything, employee).Return(nil)

	body, _ := json.Marshal(employee)
	req := httptest.NewRequest(http.MethodPost, "/employees", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	resp, _ := app.Test(req)

	assert.Equal(t, fiber.StatusCreated, resp.StatusCode)
	mockSvc.AssertExpectations(t)
}

func TestGetEmployeeHandler(t *testing.T) {
	mockSvc := new(MockService)
	app := setupApp(mockSvc)

	employee := &schemas.Employee{
		Id:   1,
		Name: "John Doe",
	}
	mockSvc.On("GetEmployee", mock.Anything, 1).Return(employee, nil)

	req := httptest.NewRequest(http.MethodGet, "/employees/1", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	mockSvc.AssertExpectations(t)
}

func TestGetEmployeesHandler(t *testing.T) {
	mockSvc := new(MockService)
	app := setupApp(mockSvc)

	employees := []schemas.Employee{
		{Id: 1, Name: "John Doe"},
		{Id: 2, Name: "Jane Doe"},
	}
	mockSvc.On("GetEmployees", mock.Anything).Return(employees, nil)

	req := httptest.NewRequest(http.MethodGet, "/employees", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	mockSvc.AssertExpectations(t)
}

func TestUpdateEmployeeHandler(t *testing.T) {
	mockSvc := new(MockService)
	app := setupApp(mockSvc)

	employee := &schemas.Employee{
		Id:   1,
		Name: "John Doe",
	}
	mockSvc.On("UpdateEmployee", mock.Anything, employee).Return(nil)

	body, _ := json.Marshal(employee)
	req := httptest.NewRequest(http.MethodPut, "/employees/1", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	resp, _ := app.Test(req)

	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	mockSvc.AssertExpectations(t)
}

func TestDeleteEmployeeHandler(t *testing.T) {
	mockSvc := new(MockService)
	app := setupApp(mockSvc)

	mockSvc.On("DeleteEmployee", mock.Anything, 1).Return(nil)

	req := httptest.NewRequest(http.MethodDelete, "/employees/1", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, fiber.StatusNoContent, resp.StatusCode)
	mockSvc.AssertExpectations(t)
}
