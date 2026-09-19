package utils


import (
	"math/rand"
	"time"
)



func GenerateCode(length int) string {


	characters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"


	rand.Seed(time.Now().UnixNano())


	code := ""


	for i:=0; i<length; i++ {


		code += string(
			characters[
				rand.Intn(len(characters)),
			],
		)

	}


	return code
}