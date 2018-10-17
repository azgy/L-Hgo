package utils

import (
	"errors"
	"fmt"
)

func Throw(str string) {
	panic(errors.New(str))
}

func Try() {
	if err := recover(); err != nil {
		fmt.Println(err)
	}
}
