package logrusltsv

import (
	"testing"

	"github.com/Sirupsen/logrus"
)

func TestLtsvFormatter(t *testing.T) {
	lf := LtsvFormatter{}

	// Test Case
	fields := logrus.Fields{
		"message":              "def",
		"level":                "ijk",
		"one":                  1,
		"pi":                   3.14,
		"bool":                 true,
		"value_has_escape_seq": "\acontain\nescape\\seq\r\t\btest",
		"value_has_space":      "contain spaces",
		"value_has_colon":      "colon:value",
		"not:visible":          "should not printed",
		"not<visible":          "should not printed",
	}
	entry := logrus.WithFields(fields)
	entry.Message = "msg"
	entry.Level = logrus.InfoLevel
	// Expect
	correct := "bool:true\tfield.level:ijk\tfield.message:def\tlevel:info\tmessage:msg\tone:1\tpi:3.14\tvalue_has_colon:colon:value\tvalue_has_escape_seq:\\acontain\\nescape\\\\seq\\r\\t\\btest\tvalue_has_space:contain spaces\n"

	b, err := lf.Format(entry)
	if err != nil {
		t.Errorf("entry=%#v cannot format.", entry)
	}
	if correct != string(b) {
		t.Error("not equal")
	}

}
