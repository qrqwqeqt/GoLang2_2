package lab2

import (
	"bytes"
	"errors"
	"strings"

	. "gopkg.in/check.v1"
)

// Объявляем тестовую структуру
type ComputeHandlerSuite struct{}

// Регистрируем тестовую структуру с библиотекой gocheck
func init() {
	Suite(&ComputeHandlerSuite{})
}

func (s *ComputeHandlerSuite) TestCompute(c *C) {
	type testCase struct {
		name           string
		input          string
		expectedOutput string
		expectedError  error
	}

	var testTable = []testCase{
		{
			name:           "Simple input",
			input:          "4 0 4 + +",
			expectedOutput: "8",
			expectedError:  nil,
		},
		{
			name:           "Error: expression_incorrect",
			input:          "4 0 4",
			expectedOutput: "",
			expectedError:  errors.New("expression_incorrect"),
		},
		{
			name:           "Error: zero_division",
			input:          "12 78 3 0 / + *",
			expectedOutput: "",
			expectedError:  errors.New("zero_division"),
		},
	}

	for _, test := range testTable {
		input := strings.NewReader(test.input)
		output := new(bytes.Buffer)
		handler := ComputeHandler{
			Input:  input,
			Output: output,
		}

		err := handler.Compute()

		outputString := output.String()

		c.Check(outputString, Equals, test.expectedOutput, Commentf("%s: actual output %s", test.name, outputString))
		c.Check(err, DeepEquals, test.expectedError, Commentf("%s: actual error %v", test.name, err))
	}
}
