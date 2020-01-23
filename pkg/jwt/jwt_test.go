package jwt

import (
	"fmt"
	"testing"
)

func TestCreate(t *testing.T) {
	token, err := Create("test", "test.rsa")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(token)
}