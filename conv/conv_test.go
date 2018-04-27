package conv

import (
	"bytes"
	"testing"
)

var testCase = []struct {
	input    string
	length   int
	expected string
}{
	{"a", 1, "a"},
	{"abc", 3, "abc"},
	{"0123456789", 10, "0123456789"},
	{"我是中国人", 15, string([]byte{206, 210, 202, 199, 214, 208, 185, 250, 200, 203, 0, 0, 0, 0, 0})},
	{"~!@#$%^&*()_+{}  ", 17, string([]byte{126, 33, 64, 35, 36, 37, 94, 38, 42, 40, 41, 95, 43, 123, 125, 32, 32})},
}

func TestConvertMsgFixLen(t *testing.T) {
	for _, i := range testCase {
		output, err := ConvertMsgFixLen([]byte(i.input), i.length)
		if err != nil || bytes.Compare(output, []byte(i.expected)) != 0 {
			t.Errorf("input:%v, error:%v, output:%v, expected:%v", i.input, err, output, []byte(i.expected))
		}
	}
}
