// Package junit for generation and parsing of JUnit xml files.
// The xml structures are not clearly defined anywhere so this implementation is
// based on [StackOverflow](https://stackoverflow.com/questions/4922867/junit-xml-format-specification-that-hudson-supports)
// and heavily depends on how [Jenkins  pluging](https://wiki.jenkins.io/display/JENKINS/+Plugin)
// understands the xml produced.
package junit

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
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

// SearchTestSuiteByName returns a reference to a TestSuite or an error.
func (testSuites *TestSuites) SearchTestSuiteByName(testSuiteName string) (*TestSuite, error) {
	if len(testSuites.TestSuites) == 0 {
		return nil, errors.New("No test suites to search")
	}
	for i := 0; i < len(testSuites.TestSuites); i++ {
		testSuite := &testSuites.TestSuites[i]
		if testSuite.Name == testSuiteName {
			return testSuite, nil
		}
	}
	return nil, fmt.Errorf("Test suite %s not found", testSuiteName)
}

// GetTestSuite returns a reference to a TestSuite, if it does not exists
// an empty one will be created.
func (testSuites *TestSuites) GetTestSuite(name string) *TestSuite {
	var testSuite *TestSuite
	testSuite, err := testSuites.SearchTestSuiteByName(name)
	if err == nil {
		return testSuite
	}
	testSuite = new(TestSuite)
	testSuite.Name = name
	testSuites.TestSuites = append(testSuites.TestSuites, *testSuite)
	testSuite, err = testSuites.SearchTestSuiteByName(name)
	if err != nil {
		panic(err)
	}
	return testSuite
}

// Update JUnit data (counters, times and maybe others)
func (testSuites *TestSuites) Update() {
	// Reset TestSuites values.
	testSuites.XMLName = xml.Name{Local: "testsuites"}
	testSuites.Tests = 0
	testSuites.Failures = 0
	testSuites.Errors = 0
	testSuites.Skips = 0
	testSuites.Time = 0
	// Checking values again.
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

// Load loads JUnit tests from a file and returns always a TestSuites
// object. Even if the JUnit file is formatted using individual TestSuite
// objects.
func Load(path string) (*TestSuites, error) {
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

// Save saves the JUnit into a file.
func (testSuites *TestSuites) Save(path string) error {
	output, err := xml.MarshalIndent(testSuites, "  ", "    ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(path, output, 0664)
}

//LoadOrCreate loads a TestSuites structure from a JUnit report file. If it does
//not exist it just creates a new empty structure.
func LoadOrCreate(path string) *TestSuites {
	var testSuites *TestSuites

	if _, err := os.Stat(path); os.IsNotExist(err) {
		testSuites = new(TestSuites)
		testSuites.Update()
	} else {
		testSuites, err = Load(path)
		if err != nil {
			panic(err)
		}
	}

	return testSuites
}
