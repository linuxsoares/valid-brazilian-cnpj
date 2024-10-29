package main

type valueTestCase struct {
	description string
	input       string
	expected    bool
}

var validCNPJTestCases = []valueTestCase{
	{
		description: "Valid CNPJ",
		input:       "04.252.011/0001-10",
		expected:    true,
	},
	{
		description: "Valid CNPJ",
		input:       "72.419.986/0001-40",
		expected:    true,
	},
	{
		description: "Valid CNPJ",
		input:       "12.ABC.345/01DE-35",
		expected:    true,
	},
	{
		description: "Valid CNPJ",
		input:       "36995147000106",
		expected:    true,
	},
	{
		description: "Invalid CNPJ",
		input:       "369951470001061",
		expected:    false,
	},
	{
		description: "Invalid CNPJ",
		input:       "13.ABC.345/01DE-35",
		expected:    false,
	},
	{
		description: "Invalid CNPJ",
		input:       "12.ABC.345/01DE",
		expected:    false,
	},
	{
		description: "Valid CNPJ",
		input:       "12.ABC.345/01DE-35",
		expected:    true,
	},
}
