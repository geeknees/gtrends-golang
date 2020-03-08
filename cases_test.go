package gtrends

var testCases = []struct {
	description string
	input       string
	expected    bool
}{
	{
		description: "empty string",
		input:       "JA",
		expected:    true,
	},
}
