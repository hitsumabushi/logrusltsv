// LTSV format : http://ltsv.org/
// LTSV format is Labeled Tab-Separated Values.
// The each value has key-value pair splited by ":".
// The key can ONLY contain [0-9A-Za-z_.-].
// In this LTSV formatter, ignore key who has ":".
package logrusltsv

import (
	"bytes"
	"fmt"
	"regexp"
	"sort"
	"strconv"

	"github.com/Sirupsen/logrus"
)

type LtsvFormatter struct {
	TimestampFormat string
}

type fields [][2]string

func (f *LtsvFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	if f.TimestampFormat == "" {
		f.TimestampFormat = logrus.DefaultTimestampFormat
	}
	_, ok := entry.Data["level"]
	if ok {
		entry.Data["field.level"] = entry.Data["level"]
	}
	entry.Data["level"] = entry.Level.String()
	_, ok = entry.Data["message"]
	if ok {
		entry.Data["field.message"] = entry.Data["message"]
	}
	entry.Data["message"] = entry.Message

	items := fields{}
	for k, v := range entry.Data {
		if !isValidKey(k) {
			continue
		}
		items = append(items, [2]string{k, fmt.Sprint(v)})
	}
	sort.Sort(items)

	return items.encode(), nil
}

func (items fields) Len() int {
	return len(items)
}
func (items fields) Less(i, j int) bool {
	return items[i][0] < items[j][0]
}
func (items fields) Swap(i, j int) {
	items[i], items[j] = items[j], items[i]
}

func (items fields) encode() []byte {
	bs := bytes.Buffer{}
	for i, item := range items {
		bs.WriteString(item[0])
		bs.WriteString(":")
		bs.WriteString(escapeValue(item[1]))
		if i < items.Len()-1 {
			bs.WriteString("\t")
		} else {
			bs.WriteString("\n")
		}
	}
	return bs.Bytes()
}

// validate key: key should be consist of [0-9A-Za-z_.-]
func isValidKey(key string) bool {
	re := regexp.MustCompile("[^0-9A-Za-z_.-]")
	return !re.MatchString(key)
}

// escape filed value
func escapeValue(value string) string {
	quoted := strconv.Quote(value)
	return quoted[1 : len(quoted)-1]
}
