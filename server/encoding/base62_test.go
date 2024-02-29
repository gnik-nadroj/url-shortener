package encoding

import (
    "testing"
)

func TestBase62Encode(t *testing.T) {
    tests := []struct {
        input uint64
        output  string
    }{
        {input: 0, output: "0"},
        {input: 1, output: "1"},
        {input: 61, output: "z"},
        {input: 62, output: "10"},
        {input: 3844, output: "100"},
    }

    for _, test := range tests {
        got := Base62Encode(test.input)
        if got != test.output {
            t.Errorf("base62Encode(%d) = %s; output %s", test.input, got, test.output)
        }
    }
}
