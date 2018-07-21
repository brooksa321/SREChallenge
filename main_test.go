package main

import "testing"

func TestConvert(t *testing.T) {
	for _, test := range fahrenheitTests {
		if actual := fahrenheitConversions(test.input, test.target); actual != test.expected {
			t.Errorf("Convert(%v) = %v, expected %v.",
				test.input, actual, test.expected)
		}
	}
	for _, test := range celsiusTests {
		if actual := celsiusConversions(test.input, test.target); actual != test.expected {
			t.Errorf("Convert(%v) = %v, expected %v.",
				test.input, actual, test.expected)
		}
	}
	// for _, test := range kelvinTests {
	// 	if actual := kelvinConversions(test.input, test.target); actual != test.expected {
	// 		t.Errorf("Convert(%v) = %v, expected %v.",
	// 			test.input, actual, test.expected)
	// 	}
	// }
	// for _, test := range rankineTests {
	// 	if actual := rankineConversions(test.input, test.target); actual != test.expected {
	// 		t.Errorf("Convert(%v) = %v, expected %v.",
	// 			test.input, actual, test.expected)
	// 	}
	// }
}

// func BenchmarkConvert(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		for _, test := range tests {
// 			Convert(test.input)
// 		}
// 	}
// }
