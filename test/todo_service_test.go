package test

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	route "todo-app/config"
	"todo-app/mocks"
	"todo-app/model"
	"todo-app/service"
)

func TestTodoService_GetTodos(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	todos := []model.Todo{
		{
			ID:          1,
			Title:       "asd",
			Description: "asd",
		},
	}

	mockRepo := mocks.NewMockTodoRepository(ctrl)
	mockRepo.EXPECT().FindAllTodos(gomock.Any(), gomock.Any()).Return(todos, nil)
	todoService := service.NewTodoService(mockRepo)

	router := route.SetupRouter(todoService)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/protected/v1/todo?pageNo=0&pageSize=3", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(
		t,
		`{"error":{"ErrorCode":"","ErrorDescription":"","ErrorMessage":""},"data":{"todos":[{"id":1,"title":"asd","description":"asd"}]}}`,
		w.Body.String(),
	)
}
