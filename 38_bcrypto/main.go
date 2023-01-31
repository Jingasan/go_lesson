// lesson38: bcryptoによるパスワードのハッシュ値化
package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	password := []byte("password")

	// パスワードのハッシュ値化（第2引数はストレッチングの回数（4-31の範囲で指定する））
	hashed, err := bcrypt.GenerateFromPassword(password, 10)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(hashed))
	// => $2a$10$P9Zr9LES1Yv/n6k77pDy0OVwCRBeHRhHsFMQyU6GfkfpOXfHOPjgG

	// パスワードの検証（正しい場合はnil, 誤っている場合はbcrypt.ErrMismatchedHashAndPasswordを返す）
	err = bcrypt.CompareHashAndPassword(hashed, password)
	fmt.Println(err)
	// => <nil>
	err = bcrypt.CompareHashAndPassword(hashed, []byte("INCORRECT_PASSWORD"))
	fmt.Println(err == bcrypt.ErrMismatchedHashAndPassword)
	// => true
}
