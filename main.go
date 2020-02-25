package main

import (
	"fmt"
	"github.com/1iza/goproject/lib/testpkg"
)

func main() {
	m := testpkg.NewMgr()
	m.Add()
	fmt.Printf("I want to test, m:%v!\n", m.Get())
}
