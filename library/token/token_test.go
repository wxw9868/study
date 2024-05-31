package token

import (
	"fmt"
	"study/utils"
	"testing"
)

func TestMain(t *testing.T) {
	utils.MethodRuntime(func() {
		token, err := GenerateJwtToken()
		fmt.Printf("token: %s;\n err: %s\n", token, err)
		claims, err := ParseJwtToken(token)
		fmt.Printf("claims: %+v;\n err: %s\n", claims, err)
	})
}
