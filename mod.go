package main

import (
	"fmt"

	"github.com/cristalhq/jwt/v5"
)

func main() {
	fmt.Println("hello mod.")
	key := []byte(`secret`)
	signer, err := jwt.NewSignerHS(jwt.HS256, key)
	fmt.Println(err)
	// checkErr(err)

	// create claims (you can create your own, see: Example_BuildUserClaims)
	claims := &jwt.RegisteredClaims{
		Audience: []string{"admin"},
		ID:       "random-unique-string",
	}

	// create a Builder
	builder := jwt.NewBuilder(signer)

	// and build a Token
	token, err := builder.Build(claims)
	fmt.Println(err)
	// checkErr(err)

	// here is token as a string
	var s string = token.String()
	fmt.Println(s)
}
