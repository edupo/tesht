package junit

import (
	"encoding/xml"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestJUnitUpdateFunctions(t *testing.T) {

	assert := assert.New(t)
	tss := TestSuites{}
	xml.Unmarshal([]byte(data), &tss)
	tss.Update() // Important step to check
	ts := tss.TestSuites[0]

	// Error count
	assert.Equal(2, tss.Errors, "TestSuites.Errors")
	assert.Equal(1, ts.Errors, "TestSuite.Errors")

	// Test counts
	assert.Equal(2, len(tss.TestSuites), "len(TestSuites.TestSuites)")
	assert.Equal(4, tss.Tests, "TestSuites.Tests")
	assert.Equal(len(ts.TestCases), tss.TestSuites[0].Tests, "TestSuite.Tests")

	// Time
	assert.Equal(2.0, tss.Time, "TestSuites.Time")
	assert.Equal(1.0, ts.Time, "TestSuite.Time")

}

const data = `
<?xml version="1.0" encoding="utf-8"?>
<testsuites>
  <testsuite name="suite1" tests="0">
    <testcase classname="tests.test1" file="tests/test.py" line="42" name="test_empty" time="0.5"/>
    <testcase classname="tests.test2" file="tests/test.pyto" line="920430243444" name="test_notempty" time="0.5">
	  <error msg="This is anonther" type="1"/>
	</testcase>
  </testsuite>
  <testsuite name="suite2" tests="0">
    <testcase classname="tests.test3" file="tests/test.py" line="42" name="test_empty" time="0.5"/>
    <testcase classname="tests.test4" file="tests/test.pyto" line="920430243444" name="test_notempty" time="0.5">
	  <error msg="This is an error" type="1"/>
	</testcase>
  </testsuite>
</testsuites>
`
