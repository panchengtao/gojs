package main

import (
	"fmt"
	"github.com/percentor/gojs"
)

func main() {
	ctx := gojs.NewContext()
	defer ctx.Release()

	ret, err := ctx.EvaluateScript("['hello', 'world'].join(' ')", nil, ".", 0)

	if err != nil {
		fmt.Println("Script had an error :(", err)
		return
	}

	if ret == nil {
		fmt.Println("Nothing returned...")
		return
	}

	retstr := ret.ToStringOrDie()

	fmt.Println(retstr)
}
