package main

import "testing"

func TestValidCNPJ(t *testing.T) {
	for _, tc := range validCNPJTestCases {
		t.Run(tc.description, func(t *testing.T) {
			actual := ValidCNPJ(tc.input)

			if actual != tc.expected {
				t.Fatalf("validCNPJ(%+q): expected %t, actual %t", tc.input, tc.expected, actual)
			}
		})
	}
}
