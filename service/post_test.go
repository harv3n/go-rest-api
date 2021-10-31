package service

import (
	"go-rest-api/entity"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func (mock *MockRepository) Save(post *entity.Post) (*entity.Post, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(*entity.Post), args.Error(1)
}

func (mock *MockRepository) FindAll() ([]entity.Post, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.([]entity.Post), args.Error(1)
}

func TestValidateEmptyPost(t *testing.T) {
	testService := NewPostService(nil)

	err := testService.Validate(nil)

	assert.NotNil(t, err)
	assert.Equal(t, "the post is empty", err.Error())
}

func TestValidateEmptyPostTitle(t *testing.T) {
	post := entity.Post{Id: 1, Title: "", Text: "Test Text"}
	testService := NewPostService(nil)
	err := testService.Validate(&post)

	assert.NotNil(t, err)
	assert.Equal(t, "the post title is empty", err.Error())
}

func TestFindAll(t *testing.T) {
	mockRepository := new(MockRepository)

	var id int64 = 1
	post := entity.Post{Id: id, Title: "A", Text: "A"}
	mockRepository.On("FindAll").Return([]entity.Post{post}, nil)

	testService := NewPostService(mockRepository)
	result, _ := testService.FindAll()

	mockRepository.AssertExpectations(t)
	assert.Equal(t, "A", result[0].Title)
	assert.Equal(t, "A", result[0].Text)
	assert.Equal(t, id, result[0].Id)
}

func TestCreate(t *testing.T) {
	mockRepository := new(MockRepository)

	var id int64 = 1
	post := entity.Post{Id: id, Title: "A", Text: "A"}
	mockRepository.On("Save").Return(&post, nil)

	testService := NewPostService(mockRepository)
	result, err := testService.Create(&post)

	mockRepository.AssertExpectations(t)
	assert.NotNil(t, result.Id)
	assert.Equal(t, "A", result.Title)
	assert.Equal(t, "A", result.Text)
	assert.Nil(t, err)
}
