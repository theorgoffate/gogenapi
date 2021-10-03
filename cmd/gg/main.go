package main

import (
	"embed"
	"fmt"
	"os"
)

type ExitCondition int

const (
	ExitOk ExitCondition = iota
	ExitMissingCmd
	ExitMissingCmdOpt
)

//go:embed generators
var generators embed.FS

func main() {
	if len(os.Args) < 2 {
		printhelp()
		os.Exit(int(ExitMissingCmd))
	}

	var (
		result []byte
		err    error
	)

	cmd := os.Args[1]
	cmdOpt := os.Args[2:len(os.Args)]

	switch cmd {
	case "controller":
		result, err = gencontroller(cmdOpt...)
	}

	if err != nil {
		fmt.Println(err)
		os.Exit(int(ExitMissingCmdOpt))
	}
	fmt.Println(result)
}

func printhelp() {
	fmt.Println("Syntax: gg <command> <command options>")
}

func gencontroller(opt ...string) ([]byte, error) {
	if len(opt) < 1 {
		return nil, fmt.Errorf("missing controller name\nSyntax: gg controller <controller name>")
	}
	packageName := opt[0]
	fmt.Printf("packageName is: %s\n", packageName)

	return nil, nil
}
