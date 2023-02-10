package service

import (
	domain "example.com/m/v2/domain"
	mock_domain "example.com/m/v2/domain/mocks"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

var testErr = errors.New("test error")

func TestAddUser(t *testing.T) {
	testTable := []struct {
		name                   string
		inputUser              domain.User
		returnUser             domain.User
		returnError            error
		isCallRepositoryMethod bool
		returnErrorRepo        error
		isCheckError           bool
	}{
		{
			name: "ok",
			inputUser: domain.User{
				Login:    "testingLogin",
				Surname:  "testingSurname",
				Name:     "testungName",
				Role:     1,
				Password: "qeqwe12",
			},
			returnUser: domain.User{
				ID:       1,
				Login:    "testingLogin",
				Surname:  "testingSurname",
				Name:     "testungName",
				Role:     1,
				Password: "qeqwe12",
			},
			returnError:            nil,
			isCallRepositoryMethod: true,
			returnErrorRepo:        nil,
			isCheckError:           false,
		},
		{
			name: "no login",
			inputUser: domain.User{
				Login:    "",
				Surname:  "testingSurname",
				Name:     "testungName",
				Role:     1,
				Password: "qeqwe12",
			},
			returnUser:             domain.User{},
			returnError:            domain.ErrNoLogin,
			isCallRepositoryMethod: false,
			returnErrorRepo:        nil,
			isCheckError:           true,
		},
		{
			name: "no surname",
			inputUser: domain.User{
				Login:    "tetetet",
				Surname:  "",
				Name:     "testungName",
				Role:     1,
				Password: "qeqwe12",
			},
			returnUser:             domain.User{},
			returnError:            domain.ErrNoSurname,
			isCallRepositoryMethod: false,
			returnErrorRepo:        nil,
			isCheckError:           true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			mockUsersRepository := mock_domain.NewMockUsersStorage(c)
			usersService := NewUserService(mockUsersRepository)

			if testCase.isCallRepositoryMethod {
				mockUsersRepository.
					EXPECT().
					Add(testCase.inputUser).
					Return(testCase.returnUser, testCase.returnErrorRepo)
			}

			respUser, err := usersService.Add(testCase.inputUser)
			if testCase.isCheckError {
				assert.ErrorIs(t, err, testCase.returnError)
			} else {
				assert.NoError(t, err)
				assert.EqualValues(t, testCase.returnUser, respUser)
			}
		})
	}
}

func TestGetUser(t *testing.T) {

	testTable := []struct {
		name            string
		idUser          int
		returnError     error
		returnResp      domain.User
		returnRespRepo  *domain.User
		returnErrorRepo error
		isCheckError    bool
	}{
		{
			name:   "ok",
			idUser: 1,
			returnRespRepo: &domain.User{
				ID:       1,
				Login:    "testingLogin",
				Surname:  "testingSurname",
				Name:     "testungName",
				Role:     1,
				Password: "qeqwe12",
			},
			returnResp: domain.User{
				ID:       1,
				Login:    "testingLogin",
				Surname:  "testingSurname",
				Name:     "testungName",
				Role:     1,
				Password: "qeqwe12",
			},
			returnError:     nil,
			returnErrorRepo: nil,
			isCheckError:    false,
		},
		{
			name:            "error",
			idUser:          1,
			returnRespRepo:  nil,
			returnResp:      domain.User{},
			returnError:     testErr,
			returnErrorRepo: testErr,
			isCheckError:    true,
		},
		{
			name:            "no user",
			idUser:          1,
			returnRespRepo:  nil,
			returnResp:      domain.User{},
			returnError:     domain.ErrUserNotFound,
			returnErrorRepo: nil,
			isCheckError:    true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			mockUsersRepository := mock_domain.NewMockUsersStorage(c)
			usersService := NewUserService(mockUsersRepository)

			mockUsersRepository.
				EXPECT().
				GetUser(testCase.idUser).
				Return(testCase.returnRespRepo, testCase.returnErrorRepo)

			respUser, err := usersService.GetUser(testCase.idUser)
			if testCase.isCheckError {
				assert.ErrorIs(t, err, testCase.returnError)
			} else {
				assert.NoError(t, err)
				assert.EqualValues(t, testCase.returnResp, respUser)
			}
		})
	}
}

