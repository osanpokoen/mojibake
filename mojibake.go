package main

import (
	"flag"
	"fmt"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	var isDecode bool
	flag.BoolVar(&isDecode, "d", false, "Execution Decode function when this flag is true")
	flag.Parse()

	body, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	if len(body) < 1 {
		fmt.Fprintln(os.Stderr, "body is empty, this command required body string")
		os.Exit(1)
	}

	text := Chomp(string(body))
	var result string
	logic := Encode
	if isDecode {
		logic = Decode
	}

	result, err = logic(text)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println(result)
}

func Transform(s string, t transform.Transformer) (string, error) {
	sr := strings.NewReader(s)
	tr := transform.NewReader(sr, t)
	buf, err := ioutil.ReadAll(tr)
	if err != nil {
		return "", err
	}
	return string(buf), err
}

func Encode(s string) (string, error) {
	return Transform(s, japanese.ShiftJIS.NewDecoder())
}

func Decode(s string) (string, error) {
	return Transform(s, japanese.ShiftJIS.NewEncoder())
}

func Chomp(s string) string {
	s = strings.TrimRight(s, "\n")
	if strings.HasSuffix(s, "\r") {
		s = strings.TrimRight(s, "\r")
	}

	return s
}