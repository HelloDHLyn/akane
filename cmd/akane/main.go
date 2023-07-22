package main

import "github.com/hellodhlyn/akane/internal/interpreter"

func main() {
	i := interpreter.NewInterpreter()
	i.Run()
}
