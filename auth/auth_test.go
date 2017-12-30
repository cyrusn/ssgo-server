package auth_test

import (
	"fmt"
	"testing"

	"github.com/cyrusn/ssgo/auth"
)

func TestMain(t *testing.T) {
	user := auth.User{Username: "lpcyn"}
	token, err := auth.GenerateJWTToken(&user)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(token)
}
