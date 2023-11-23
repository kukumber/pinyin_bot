package converter

import "testing"

func TestConvertToPinyin(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{"Empty input", "", ""},
		{"Single character", "a1", "ā"},
		{"Multiple characters", "ni3 hao3", "nǐ hǎo"},
		{"Mixed case", "Ni3 Hao3", "Nǐ Hǎo"},
		{"Punctuation", "Hello, world!", "Hello, world!"},
		{"Numbers", "12345", "12345"},
		{"Special characters", "!@#$%^&*()", "!@#$%^&*()"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := ConvertToPinyin(tc.input)
			if result != tc.expected {
				t.Errorf("ConvertToPinyin(%s): expected %s, got %s", tc.input, tc.expected, result)
			}
		})
	}
}

func TestMoveNumber(t *testing.T) {
	// define the map for marks
	moved := map[string]string{
		"kan1": "ka1n", "kan2": "ka2n", "kan3": "ka3n", "kan4": "ka4n", "kan5": "ka5n",
		"bang1": "ba1ng", "bang2": "ba2ng", "bang3": "ba3ng", "bang4": "ba4ng", "bang5": "ba5ng",
		"nar1": "na1r", "nar2": "na2r", "nar3": "na3r", "nar4": "na4r", "nar5": "na5r",
		"mai1": "ma1i", "mai2": "ma2i", "mai3": "ma3i", "mai4": "ma4i", "mai5": "ma5i",
		"dao1": "da1o", "dao2": "da2o", "dao3": "da3o", "dao4": "da4o", "dao5": "da5o",
		"lei1": "le1i", "lei2": "le2i", "lei3": "le3i", "lei4": "le4i", "lei5": "le5i",
		"dou1": "do1u", "dou2": "do2u", "dou3": "do3u", "dou4": "do4u", "dou5": "do5u",
		"nve1": "nüe1", "nve2": "nüe2", "nve3": "nüe3", "nve4": "nüe4", "nve5": "nüe5",
	}

	for k, v := range moved {
		result := moveNumber(k)
		if result != v {
			t.Errorf("moveNumber(%s): expected %s, got %s", k, v, result)
		}
	}
}

func TestConvertNumToMark(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{"Basic a1", "a1", "ā"},
		{"Basic a2", "a2", "á"},
		{"Basic a3", "a3", "ǎ"},
		{"Basic a4", "a4", "à"},
		{"Basic a5", "a5", "a"},
		{"Basic a", "a", "a"},
		{"Basic e1", "e1", "ē"},
		{"Basic e2", "e2", "é"},
		{"Basic e3", "e3", "ě"},
		{"Basic e4", "e4", "è"},
		{"Basic e5", "e5", "e"},
		{"Basic e", "e", "e"},
		{"Basic i1", "i1", "ī"},
		{"Basic i2", "i2", "í"},
		{"Basic i3", "i3", "ǐ"},
		{"Basic i4", "i4", "ì"},
		{"Basic i5", "i5", "i"},
		{"Basic i", "i", "i"},
		{"Basic o1", "o1", "ō"},
		{"Basic o2", "o2", "ó"},
		{"Basic o3", "o3", "ǒ"},
		{"Basic o4", "o4", "ò"},
		{"Basic o5", "o5", "o"},
		{"Basic o", "o", "o"},
		{"Basic u1", "u1", "ū"},
		{"Basic u2", "u2", "ú"},
		{"Basic u3", "u3", "ǔ"},
		{"Basic u4", "u4", "ù"},
		{"Basic u5", "u5", "u"},
		{"Basic u", "u", "u"},
		{"Basic v1", "v1", "ǖ"},
		{"Basic v2", "v2", "ǘ"},
		{"Basic v3", "v3", "ǚ"},
		{"Basic v4", "v4", "ǜ"},
		{"Basic v5", "v5", "ü"},
		{"Basic v", "v", "v"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := convertNumToMark(tc.input)
			if result != tc.expected {
				t.Errorf("convertNumToMark(%s): expected %s, got %s", tc.input, tc.expected, result)
			}
		})
	}
}
