package main

import "os"

func main() {
	args := os.Args
	dir := args[1]
	env, _ := ReadDir(dir)
	exitCode := RunCmd(args[2:], env)
	os.Exit(exitCode)
}
