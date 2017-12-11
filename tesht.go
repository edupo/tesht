package main

import (
	"fmt"
	//	"github.com/edupo/tesht/junit"
	"os"
	"os/exec"
	"strings"
)

func main() {

	//testSuites, err := junit.LoadFromFile("results.xml")

	// Command creation
	passedCmd := strings.Join(os.Args[1:], " ")
	cmd := exec.Command("bash", "-c", passedCmd)

	// Creation of the test case (it also initializes time)
	//testCase := NewTestCase(passedCmd)

	// Command execution
	output, err := cmd.CombinedOutput()

	printError(err)
	printOutput(output)
}

func printError(err error) {
	if err != nil {
		os.Stderr.WriteString(fmt.Sprintf("-- Error: %s\n", err.Error()))
	}
}

func printOutput(outs []byte) {
	if len(outs) > 0 {
		fmt.Printf("-- Output: %s\n", string(outs))
	}
}
