package main

import "github.com/ckd/ssm-parent/cmd"

var version = "dev"

func main() {
	cmd.Execute(version)
}
