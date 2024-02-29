package internal_encoding

import (
    "testing"
)

func TestBase62Encode(t *testing.T) {
    tests := []struct {
        input uint64
        output  string
    }{
        {input: 0, output: "10000"},
        {input: 1, output: "10001"},
        {input: 61, output: "1000z"},
        {input: 62, output: "10010"},
        {input: 3844, output: "10100"},
    }

    for _, test := range tests {
        got := Base62Encode(test.input)
        if got != test.output {
            t.Errorf("base62Encode(%d) = %s; output %s", test.input, got, test.output)
        }
    }
}
