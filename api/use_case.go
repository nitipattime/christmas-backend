package api

import "github.com/gin-gonic/gin"

type UseCase struct {
	testRepo interface{}
}

type Test interface {
	gg(c *gin.Context) error
}

func New(
	test Test) *UseCase {
	return &UseCase{
		testRepo: test,
	}
}
