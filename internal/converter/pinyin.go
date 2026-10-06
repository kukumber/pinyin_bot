package converter

import (
	"regexp"
)

var (
	// veToneRe matches "ve" before a tone number (e.g. "nve4")
	veToneRe = regexp.MustCompile(`ve([1-5])`)
	// aeDiphthongRe matches a/e followed by another vowel before a tone
	// number; the tone mark goes on the a/e
	aeDiphthongRe = regexp.MustCompile(`([ae])([iou])([1-5])`)
	// ouRe matches "ou" before a tone number; the tone mark goes on the o
	ouRe = regexp.MustCompile(`ou([1-5])`)
	// nasalRe matches a vowel followed by n/ng/r before a tone number;
	// the tone mark goes on the vowel. Consonant-only syllables (n, ng, r)
	// are left untouched.
	nasalRe = regexp.MustCompile(`([aeiouvü])(ng?|r)([1-5])`)
	// numberedVowelRe matches a vowel directly followed by a tone number
	numberedVowelRe = regexp.MustCompile(`([aeiouv])([1-5])`)
)

// marks maps a vowel with a tone number to the vowel with a tone mark
var marks = map[string]string{
	"a1": "ā", "a2": "á", "a3": "ǎ", "a4": "à", "a5": "a",
	"e1": "ē", "e2": "é", "e3": "ě", "e4": "è", "e5": "e",
	"i1": "ī", "i2": "í", "i3": "ǐ", "i4": "ì", "i5": "i",
	"o1": "ō", "o2": "ó", "o3": "ǒ", "o4": "ò", "o5": "o",
	"u1": "ū", "u2": "ú", "u3": "ǔ", "u4": "ù", "u5": "u",
	"v1": "ǖ", "v2": "ǘ", "v3": "ǚ", "v4": "ǜ", "v5": "ü",
}

// ConvertToPinyin converts text with tone numbers to pinyin with tone marks
func ConvertToPinyin(input string) string {
	return convertNumToMark(moveNumber(input))
}

// moveNumber moves the tone number onto the vowel that carries the tone mark
func moveNumber(raw string) string {
	raw = veToneRe.ReplaceAllString(raw, "üe$1")
	raw = aeDiphthongRe.ReplaceAllString(raw, "$1$3$2")
	raw = ouRe.ReplaceAllString(raw, "o${1}u")
	raw = nasalRe.ReplaceAllString(raw, "$1$3$2")
	return raw
}

// convertNumToMark converts the number to the pinyin mark
func convertNumToMark(raw string) string {
	return numberedVowelRe.ReplaceAllStringFunc(raw, func(match string) string {
		if replacement, ok := marks[match]; ok {
			return replacement
		}
		return match
	})
}
