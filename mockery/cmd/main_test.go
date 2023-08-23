package main

import (
	mocks "mockery/article/mock"
	"testing"
)

func TestGet(t *testing.T) {
	mock := mocks.NewArticleRepository(t)
	mock.On("GetById").Return()
}
