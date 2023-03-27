package pkg

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestParseString(t *testing.T) {
	tests := []struct {
		input  string
		expect []string
		msg    string
	}{
		{input: "1", expect: []string{"1"}},
		{input: "2+3", expect: []string{"2", "+", "3"}},
		{input: "-1+2", expect: []string{"-1", "+", "2"}, msg: "unary minus doesn't work"},
		{input: "1-(2+3)", expect: []string{"1", "-", "(", "2", "+", "3", ")"}},
		{input: "1-(2+2)/2", expect: []string{"1", "-", "(", "2", "+", "2", ")", "/", "2"}},
		{input: "1-(2+3*(-4/5))", expect: []string{"1", "-", "(", "2", "+", "3", "*", "(", "-4", "/", "5", ")", ")"}},
	}

	assert := assert.New(t)
	for _, test := range tests {
		arr, err := ParseString(test.input)
		if err != nil {
			t.Errorf("Failed to parse %v: %s", test.input, err)
		}

		m := test.msg
		if m == "" {
			m = test.input
		}
		assert.Equal(test.expect, arr, m)
	}
}

func TestOnlyUnaryMinusIsAllowed(t *testing.T) {
	res, err := ParseString("+1 - 3")
	require.Error(t, err)
	require.Empty(t, res, "result should be empty when error")
	assert.ErrorContains(t, err, _errOnlyUnaryMinusAllowed)
}
