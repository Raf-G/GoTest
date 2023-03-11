package service

import (
	"example.com/m/v2/internal/domain"
	mock_domain "example.com/m/v2/internal/repository/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetStatus(t *testing.T) {
	testTable := []struct {
		name            string
		statusID        int
		returnResp      domain.Status
		returnError     error
		returnRespRepo  *domain.Status
		returnErrorRepo error
		isCheckError    bool
	}{
		{
			name:     "ok",
			statusID: 1,
			returnResp: domain.Status{
				ID:   1,
				Name: "В обработке",
			},
			returnError: nil,
			returnRespRepo: &domain.Status{
				ID:   1,
				Name: "В обработке",
			},
			returnErrorRepo: nil,
			isCheckError:    false,
		},
		{
			name:     "error",
			statusID: 1,
			returnResp: domain.Status{
				ID:   1,
				Name: "В обработке",
			},
			returnError: testErr,
			returnRespRepo: &domain.Status{
				ID:   1,
				Name: "В обработке",
			},
			returnErrorRepo: nil,
			isCheckError:    false,
		},
		{
			name:     "No found status",
			statusID: 1,
			returnResp: domain.Status{
				ID:   1,
				Name: "В обработке",
			},
			returnError:     domain.ErrStatusNotFound,
			returnRespRepo:  nil,
			returnErrorRepo: nil,
			isCheckError:    true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			mockStatusesRepository := mock_domain.NewMockStatusesStorage(c)
			statusesService := NewStatusService(mockStatusesRepository)

			mockStatusesRepository.
				EXPECT().
				GetStatus(testCase.statusID).
				Return(testCase.returnRespRepo, testCase.returnErrorRepo)

			respUser, err := statusesService.GetStatus(testCase.statusID)
			if testCase.isCheckError {
				assert.ErrorIs(t, err, testCase.returnError)
			} else {
				assert.NoError(t, err)
				assert.EqualValues(t, testCase.returnResp, respUser)
			}
		})
	}
}

func TestGetStatuses(t *testing.T) {
	testTable := []struct {
		name            string
		returnResp      []domain.Status
		returnError     error
		returnRespRepo  []domain.Status
		returnErrorRepo error
		isCheckError    bool
	}{
		{
			name: "ok",
			returnResp: []domain.Status{
				{
					ID:   1,
					Name: "В обработке",
				},
			},
			returnError: nil,
			returnRespRepo: []domain.Status{
				{
					ID:   1,
					Name: "В обработке",
				},
			},
			returnErrorRepo: nil,
			isCheckError:    false,
		},
		{
			name:            "error",
			returnResp:      nil,
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

			mockStatusesRepository := mock_domain.NewMockStatusesStorage(c)
			statusesService := NewStatusService(mockStatusesRepository)

			mockStatusesRepository.
				EXPECT().
				GetStatuses().
				Return(testCase.returnRespRepo, testCase.returnErrorRepo)

			respUser, err := statusesService.GetStatuses()
			if testCase.isCheckError {
				assert.ErrorIs(t, err, testCase.returnError)
			} else {
				assert.NoError(t, err)
				assert.EqualValues(t, testCase.returnResp, respUser)
			}
		})
	}
}
