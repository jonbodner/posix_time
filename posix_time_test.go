package posix_time

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestPosixToGo(t *testing.T) {
	data := []struct {
		name   string
		input  string
		output string
		errMsg string
	}{
		{
			name:   "needed",
			input:  "%d-%b-%y",
			output: "02-Jan-06",
			errMsg: "",
		},
	}
	for _, v := range data {
		t.Run(v.name, func(t *testing.T) {
			out, err := ToGo(v.input)
			var errMsg string
			if err != nil {
				errMsg = err.Error()
			}
			if diff := cmp.Diff(out, v.output); diff != "" {
				t.Error(diff)
			}
			if diff := cmp.Diff(errMsg, v.errMsg); diff != "" {
				t.Error(diff)
			}
		})
	}
}
