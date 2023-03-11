package service

import (
	"example.com/m/v2/internal/domain"
	mock_domain "example.com/m/v2/internal/repository/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetRole(t *testing.T) {
	testTable := []struct {
		name            string
		idRole          int
		returnResp      domain.Role
		returnError     error
		returnRespRepo  *domain.Role
		returnErrorRepo error
		isCheckError    bool
	}{
		{
			name:   "ok",
			idRole: 1,
			returnResp: domain.Role{
				ID:   1,
				Name: "administrator",
			},
			returnError: nil,
			returnRespRepo: &domain.Role{
				ID:   1,
				Name: "administrator",
			},
			returnErrorRepo: nil,
			isCheckError:    false,
		},
		{
			name:   "error",
			idRole: 1,
			returnResp: domain.Role{
				ID:   1,
				Name: "administrator",
			},
			returnError:     testErr,
			returnRespRepo:  nil,
			returnErrorRepo: testErr,
			isCheckError:    true,
		},
		{
			name:   "no found",
			idRole: 1,
			returnResp: domain.Role{
				ID:   1,
				Name: "administrator",
			},
			returnError:     domain.ErrRoleNotFound,
			returnRespRepo:  nil,
			returnErrorRepo: nil,
			isCheckError:    true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			mockRolesRepository := mock_domain.NewMockRolesStorage(c)
			rolesService := NewRoleService(mockRolesRepository)

			mockRolesRepository.
				EXPECT().
				GetRole(testCase.idRole).
				Return(testCase.returnRespRepo, testCase.returnErrorRepo)

			respRole, err := rolesService.GetRole(testCase.idRole)
			if testCase.isCheckError {
				assert.ErrorIs(t, err, testCase.returnError)
			} else {
				assert.NoError(t, err)
				assert.EqualValues(t, testCase.returnResp, respRole)
			}
		})
	}
}

func TestGetRoles(t *testing.T) {
	testTable := []struct {
		name            string
		returnResp      []domain.Role
		returnError     error
		returnRespRepo  []domain.Role
		returnErrorRepo error
		isCheckError    bool
	}{
		{
			name: "ok",
			returnResp: []domain.Role{
				{
					ID:   1,
					Name: "administrator",
				},
			},
			returnError: nil,
			returnRespRepo: []domain.Role{
				{
					ID:   1,
					Name: "administrator",
				},
			},
			returnErrorRepo: nil,
			isCheckError:    false,
		},
		{
			name: "error",
			returnResp: []domain.Role{
				{
					ID:   1,
					Name: "administrator",
				},
			},
			returnError:     testErr,
			returnRespRepo:  nil,
			returnErrorRepo: testErr,
			isCheckError:    true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			mockRolesRepository := mock_domain.NewMockRolesStorage(c)
			rolesService := NewRoleService(mockRolesRepository)

			mockRolesRepository.
				EXPECT().
				GetRoles().
				Return(testCase.returnRespRepo, testCase.returnErrorRepo)

			respRoles, err := rolesService.GetRoleAll()
			if testCase.isCheckError {
				assert.ErrorIs(t, err, testCase.returnError)
			} else {
				assert.NoError(t, err)
				assert.EqualValues(t, testCase.returnResp, respRoles)
			}
		})
	}
}
