package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		}, {
			input:    "Charmander Bulbasaur PIKACHU",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		// Kiểm tra độ dài của slice 'actual' so với slice 'expected'
		lenExpected := len(c.expected)
		lenActual := len(actual)
		// nếu chúng không khớp, dùng t.Errorf để in ra thông báo lỗi
		if lenActual != lenExpected {
			t.Errorf("expected %v, got %v", lenExpected, lenActual)
		}
		// và đánh rớt (fail) bài test
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			// Kiểm tra từng từ trong slice
			// nếu chúng không khớp, dùng t.Errorf để in ra thông báo lỗi
			if word != expectedWord {
				t.Errorf("FAIL")
			}
			// và đánh rớt bài test
		}
	}
}
