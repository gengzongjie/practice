package aaa

import (
	"bufio"
	"io"
	"os"
	"strings"
	"testing"
)

func Test_Input(t *testing.T)  {
	var inputs []string

	reader := bufio.NewReader(os.Stdin)
	for {
		in, err := reader.ReadString('\n')
		if err != nil {
			t.Error(err)
		}
		in = strings.Trim(in, "\n")
		if input, err := reader.ReadString('\n'); err != nil || err == io.EOF {
			t.Error("input error")
			break
		} else {
			inputs = append(inputs, input)
		}
	}

	t.Log(inputs)
}
