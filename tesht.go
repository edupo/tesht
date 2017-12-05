package main

import (
"fmt"
"os"
"os/exec"
"strings"
)

func main() {

  passed_cmd := strings.Join(os.Args[1:], " ")

  cmd := exec.Command(passed_cmd)

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
