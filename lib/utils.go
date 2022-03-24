package lib

import (
	"errors"
	"fmt"

	"github.com/lithdew/quickjs"
)

func JSError(err error, shouldPanic bool) {
	if err != nil {
		var evalErr *quickjs.Error

		if errors.As(err, &evalErr) {
			fmt.Println(evalErr.Cause)
			fmt.Println(evalErr.Stack)
		}

		if shouldPanic {
			panic(err)
		}
	}
}