func TestGetUsers(t *testing.T) {

	testTable := []struct {
		name            string
		returnError     error
		returnResp      []domain.User
		returnRespRepo  []domain.User
		returnErrorRepo error
		isCheckError    bool
	}{
		{
			name: "ok",
			returnRespRepo: []domain.User{
				{ID: 1, Login: "testingLogin", Surname: "testingSurname", Name: "testungName", Role: 1, Password: "qeqwe12"},
				{ID: 2, Login: "testingLogin2", Surname: "testingSurname2", Name: "testungName2", Role: 2, Password: "qeqwe122"},
			},
			returnResp: []domain.User{
				{ID: 1, Login: "testingLogin", Surname: "testingSurname", Name: "testungName", Role: 1, Password: "qeqwe12"},
				{ID: 2, Login: "testingLogin2", Surname: "testingSurname2", Name: "testungName2", Role: 2, Password: "qeqwe122"},
			},
			returnError:     nil,
			returnErrorRepo: nil,
			isCheckError:    false,
		},
		{
			name:            "error",
			returnRespRepo:  nil,
			returnResp:      nil,
			returnError:     testErr,
			returnErrorRepo: testErr,
			isCheckError:    true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			mockUsersRepository := mock_domain.NewMockUsersStorage(c)
			usersService := NewUserService(mockUsersRepository)

			mockUsersRepository.
				EXPECT().
				GetUsers().
				Return(testCase.returnRespRepo, testCase.returnErrorRepo)

			respUsers, err := usersService.GetAll()
			if testCase.isCheckError {
				assert.ErrorIs(t, err, testCase.returnError)
				fmt.Println(assert.ErrorIs(t, err, testCase.returnError))
			} else {
				assert.NoError(t, err)
				assert.EqualValues(t, testCase.returnResp, respUsers)
			}
		})
	}
}

func TestEditUser(t *testing.T) {
	testTable := []struct {
		name                   string
		inputUser              domain.User
		returnUser             domain.User
		returnError            error
		isCallRepositoryMethod bool
		returnRepo             domain.User
		returnErrorRepo        error
		isCheckError           bool
	}{
		{
			name: "ok",
			inputUser: domain.User{
				ID:       1,
				Login:    "testingLogin",
				Surname:  "testingSurname",
				Name:     "testungName",
				Role:     1,
				Password: "qeqwe12",
			},
			returnUser: domain.User{
				ID:       1,
				Login:    "testingLogin",
				Surname:  "testingSurname2",
				Name:     "testungName2",
				Role:     1,
				Password: "qeqwe12",
			},
			returnError:            nil,
			isCallRepositoryMethod: true,
			returnRepo: domain.User{
				ID:       1,
				Login:    "testingLogin",
				Surname:  "testingSurname2",
				Name:     "testungName2",
				Role:     1,
				Password: "qeqwe12",
			},
			returnErrorRepo: nil,
			isCheckError:    false,
		},
		{
			name: "error",
			inputUser: domain.User{
				ID:       1,
				Login:    "grgrgrgr",
				Surname:  "testingSurname",
				Name:     "testungName",
				Role:     1,
				Password: "qeqwe12",
			},
			returnUser:             domain.User{},
			returnError:            domain.ErrUserNotEdited,
			isCallRepositoryMethod: true,
			returnErrorRepo:        testErr,
			isCheckError:           true,
		},
		{
			name: "no login",
			inputUser: domain.User{
				ID:       1,
				Login:    "",
				Surname:  "testingSurname",
				Name:     "testungName",
				Role:     1,
				Password: "qeqwe12",
			},
			returnUser:             domain.User{},
			returnError:            domain.ErrNoLogin,
			isCallRepositoryMethod: false,
			returnErrorRepo:        testErr,
			isCheckError:           true,
		},
		{
			name: "no surname",
			inputUser: domain.User{
				ID:       1,
				Login:    "qwqwqwwq",
				Surname:  "",
				Name:     "testungName",
				Role:     1,
				Password: "qeqwe12",
			},
			returnUser:             domain.User{},
			returnError:            domain.ErrNoSurname,
			isCallRepositoryMethod: false,
			returnErrorRepo:        testErr,
			isCheckError:           true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			mockUsersRepository := mock_domain.NewMockUsersStorage(c)
			usersService := NewUserService(mockUsersRepository)

			if testCase.isCallRepositoryMethod {
				mockUsersRepository.
					EXPECT().
					Edit(testCase.inputUser).
					Return(testCase.returnRepo, testCase.returnErrorRepo)
			}

			respUser, err := usersService.Edit(testCase.inputUser)
			if testCase.isCheckError {
				assert.ErrorIs(t, err, testCase.returnError)
			} else {
				assert.NoError(t, err)
				assert.EqualValues(t, testCase.returnUser, respUser)
			}
		})
	}
}

func TestDeleteUser(t *testing.T) {
	testTable := []struct {
		name                   string
		idUser                 int
		returnError            error
		isCallRepositoryMethod bool
		returnRepo             bool
		returnErrorRepo        error
		isCheckError           bool
	}{
		{
			name:                   "ok",
			idUser:                 1,
			returnError:            nil,
			isCallRepositoryMethod: true,
			returnRepo:             true,
			returnErrorRepo:        nil,
			isCheckError:           false,
		},
		{
			name:                   "error",
			idUser:                 1,
			returnError:            testErr,
			isCallRepositoryMethod: true,
			returnRepo:             true,
			returnErrorRepo:        testErr,
			isCheckError:           true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			mockUsersRepository := mock_domain.NewMockUsersStorage(c)
			usersService := NewUserService(mockUsersRepository)

			if testCase.isCallRepositoryMethod {
				mockUsersRepository.
					EXPECT().
					Delete(testCase.idUser).
					Return(testCase.returnRepo, testCase.returnErrorRepo)
			}

			err := usersService.Delete(testCase.idUser)
			if testCase.isCheckError {
				assert.ErrorIs(t, err, testCase.returnError)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
