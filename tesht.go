package main

import (
	"fmt"
	"github.com/edupo/tesht/junit"
	"os"
	"os/exec"
	"strings"
)

var filePath = "results.xml"

func main() {

	testSuites := junit.LoadOrCreate(filePath)

	passedCmd := strings.Join(os.Args[1:], " ")
	cmd := exec.Command("bash", "-c", passedCmd)
	execCommand(cmd, testSuites)

	testSuites.Save(filePath)
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

func execCommand(cmd *exec.Cmd, testSuites *junit.TestSuites) {

	// Creation of the test case (it also initializes time)
	testCase := junit.NewTestCase(strings.Join(cmd.Args[0:], " "))

	// Command execution
	output, err := cmd.CombinedOutput()
	testCase.Done(output, err)

	testSuite := testSuites.GetTestSuite("bash")
	testSuite.TestCases = append(testSuite.TestCases, *testCase)

	testSuites.Update()

	printError(err)
	printOutput(output)
}
