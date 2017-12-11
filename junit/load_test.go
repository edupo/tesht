package junit

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestJUnitLoadFromFile(t *testing.T) {
	assert := assert.New(t)

	path1 := "test/results.xml"
	path2 := "test/results_tss.xml"

	// Logic to address `go test` executed from the module root.
	if _, err := os.Stat(path1); os.IsNotExist(err) {
		path1 = "junit/" + path1
	}

	if _, err := os.Stat(path2); os.IsNotExist(err) {
		path2 = "junit/" + path2
	}

	testSuites1, err := LoadFromFile(path1)
	assert.NoError(err)
	testSuites2, err := LoadFromFile(path2)
	assert.NoError(err)

	//fmt.Printf("%#v\n", testSuites1)
	//fmt.Printf("%#v\n", testSuites2)

	assert.EqualValues(testSuites1, testSuites2)
}
