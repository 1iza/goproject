package main

import (
	"fmt"
	"github.com/1iza/goproject/lib/testpkg"
	"github.com/1iza/goproject/lib/testpkg2"
)

func main() {
	m1 := testpkg.NewMgr()
	m1.Add()
	m2 := testpkg2.NewMgr()
	m2.Add()
	fmt.Printf("I want to test,  m1:%v m2:%v!\n", m1.Get(), m2.Get())
}
