package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	password := "user_password"

	hashedPassword, err := hashPassword(password)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("原始密码:", password)
	fmt.Println("哈希密码:", hashedPassword)

	err = comparePasswords(hashedPassword, "wrong_password")
	if err != nil {
		fmt.Println("密码不匹配", err)
	} else {
		fmt.Println("密码匹配")
	}

	err = comparePasswords(hashedPassword, "user_password")
	if err != nil {
		fmt.Println("密码不匹配", err)
	} else {
		fmt.Println("密码匹配")
	}
}

func hashPassword(password string) (string, error) {
	//使用bcrypt库的GenerateFromPassword函数对密码进行哈希处理
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func comparePasswords(hashedPassword, inputPassword string) error {
	//使用bcrypt库的CompareHashAndPassword函数比较密码
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(inputPassword))
	return err
}
