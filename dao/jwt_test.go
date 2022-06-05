package dao

import (
	"fmt"
	"testing"
)

func TestJwtAuth(t *testing.T) {
	token, err := GenerateToken(4)
	if err != nil {
		panic(err)
	}
	if err != nil {
		panic(err)
	}
	userid, err := JwtAuth(token)
	if err != nil {
		panic(err)
	}
	fmt.Println(userid)
}
