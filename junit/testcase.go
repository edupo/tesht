// Package junit for generation and parsing of JUnit xml files.
// The xml structures are not clearly defined anywhere so this implementation is
// based on [StackOverflow](https://stackoverflow.com/questions/4922867/junit-xml-format-specification-that-hudson-supports)
// and heavily depends on how [Jenkins  pluging](https://wiki.jenkins.io/display/JENKINS/+Plugin)
// understands the xml produced.
package junit

import (
	"encoding/xml"
	"time"
)

// TestCase main structure to hold test results
type TestCase struct {
	XMLName     xml.Name  `xml:"testcase"`
	Name        string    `xml:"name,attr"`
	ClassName   string    `xml:"classname,attr,omitempty"`
	File        string    `xml:"file,attr,omitempty"`
	Line        int       `xml:"line,attr,omitempty"`
	Time        float64   `xml:"time,attr,omitempty"`
	Skipped     string    `xml:"skipped,omitempty"`
	Errors      []Error   `xml:"error,omitempty"`
	Failures    []Error   `xml:"failure,omitempty"`
	SystemOuts  []string  `xml:"system-out,omitempty"`
	SystemErrs  []string  `xml:"system-err,omitempty"`
	InitialTime time.Time `xml:"-"`
}

// Error can contain type and/or message
type Error struct {
	Message string `xml:"message,attr,omitempty"`
	Type    string `xml:"type,attr,omitempty"`
}

// NewTestCase returns a new TestCase with the passed name
func NewTestCase(name string) *TestCase {
	testCase := new(TestCase)
	testCase.XMLName = xml.Name{Local: "testcase"}
	testCase.Name = name
	testCase.InitialTime = time.Now()
	return testCase
}

// Done function sets the parameters of the test case depending on the passed
// values
func (testCase *TestCase) Done(output []byte, err error) {
	// TODO: Consider ExitError or other error types
	testCase.Update()
	if err != nil {
		errStruct := Error{}
		errStruct.Type = err.Error()
		if len(output) > 0 {
			errStruct.Message = string(output)
		}
		testCase.Errors = append(testCase.Errors, errStruct)
	}
}

// Update test case time data. Measures the time elapsed since InitialTime filed
func (testCase *TestCase) Update() {
	d := time.Since(testCase.InitialTime)
	testCase.Time = d.Seconds()
}

// DoneOk just closes the test case measuring the execution time
func (testCase *TestCase) DoneOk() {
	testCase.Update()
}

// DoneError closes the test case appending an error to it
func (testCase *TestCase) DoneError(errorString string) {
	testCase.Errors = append(testCase.Errors, Error{Message: errorString})
	testCase.Update()
}

// DoneErrorWithLine closes the test case appending an error to it
func (testCase *TestCase) DoneErrorWithLine(errorString string, file string, line int) {
	testCase.Errors = append(testCase.Errors, Error{Message: errorString})
	testCase.Line = line
	testCase.File = file
	testCase.Update()
}
