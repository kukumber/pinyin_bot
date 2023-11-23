package converter

import (
	"regexp"
	"strings"
)

// ConvertToPinyin converts to pinyin
func ConvertToPinyin(input string) string {
	result := moveNumber(input)
	result = convertNumToMark(result)
	return result
}

// moveNumber moves the number from the end of the syllable
func moveNumber(raw string) string {
	moved := map[string]string{
		"n1": "1n", "n2": "2n", "n3": "3n", "n4": "4n", "n5": "5n",
		"ng1": "1ng", "ng2": "2ng", "ng3": "3ng", "ng4": "4ng", "ng5": "5ng",
		"r1": "1r", "r2": "2r", "r3": "3r", "r4": "4r", "r5": "5r",
		"ai1": "a1i", "ai2": "a2i", "ai3": "a3i", "ai4": "a4i", "ai5": "a5i",
		"ao1": "a1o", "ao2": "a2o", "ao3": "a3o", "ao4": "a4o", "ao5": "a5o",
		"ei1": "e1i", "ei2": "e2i", "ei3": "e3i", "ei4": "e4i", "ei5": "e5i",
		"ou1": "o1u", "ou2": "o2u", "ou3": "o3u", "ou4": "o4u", "ou5": "o5u",
		"ve1": "üe1", "ve2": "üe2", "ve3": "üe3", "ve4": "üe4", "ve5": "üe5",
	}

	for k, v := range moved {
		raw = strings.Replace(raw, k, v, -1)
	}

	return raw
}

// convertNumToMark converts the number to the pinyin mark
func convertNumToMark(raw string) string {
	// define the regular expression to match vowel followed by a number
	re := regexp.MustCompile(`([aeiouv])([1-5])`)

	// define the map for marks
	marks := map[string]string{
		"a1": "ā", "a2": "á", "a3": "ǎ", "a4": "à", "a5": "a",
		"e1": "ē", "e2": "é", "e3": "ě", "e4": "è", "e5": "e",
		"i1": "ī", "i2": "í", "i3": "ǐ", "i4": "ì", "i5": "i",
		"o1": "ō", "o2": "ó", "o3": "ǒ", "o4": "ò", "o5": "o",
		"u1": "ū", "u2": "ú", "u3": "ǔ", "u4": "ù", "u5": "u",
		"v1": "ǖ", "v2": "ǘ", "v3": "ǚ", "v4": "ǜ", "v5": "ü",
	}

	return re.ReplaceAllStringFunc(raw, func(match string) string {
		if replacement, ok := marks[match]; ok {
			return replacement
		}
		return match
	})
}
