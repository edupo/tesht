// Package junit for generation and parsing of JUnit xml files.
// The xml structures are not clearly defined anywhere so this implementation is
// based on [StackOverflow](https://stackoverflow.com/questions/4922867/junit-xml-format-specification-that-hudson-supports)
// and heavily depends on how [Jenkins  pluging](https://wiki.jenkins.io/display/JENKINS/+Plugin)
// understands the xml produced.
package junit

import (
	"encoding/xml"
)

// TestSuite container for test cases among other useful
// information.
type TestSuite struct {
	XMLName     xml.Name   `xml:"testsuite"`
	Name        string     `xml:"name,attr"`
	Tests       int        `xml:"tests,attr"`
	Time        float64    `xml:"time,attr,omitempty"`
	Failures    int        `xml:"failures,attr,omitempty"`
	Disabled    int        `xml:"disabled,attr,omitempty"`
	Errors      int        `xml:"errors,attr,omitempty"`
	Skips       int        `xml:"skips,attr,omitempty"`
	Hostname    string     `xml:"hostname,attr,omitempty"`
	TestCases   []TestCase `xml:"testcase,omitempty"`
	Propertires []Property `xml:"properties,omitempty>property,omitempty"`
	SystemOut   string     `xml:"system-out,omitempty"`
	SystemErr   string     `xml:"system-err,omitempty"`
}

// Property contains a value mapped to a key.
type Property struct {
	Name  string `xml:"name,attr"`
	Value string `xml:"value,attr"`
}

// Update JUnit data (counters, times and maybe others)
func (testSuite *TestSuite) Update() {
	// Reset TestSuite values.
	testSuite.XMLName = xml.Name{Local: "testsuite"}
	testSuite.Tests = 0
	testSuite.Failures = 0
	testSuite.Disabled = 0
	testSuite.Errors = 0
	testSuite.Skips = 0
	testSuite.Time = 0
	// Checking values again.
	for _, testCase := range testSuite.TestCases {
		if testCase.Skipped != "" {
			testSuite.Skips++
		}
		testSuite.Failures += len(testCase.Failures)
		testSuite.Errors += len(testCase.Errors)
		testSuite.Time += testCase.Time
	}
	testSuite.Tests = len(testSuite.TestCases)
}
