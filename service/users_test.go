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
		name                 string
		inputUser            domain.User
		returnUser           *domain.User
		returnError          error
		callRepositoryMethod bool
		returnErrorRepo      error
		checkError           bool
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
			returnUser: &domain.User{
				ID:       1,
				Login:    "testingLogin",
				Surname:  "testingSurname",
				Name:     "testungName",
				Role:     1,
				Password: "qeqwe12",
			},
			returnError:          nil,
			callRepositoryMethod: true,
			returnErrorRepo:      nil,
			checkError:           false,
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
			returnUser:           nil,
			returnError:          domain.ErrNoLogin,
			callRepositoryMethod: false,
			returnErrorRepo:      nil,
			checkError:           true,
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
			returnUser:           nil,
			returnError:          domain.ErrNoSurname,
			callRepositoryMethod: false,
			returnErrorRepo:      nil,
			checkError:           true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			mockUsersRepository := mock_domain.NewMockUsersStorage(c)
			usersService := NewUserService(mockUsersRepository)

			if testCase.callRepositoryMethod {
				mockUsersRepository.EXPECT().Add(testCase.inputUser).Return(testCase.returnUser, testCase.returnErrorRepo)
			}

			respUser, err := usersService.Add(testCase.inputUser)
			if testCase.checkError {
				assert.ErrorIs(t, err, testCase.returnError)
			} else {
				assert.NoError(t, err)
				assert.EqualValues(t, *testCase.returnUser, respUser)
			}

			defer c.Finish()
		})
	}
}

func TestGetUser(t *testing.T) {

	testTable := []struct {
		name                 string
		idUser               int
		returnRespRepo       *domain.User
		returnError          error
		returnResp           domain.User
		callRepositoryMethod bool
		returnErrorRepo      error
		checkError           bool
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
			returnError:          nil,
			callRepositoryMethod: true,
			returnErrorRepo:      nil,
			checkError:           false,
		},
		{
			name:                 "error",
			idUser:               1,
			returnRespRepo:       nil,
			returnResp:           domain.User{},
			returnError:          testErr,
			callRepositoryMethod: true,
			returnErrorRepo:      testErr,
			checkError:           true,
		},
		{
			name:                 "no user",
			idUser:               1,
			returnRespRepo:       nil,
			returnResp:           domain.User{},
			returnError:          domain.ErrUserNotFound,
			callRepositoryMethod: true,
			returnErrorRepo:      nil,
			checkError:           true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			mockUsersRepository := mock_domain.NewMockUsersStorage(c)
			usersService := NewUserService(mockUsersRepository)

			if testCase.callRepositoryMethod {
				mockUsersRepository.EXPECT().GetUser(testCase.idUser).Return(testCase.returnRespRepo, testCase.returnErrorRepo)
			}

			respUser, err := usersService.GetUser(testCase.idUser)
			if testCase.checkError {
				assert.ErrorIs(t, err, testCase.returnError)
			} else {
				assert.NoError(t, err)
				assert.EqualValues(t, testCase.returnResp, respUser)
			}

			defer c.Finish()
		})
	}
}

func TestGetUsers(t *testing.T) {

	testTable := []struct {
		name                 string
		returnRespRepo       []domain.User
		returnError          error
		returnResp           []domain.User
		callRepositoryMethod bool
		returnErrorRepo      error
		checkError           bool
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
			returnError:          nil,
			callRepositoryMethod: true,
			returnErrorRepo:      nil,
			checkError:           false,
		},
		{
			name:                 "error",
			returnRespRepo:       nil,
			returnResp:           nil,
			returnError:          testErr,
			callRepositoryMethod: true,
			returnErrorRepo:      testErr,
			checkError:           true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			mockUsersRepository := mock_domain.NewMockUsersStorage(c)
			usersService := NewUserService(mockUsersRepository)

			if testCase.callRepositoryMethod {
				mockUsersRepository.EXPECT().GetUsers().Return(testCase.returnRespRepo, testCase.returnErrorRepo)
			}

			respUsers, err := usersService.GetAll()
			if testCase.checkError {
				assert.ErrorIs(t, err, testCase.returnError)
				fmt.Println(assert.ErrorIs(t, err, testCase.returnError))
			} else {
				assert.NoError(t, err)
				assert.EqualValues(t, testCase.returnResp, respUsers)
			}

			defer c.Finish()
		})
	}
}

func TestEditUser(t *testing.T) {
	testTable := []struct {
		name                 string
		inputUser            domain.User
		returnUser           *domain.User
		returnError          error
		callRepositoryMethod bool
		returnRepo           *domain.User
		returnErrorRepo      error
		checkError           bool
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
			returnUser: &domain.User{
				ID:       1,
				Login:    "testingLogin",
				Surname:  "testingSurname2",
				Name:     "testungName2",
				Role:     1,
				Password: "qeqwe12",
			},
			returnError:          nil,
			callRepositoryMethod: true,
			returnRepo: &domain.User{
				ID:       1,
				Login:    "testingLogin",
				Surname:  "testingSurname2",
				Name:     "testungName2",
				Role:     1,
				Password: "qeqwe12",
			},
			returnErrorRepo: nil,
			checkError:      false,
		},
		{
			name:                 "error",
			inputUser:            domain.User{},
			returnUser:           nil,
			returnError:          domain.ErrUserNotFound,
			callRepositoryMethod: true,
			returnErrorRepo:      testErr,
			checkError:           true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			mockUsersRepository := mock_domain.NewMockUsersStorage(c)
			usersService := NewUserService(mockUsersRepository)

			if testCase.callRepositoryMethod {
				mockUsersRepository.EXPECT().Edit(testCase.inputUser).Return(testCase.returnRepo, testCase.returnErrorRepo)
			}

			respUser, err := usersService.Edit(testCase.inputUser)
			if testCase.checkError {
				assert.ErrorIs(t, err, testCase.returnError)
			} else {
				assert.NoError(t, err)
				assert.EqualValues(t, *testCase.returnUser, *respUser)
			}

			defer c.Finish()
		})
	}
}

func TestDeleteUser(t *testing.T) {
	testTable := []struct {
		name                 string
		idUser               int
		returnError          error
		callRepositoryMethod bool
		returnRepo           bool
		returnErrorRepo      error
		checkError           bool
	}{
		{
			name:                 "ok",
			idUser:               1,
			returnError:          nil,
			callRepositoryMethod: true,
			returnRepo:           true,
			returnErrorRepo:      nil,
			checkError:           false,
		},
		{
			name:                 "error",
			idUser:               1,
			returnError:          testErr,
			callRepositoryMethod: true,
			returnRepo:           true,
			returnErrorRepo:      testErr,
			checkError:           true,
		},
		{
			name:                 "false isDeleted",
			idUser:               1,
			returnError:          testErr,
			callRepositoryMethod: true,
			returnRepo:           true,
			returnErrorRepo:      testErr,
			checkError:           true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			mockUsersRepository := mock_domain.NewMockUsersStorage(c)
			usersService := NewUserService(mockUsersRepository)

			if testCase.callRepositoryMethod {
				mockUsersRepository.EXPECT().Delete(testCase.idUser).Return(testCase.returnRepo, testCase.returnErrorRepo)
			}

			err := usersService.Delete(testCase.idUser)
			if testCase.checkError {
				assert.ErrorIs(t, err, testCase.returnError)
			} else {
				assert.NoError(t, err)
			}

			defer c.Finish()
		})
	}
}
