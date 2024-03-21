package lab2

import (
	"errors"
	"fmt"

	. "gopkg.in/check.v1"
)

type PostfixEvaluatorSuite struct{}

func init() {
	Suite(&PostfixEvaluatorSuite{})
}

func (s *PostfixEvaluatorSuite) TestEvaluatePostfix(c *C) {
	type testCase struct {
		name           string
		input          string
		expectedResult string
		expectedError  error
	}

	var testTable = []testCase{
		//Прості операції
		{
			name:           "Додавання",
			input:          "7 2 4 + +",
			expectedResult: "13",
			expectedError:  nil,
		},
		{
			name:           "Віднімання",
			input:          "128 28 -",
			expectedResult: "100",
			expectedError:  nil,
		},
		{
			name:           "Множення",
			input:          "6 6 *",
			expectedResult: "36",
			expectedError:  nil,
		},
		{
			name:           "Ділення",
			input:          "14 -7 /",
			expectedResult: "-2",
			expectedError:  nil,
		},
		{
			name:           "Піднесення до ступеню",
			input:          "11 2 ^",
			expectedResult: "121",
			expectedError:  nil,
		},
		// Складні приклади
		{
			name:           "Приклад 1",
			input:          "4 5 * 3 + 657 + -1 / -11 / 0 ^",
			expectedResult: "1",
			expectedError:  nil,
		},
		{
			name:           "Приклад 2",
			input:          "5 11 + 6 * 12 + -1 *  5 + -11 + 2 /",
			expectedResult: "-57",
			expectedError:  nil,
		},
		// Некорректні
		{
			name:           "Ділення на 0",
			input:          "10 0 /",
			expectedResult: "",
			expectedError:  errors.New("zero_division"),
		},
		// Числа з комою
		{
			name:           "Числа з комою",
			input:          "4.7 -1.2 + 2.8 *",
			expectedResult: "9.8",
			expectedError:  nil,
		},
		// Помилки
		{
			name:           "Недостатньо операндів у стеці",
			input:          "1 4 7 + / -",
			expectedResult: "",
			expectedError:  errors.New("expression_incorrect"),
		},
		{
			name:           "Недостатньо операторів у стеці",
			input:          "4 2 2 +",
			expectedResult: "",
			expectedError:  errors.New("expression_incorrect"),
		},
		{
			name:           "Немає операндів і операторів",
			input:          " ",
			expectedResult: "",
			expectedError:  errors.New("expression_incorrect"),
		},
		{
			name:           "Додатковий пробіл між операндами",
			input:          "2 6  2 - -",
			expectedResult: "-2",
			expectedError:  nil,
		},
		{
			name:           "Додатковий пробіл між операторами",
			input:          "2 6 2 - -",
			expectedResult: "-2",
			expectedError:  nil,
		},
		{
			name:           "Символ: не число або не оператор 1",
			input:          "6! 5! -",
			expectedResult: "",
			expectedError:  errors.New("invalid_operand"),
		},
		{
			name:           "Символ: не число або не оператор 2",
			input:          "40e 39 %",
			expectedResult: "",
			expectedError:  errors.New("invalid_operand"),
		},
		{
			name:           "Символ: не число та не оператор 3",
			input:          "qwe qwe qwe qwe",
			expectedResult: "",
			expectedError:  errors.New("invalid_operand"),
		},
	}

	for _, test := range testTable {
		actual, err := EvaluatePostfix(test.input)
		c.Assert(actual, Equals, test.expectedResult, Commentf("%s: actual result %s", test.name, actual))
		c.Assert(err, DeepEquals, test.expectedError, Commentf("%s: actual error %v", test.name, err))
	}
}

func (s *PostfixEvaluatorSuite) ExampleEvaluatePostfix(c *C) {
	outputExample := func(name string, exp string) {
		result, error := EvaluatePostfix(exp)
		if result != "" {
			fmt.Printf("%s: %s\n", name, result)
		}
		if error != nil {
			fmt.Printf("Error at %s: %v\n", name, error)
		}
	}
	outputExample("Ex1", "9 7 - 2 * ")
	outputExample("Ex2", "11 +")
	outputExample("Ex3", "4 -2 * 6 +")
	outputExample("Ex4", "11 2 ^")
	outputExample("Ex5", "qwer qewr qwer")

	// Output:
	// Ex1: 4
	// Error at Ex2: expression_incorrect
	// Ex3: -2
	// Ex4: 121
	// Error at Ex5: invalid_operand
}
