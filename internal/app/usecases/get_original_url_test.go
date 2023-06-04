package usecases

import (
	"go-shortener-url/internal/app/entities"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockURLRepository struct {
	mock.Mock
}

func (m *MockURLRepository) SaveURL(url *entities.URL) error {
	args := m.Called(url)
	return args.Error(0)
}

func (m *MockURLRepository) GetURLByID(ID string) (*entities.URL, error) {
	args := m.Called(ID)
	return args.Get(0).(*entities.URL), args.Error(1)
}

func TestOriginalURLUseCase_Execute(t *testing.T) {
	repo := new(MockURLRepository)

	useCase := OriginalURLUseCase{
		Repository: repo,
	}

	id := "aWQxMjM="

	expectedURL := "http://example.com"

	repo.On("GetURLByID", "id123").Return(&entities.URL{
		ID:          "id123",
		OriginalURL: expectedURL,
	}, nil)

	resultURL, err := useCase.Execute(id)

	assert.NoError(t, err)

	assert.Equal(t, expectedURL, resultURL)

	repo.AssertCalled(t, "GetURLByID", "id123")
}

func TestOriginalURLUseCase_Execute_Error(t *testing.T) {
	repo := new(MockURLRepository)

	useCase := OriginalURLUseCase{
		Repository: repo,
	}

	id := "aWQxMjM="

	expectedURL := &entities.URL{}

	repo.On("GetURLByID", "id123").Return(expectedURL, nil)

	resultURL, err := useCase.Execute(id)

	assert.NoError(t, err)

	assert.Equal(t, "", resultURL)
}

func TestOriginalURLUseCase_Execute_Error_Decode(t *testing.T) {
	repo := new(MockURLRepository)

	useCase := OriginalURLUseCase{
		Repository: repo,
	}

	id := "aWQxMjM"

	resultURL, err := useCase.Execute(id)

	assert.Error(t, err)

	assert.Equal(t, "", resultURL)
}
