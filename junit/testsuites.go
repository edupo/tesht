// Package junit for generation and parsing of JUnit xml files.
// The xml structures are not clearly defined anywhere so this implementation is
// based on [StackOverflow](https://stackoverflow.com/questions/4922867/junit-xml-format-specification-that-hudson-supports)
// and heavily depends on how [Jenkins  pluging](https://wiki.jenkins.io/display/JENKINS/+Plugin)
// understands the xml produced.
package junit

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

// TestSuites is a container for a list of individual test suites.
type TestSuites struct {
	XMLName    xml.Name    `xml:"testsuites"`
	Name       string      `xml:"name,attr,omitempty"`
	Time       float64     `xml:"time,attr,omitempty"`
	Tests      int         `xml:"tests,attr,omitempty"`
	Failures   int         `xml:"failures,attr,omitempty"`
	Disabled   int         `xml:"disabled,attr,omitempty"`
	Errors     int         `xml:"errors,attr,omitempty"`
	Skips      int         `xml:"skips,attr,omitempty"`
	TestSuites []TestSuite `xml:"testsuite"`
}

/*
func (testSuites *TestSuites) DoneOk(testCaseName string, testSuiteName string) {
	// Locate the test suit required

	testSuite, err := testSuites.FindTestSuiteByName(testSuiteName)
	if err != nil && len(testSuites.TestSuites) != 0 {
		testSuite = &testSuites.TestSuites[0]
	} else {
		testSuite := NewTestCase(testCaseName)
	}
}
*/

// FindTestSuiteByName returns a reference to a TestSuite or an error.
func (testSuites *TestSuites) FindTestSuiteByName(testSuiteName string) (*TestSuite, error) {
	for i := 0; i < len(testSuites.TestSuites); i++ {
		testSuite := &testSuites.TestSuites[i]
		if testSuite.Name == testSuiteName {
			return testSuite, nil
		}
	}
	return nil, fmt.Errorf("Test suite %s not found", testSuiteName)
}

// Update JUnit data (counters, times and maybe others)
func (testSuites *TestSuites) Update() {
	testSuites.XMLName = xml.Name{Local: "testsuites"}
	for i := 0; i < len(testSuites.TestSuites); i++ {
		testSuite := &testSuites.TestSuites[i]
		testSuite.Update()
		testSuites.Tests += testSuite.Tests
		testSuites.Failures += testSuite.Failures
		testSuites.Errors += testSuite.Errors
		testSuites.Skips += testSuite.Skips
		testSuites.Time += testSuite.Time
	}
}

// LoadFromFile loads JUnit tests from a file and returns always a TestSuites
// object. Even if the JUnit file is formatted using individual TestSuite
// objects.
func LoadFromFile(path string) (*TestSuites, error) {
	testSuites := new(TestSuites)

	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	// If we unmarshal 'testsuites' then return. Everything is ok.
	if err = xml.Unmarshal(bytes, testSuites); err == nil {
		testSuites.Update()
		return testSuites, nil
	}

	// If not we attempt to unmarshal an individual 'testsuite'
	ts := TestSuite{}
	if err = xml.Unmarshal(bytes, &ts); err == nil {
		testSuites.TestSuites = append(testSuites.TestSuites, ts)
		testSuites.Update()
		return testSuites, nil
	}

	// If any of these succeed there is no further check.
	return nil, fmt.Errorf("File %s does not contain usable JUnit reports", path)
}
