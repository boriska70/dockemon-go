package collectors

import (
	"testing"
	"strings"
)

func TestMapToArray(t *testing.T) {

	var result = []string{}

	m := make(map[string]string)

	m["one"]="foo"
	m["two"]="bar"

	result = MapToArray(m)
	var resultString = strings.Join(result, ",")

	if !strings.Contains(resultString,"one=foo") {
		println("not ok")
		t.Error("Expected to have one=foo in the result that is " + resultString)
	}

	if !strings.Contains(resultString,"two=bar") {
		println("not ok")
		t.Error("Expected to have two=bar in the result that is " + resultString)
	}
}

