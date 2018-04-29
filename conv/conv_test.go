package conv

import (
	"bufio"
	"bytes"
	"encoding/hex"
	"fmt"
	"os"
	"testing"

	"github.com/djimenez/iconv-go"
)

var testCase = []struct {
	input string
}{
	{"a"},
	{"abc"},
	{"0123456789"},
	{"我是中国人     "},
	{"~!@#$%^&*()_+{}  "},
	{"1我0是9中国人                "},
	{""},
	{"我 "},
}

const (
	fromEncoding = "utf8"
	toEncoding   = "gbk"
)

func TestConvertMsgFixLen(t *testing.T) {
	cd, err := iconv.NewConverter(fromEncoding, toEncoding)
	if err != nil {
		t.Fatal("create converter error", err, fromEncoding, "->", toEncoding)
	}
	defer cd.Close()

	dc, err := iconv.NewConverter(toEncoding, fromEncoding)
	if err != nil {
		t.Fatal("create converter error", err, toEncoding, "->", fromEncoding)
	}

	// Print hex
	stdoutDumper := hex.Dumper(os.Stdout)
	defer stdoutDumper.Close()

	for _, item := range testCase {
		fmt.Println("------------------------------", item.input)
		output, err := ConvertMsgFixLen(cd, []byte(item.input))
		if err != nil {
			t.Errorf("error:%v, input:%v, output:%v", err, item.input, output)
			return
		}
		fmt.Printf("%s->%s, input:%x -> output:%x \n", fromEncoding, toEncoding, []byte(item.input), output)

		output2, err := ConvertMsgFixLen(dc, output)
		if err != nil {
			t.Errorf("error:%v, input:%v, output:%x", err, output, output2)
			return
		}

		fmt.Printf("%s->%s, input:%x -> output:%x\n", toEncoding, fromEncoding, output, output2)

		// Check result
		if bytes.Compare([]byte(item.input), output2) != 0 {
			t.Errorf("error:%v, input:%x, output:%x", err, item.input, output2)
			return
		}
	}

}

func TestConvertMsgFixLen2(t *testing.T) {
	cd, err := iconv.NewConverter(fromEncoding, toEncoding)
	if err != nil {
		t.Fatal("create converter error", err, fromEncoding, "->", toEncoding)
	}
	defer cd.Close()

	dc, err := iconv.NewConverter(toEncoding, fromEncoding)
	if err != nil {
		t.Fatal("create converter error", err, toEncoding, "->", fromEncoding)
	}

	file, err := os.Open("../testdata/utf8")
	if err != nil {
		t.Fatal("open file failed", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Bytes()
		buf := make([]byte, len(line)*3)
		copy(buf, line)

		fmt.Println("------------------------------", buf)
		output, err := ConvertMsgFixLen(cd, buf)
		if err != nil {
			t.Errorf("error:%v, input:%v, output:%v", err, buf, output)
			return
		}
		fmt.Printf("%s->%s, input:%x -> output:%x \n", fromEncoding, toEncoding, buf, output)

		output2, err := ConvertMsgFixLen(dc, output)
		if err != nil {
			t.Errorf("error:%v, input:%v, output:%x", err, output, output2)
			return
		}

		fmt.Printf("%s->%s, input:%x -> output:%x\n", toEncoding, fromEncoding, output, output2)

		// Check result
		if bytes.Compare(buf, output2) != 0 {
			t.Errorf("error:%v, input:%x, output:%x", err, buf, output2)
			return
		}

	}
}
