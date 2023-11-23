package converter

import (
	"testing"
	"unicode"
)

func TestConvertToPalladyA(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "a",
			input:    "a",
			expected: "а",
		},
		{
			name:     "ai",
			input:    "ai",
			expected: "ай",
		},
		{
			name:     "an",
			input:    "an",
			expected: "ань",
		},
		{
			name:     "ang",
			input:    "ang",
			expected: "ан",
		},
		{
			name:     "ao",
			input:    "ao",
			expected: "ао",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := tt.input
			got := ConvertToPallady(input)
			if got != tt.expected {
				t.Errorf("got %s, expected %s", got, tt.expected)
			}

			// convert first letter to uppercase
			ttRunes := []rune(tt.input)
			ttRunes[0] = unicode.ToUpper(ttRunes[0])
			upInput := string(ttRunes)

			// convert first letter to uppercase
			upRunes := []rune(tt.expected)
			upRunes[0] = unicode.ToUpper(upRunes[0])
			upExp := string(upRunes)

			got = ConvertToPallady(upInput)

			if got != upExp {
				t.Errorf("got %s, expected %s", got, upExp)
			}
		})
	}
}

func TestConvertToPalladyB(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "ba",
			input:    "ba",
			expected: "ба",
		},
		{
			name:     "bai",
			input:    "bai",
			expected: "бай",
		},
		{
			name:     "ban",
			input:    "ban",
			expected: "бань",
		},
		{
			name:     "bang",
			input:    "bang",
			expected: "бан",
		},
		{
			name:     "bao",
			input:    "bao",
			expected: "бао",
		},
		{
			name:     "bei",
			input:    "bei",
			expected: "бэй",
		},
		{
			name:     "ben",
			input:    "ben",
			expected: "бэнь",
		},
		{
			name:     "beng",
			input:    "beng",
			expected: "бэн",
		},
		{
			name:     "bi",
			input:    "bi",
			expected: "би",
		},
		{
			name:     "bian",
			input:    "bian",
			expected: "бянь",
		},
		{
			name:     "biao",
			input:    "biao",
			expected: "бяо",
		},
		{
			name:     "bie",
			input:    "bie",
			expected: "бе",
		},
		{
			name:     "bin",
			input:    "bin",
			expected: "бинь",
		},
		{
			name:     "bing",
			input:    "bing",
			expected: "бин",
		},
		{
			name:     "bo",
			input:    "bo",
			expected: "бо",
		},
		{
			name:     "bu",
			input:    "bu",
			expected: "бу",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := tt.input
			got := ConvertToPallady(input)
			if got != tt.expected {
				t.Errorf("got %s, expected %s", got, tt.expected)
			}

			// convert first letter to uppercase
			ttRunes := []rune(tt.input)
			ttRunes[0] = unicode.ToUpper(ttRunes[0])
			upInput := string(ttRunes)

			// convert first letter to uppercase
			upRunes := []rune(tt.expected)
			upRunes[0] = unicode.ToUpper(upRunes[0])
			upExp := string(upRunes)

			got = ConvertToPallady(upInput)

			if got != upExp {
				t.Errorf("got %s, expected %s", got, upExp)
			}
		})
	}
}

func TestConvertToPalladyС(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "ca",
			input:    "ca",
			expected: "ца",
		},
		{
			name:     "cai",
			input:    "cai",
			expected: "цай",
		},
		{
			name:     "can",
			input:    "can",
			expected: "цань",
		},
		{
			name:     "cang",
			input:    "cang",
			expected: "цан",
		},
		{
			name:     "cao",
			input:    "cao",
			expected: "цао",
		},
		{
			name:     "ce",
			input:    "ce",
			expected: "цэ",
		},
		{
			name:     "cen",
			input:    "cen",
			expected: "цэнь",
		},
		{
			name:     "ceng",
			input:    "ceng",
			expected: "цэн",
		},
		{
			name:     "ci",
			input:    "ci",
			expected: "цы",
		},
		{
			name:     "cong",
			input:    "cong",
			expected: "цун",
		},
		{
			name:     "cou",
			input:    "cou",
			expected: "цоу",
		},
		{
			name:     "cu",
			input:    "cu",
			expected: "цу",
		},
		{
			name:     "cuan",
			input:    "cuan",
			expected: "цуань",
		},
		{
			name:     "cui",
			input:    "cui",
			expected: "цуй",
		},
		{
			name:     "cun",
			input:    "cun",
			expected: "цунь",
		},
		{
			name:     "cuo",
			input:    "cuo",
			expected: "цо",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := tt.input
			got := ConvertToPallady(input)
			if got != tt.expected {
				t.Errorf("got %s, expected %s", got, tt.expected)
			}

			// convert first letter to uppercase
			ttRunes := []rune(tt.input)
			ttRunes[0] = unicode.ToUpper(ttRunes[0])
			upInput := string(ttRunes)

			// convert first letter to uppercase
			upRunes := []rune(tt.expected)
			upRunes[0] = unicode.ToUpper(upRunes[0])
			upExp := string(upRunes)

			got = ConvertToPallady(upInput)

			if got != upExp {
				t.Errorf("got %s, expected %s", got, upExp)
			}
		})
	}
}

func TestConvertToPalladyCH(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "cha",
			input:    "cha",
			expected: "ча",
		},
		{
			name:     "chai",
			input:    "chai",
			expected: "чай",
		},
		{
			name:     "chan",
			input:    "chan",
			expected: "чань",
		},
		{
			name:     "chang",
			input:    "chang",
			expected: "чан",
		},
		{
			name:     "chao",
			input:    "chao",
			expected: "чао",
		},
		{
			name:     "che",
			input:    "che",
			expected: "чэ",
		},
		{
			name:     "chen",
			input:    "chen",
			expected: "чэнь",
		},
		{
			name:     "cheng",
			input:    "cheng",
			expected: "чэн",
		},
		{
			name:     "chi",
			input:    "chi",
			expected: "чи",
		},
		{
			name:     "chong",
			input:    "chong",
			expected: "чун",
		},
		{
			name:     "chou",
			input:    "chou",
			expected: "чоу",
		},
		{
			name:     "chu",
			input:    "chu",
			expected: "чу",
		},
		{
			name:     "chua",
			input:    "chua",
			expected: "чуа",
		},
		{
			name:     "chuai",
			input:    "chuai",
			expected: "чуай",
		},
		{
			name:     "chuan",
			input:    "chuan",
			expected: "чуань",
		},
		{
			name:     "chuang",
			input:    "chuang",
			expected: "чуан",
		},
		{
			name:     "chui",
			input:    "chui",
			expected: "чуй",
		},
		{
			name:     "chun",
			input:    "chun",
			expected: "чунь",
		},
		{
			name:     "chuo",
			input:    "chuo",
			expected: "чо",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := tt.input
			got := ConvertToPallady(input)
			if got != tt.expected {
				t.Errorf("got %s, expected %s", got, tt.expected)
			}

			// convert first letter to uppercase
			ttRunes := []rune(tt.input)
			ttRunes[0] = unicode.ToUpper(ttRunes[0])
			upInput := string(ttRunes)

			// convert first letter to uppercase
			upRunes := []rune(tt.expected)
			upRunes[0] = unicode.ToUpper(upRunes[0])
			upExp := string(upRunes)

			got = ConvertToPallady(upInput)

			if got != upExp {
				t.Errorf("got %s, expected %s", got, upExp)
			}
		})
	}
}

func TestConvertToPalladyD(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "da",
			input:    "da",
			expected: "да",
		},
		{
			name:     "dai",
			input:    "dai",
			expected: "дай",
		},
		{
			name:     "dan",
			input:    "dan",
			expected: "дань",
		},
		{
			name:     "dang",
			input:    "dang",
			expected: "дан",
		},
		{
			name:     "dao",
			input:    "dao",
			expected: "дао",
		},
		{
			name:     "de",
			input:    "de",
			expected: "дэ",
		},
		{
			name:     "dei",
			input:    "dei",
			expected: "дэй",
		},
		{
			name:     "den",
			input:    "den",
			expected: "дэнь",
		},
		{
			name:     "di",
			input:    "di",
			expected: "ди",
		},
		{
			name:     "dia",
			input:    "dia",
			expected: "дя",
		},
		{
			name:     "dian",
			input:    "dian",
			expected: "дянь",
		},
		{
			name:     "diang",
			input:    "diang",
			expected: "дян",
		},
		{
			name:     "diao",
			input:    "diao",
			expected: "дяо",
		},
		{
			name:     "die",
			input:    "die",
			expected: "де",
		},
		{
			name:     "ding",
			input:    "ding",
			expected: "дин",
		},
		{
			name:     "diu",
			input:    "diu",
			expected: "дю",
		},
		{
			name:     "dong",
			input:    "dong",
			expected: "дун",
		},
		{
			name:     "dou",
			input:    "dou",
			expected: "доу",
		},
		{
			name:     "du",
			input:    "du",
			expected: "ду",
		},
		{
			name:     "duan",
			input:    "duan",
			expected: "дуань",
		},
		{
			name:     "dui",
			input:    "dui",
			expected: "дуй",
		},
		{
			name:     "dun",
			input:    "dun",
			expected: "дунь",
		},
		{
			name:     "duo",
			input:    "duo",
			expected: "до",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := tt.input
			got := ConvertToPallady(input)
			if got != tt.expected {
				t.Errorf("got %s, expected %s", got, tt.expected)
			}

			// convert first letter to uppercase
			ttRunes := []rune(tt.input)
			ttRunes[0] = unicode.ToUpper(ttRunes[0])
			upInput := string(ttRunes)

			// convert first letter to uppercase
			upRunes := []rune(tt.expected)
			upRunes[0] = unicode.ToUpper(upRunes[0])
			upExp := string(upRunes)

			got = ConvertToPallady(upInput)

			if got != upExp {
				t.Errorf("got %s, expected %s", got, upExp)
			}
		})
	}
}

func TestConvertToPalladyE(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "e",
			input:    "e",
			expected: "э",
		},
		{
			name:     "ei",
			input:    "ei",
			expected: "эй",
		},
		{
			name:     "en",
			input:    "en",
			expected: "энь",
		},
		{
			name:     "eng",
			input:    "eng",
			expected: "эн",
		},
		{
			name:     "er",
			input:    "er",
			expected: "эр",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := tt.input
			got := ConvertToPallady(input)
			if got != tt.expected {
				t.Errorf("got %s, expected %s", got, tt.expected)
			}

			// convert first letter to uppercase
			ttRunes := []rune(tt.input)
			ttRunes[0] = unicode.ToUpper(ttRunes[0])
			upInput := string(ttRunes)

			// convert first letter to uppercase
			upRunes := []rune(tt.expected)
			upRunes[0] = unicode.ToUpper(upRunes[0])
			upExp := string(upRunes)

			got = ConvertToPallady(upInput)

			if got != upExp {
				t.Errorf("got %s, expected %s", got, upExp)
			}
		})
	}
}

func TestConvertToPalladyF(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "fa",
			input:    "fa",
			expected: "фа",
		},
		{
			name:     "fan",
			input:    "fan",
			expected: "фань",
		},
		{
			name:     "fang",
			input:    "fang",
			expected: "фан",
		},
		{
			name:     "fei",
			input:    "fei",
			expected: "фэй",
		},
		{
			name:     "fen",
			input:    "fen",
			expected: "фэнь",
		},
		{
			name:     "feng",
			input:    "feng",
			expected: "фэн",
		},
		{
			name:     "fiao",
			input:    "fiao",
			expected: "фяо",
		},
		{
			name:     "fo",
			input:    "fo",
			expected: "фо",
		},
		{
			name:     "fou",
			input:    "fou",
			expected: "фоу",
		},
		{
			name:     "fu",
			input:    "fu",
			expected: "фу",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := tt.input
			got := ConvertToPallady(input)
			if got != tt.expected {
				t.Errorf("got %s, expected %s", got, tt.expected)
			}

			// convert first letter to uppercase
			ttRunes := []rune(tt.input)
			ttRunes[0] = unicode.ToUpper(ttRunes[0])
			upInput := string(ttRunes)

			// convert first letter to uppercase
			upRunes := []rune(tt.expected)
			upRunes[0] = unicode.ToUpper(upRunes[0])
			upExp := string(upRunes)

			got = ConvertToPallady(upInput)

			if got != upExp {
				t.Errorf("got %s, expected %s", got, upExp)
			}
		})
	}
}

func TestConvertToPalladyG(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "ga",
			input:    "ga",
			expected: "га",
		},

		{
			name:     "gai",
			input:    "gai",
			expected: "гай",
		},

		{
			name:     "gan",
			input:    "gan",
			expected: "гань",
		},

		{
			name:     "gang",
			input:    "gang",
			expected: "ган",
		},

		{
			name:     "gao",
			input:    "gao",
			expected: "гао",
		},
		{
			name:     "ge",
			input:    "ge",
			expected: "гэ",
		},
		{
			name:     "gei",
			input:    "gei",
			expected: "гэй",
		},
		{
			name:     "gen",
			input:    "gen",
			expected: "гэнь",
		},
		{
			name:     "geng",
			input:    "geng",
			expected: "гэн",
		},
		{
			name:     "go",
			input:    "go",
			expected: "го",
		},
		{
			name:     "gong",
			input:    "gong",
			expected: "гун",
		},
		{
			name:     "gou",
			input:    "gou",
			expected: "гоу",
		},
		{
			name:     "gu",
			input:    "gu",
			expected: "гу",
		},
		{
			name:     "gua",
			input:    "gua",
			expected: "гуа",
		},
		{
			name:     "guai",
			input:    "guai",
			expected: "гуай",
		},
		{
			name:     "guan",
			input:    "guan",
			expected: "гуань",
		},
		{
			name:     "guang",
			input:    "guang",
			expected: "гуан",
		},
		{
			name:     "gui",
			input:    "gui",
			expected: "гуй",
		},
		{
			name:     "gun",
			input:    "gun",
			expected: "гунь",
		},
		{
			name:     "guo",
			input:    "guo",
			expected: "го",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := tt.input
			got := ConvertToPallady(input)
			if got != tt.expected {
				t.Errorf("got %s, expected %s", got, tt.expected)
			}

			// convert first letter to uppercase
			ttRunes := []rune(tt.input)
			ttRunes[0] = unicode.ToUpper(ttRunes[0])
			upInput := string(ttRunes)

			// convert first letter to uppercase
			upRunes := []rune(tt.expected)
			upRunes[0] = unicode.ToUpper(upRunes[0])
			upExp := string(upRunes)

			got = ConvertToPallady(upInput)

			if got != upExp {
				t.Errorf("got %s, expected %s", got, upExp)
			}
		})
	}
}

func TestConvertToPalladyH(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "ha",
			input:    "ha",
			expected: "ха",
		},
		{
			name:     "hai",
			input:    "hai",
			expected: "хай",
		},
		{
			name:     "han",
			input:    "han",
			expected: "хань",
		},
		{
			name:     "hang",
			input:    "hang",
			expected: "хан",
		},
		{
			name:     "hao",
			input:    "hao",
			expected: "хао",
		},
		{
			name:     "he",
			input:    "he",
			expected: "хэ",
		},
		{
			name:     "hei",
			input:    "hei",
			expected: "хэй",
		},
		{
			name:     "hen",
			input:    "hen",
			expected: "хэнь",
		},
		{
			name:     "heng",
			input:    "heng",
			expected: "хэн",
		},
		{
			name:     "hm",
			input:    "hm",
			expected: "хм",
		},
		{
			name:     "hng",
			input:    "hng",
			expected: "хн",
		},
		{
			name:     "hong",
			input:    "hong",
			expected: "хун",
		},
		{
			name:     "hou",
			input:    "hou",
			expected: "хоу",
		},
		{
			name:     "hu",
			input:    "hu",
			expected: "ху",
		},
		{
			name:     "hua",
			input:    "hua",
			expected: "хуа",
		},
		{
			name:     "huai",
			input:    "huai",
			expected: "хуай",
		},
		{
			name:     "huan",
			input:    "huan",
			expected: "хуань",
		},
		{
			name:     "huang",
			input:    "huang",
			expected: "хуан",
		},
		{
			name:     "hui",
			input:    "hui",
			expected: "хуэй",
		},
		{
			name:     "hun",
			input:    "hun",
			expected: "хунь",
		},
		{
			name:     "huo",
			input:    "huo",
			expected: "хо",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := tt.input
			got := ConvertToPallady(input)
			if got != tt.expected {
				t.Errorf("got %s, expected %s", got, tt.expected)
			}

			// convert first letter to uppercase
			ttRunes := []rune(tt.input)
			ttRunes[0] = unicode.ToUpper(ttRunes[0])
			upInput := string(ttRunes)

			// convert first letter to uppercase
			upRunes := []rune(tt.expected)
			upRunes[0] = unicode.ToUpper(upRunes[0])
			upExp := string(upRunes)

			got = ConvertToPallady(upInput)

			if got != upExp {
				t.Errorf("got %s, expected %s", got, upExp)
			}
		})
	}
}

func TestConvertToPalladyJ(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "ji",
			input:    "ji",
			expected: "цзи",
		},
		{
			name:     "jia",
			input:    "jia",
			expected: "цзя",
		},
		{
			name:     "jian",
			input:    "jian",
			expected: "цзянь",
		},
		{
			name:     "jiang",
			input:    "jiang",
			expected: "цзян",
		},
		{
			name:     "jiao",
			input:    "jiao",
			expected: "цзяо",
		},
		{
			name:     "jie",
			input:    "jie",
			expected: "цзе",
		},
		{
			name:     "jin",
			input:    "jin",
			expected: "цзинь",
		},
		{
			name:     "jing",
			input:    "jing",
			expected: "цзин",
		},
		{
			name:     "jiong",
			input:    "jiong",
			expected: "цзюн",
		},
		{
			name:     "jiu",
			input:    "jiu",
			expected: "цзю",
		},
		{
			name:     "ju",
			input:    "ju",
			expected: "цзюй",
		},
		{
			name:     "juan",
			input:    "juan",
			expected: "цзюань",
		},
		{
			name:     "jue",
			input:    "jue",
			expected: "цзюэ",
		},
		{
			name:     "jun",
			input:    "jun",
			expected: "цзюнь",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := tt.input
			got := ConvertToPallady(input)
			if got != tt.expected {
				t.Errorf("got %s, expected %s", got, tt.expected)
			}

			// convert first letter to uppercase
			ttRunes := []rune(tt.input)
			ttRunes[0] = unicode.ToUpper(ttRunes[0])
			upInput := string(ttRunes)

			// convert first letter to uppercase
			upRunes := []rune(tt.expected)
			upRunes[0] = unicode.ToUpper(upRunes[0])
			upExp := string(upRunes)

			got = ConvertToPallady(upInput)

			if got != upExp {
				t.Errorf("got %s, expected %s", got, upExp)
			}
		})
	}
}

func TestConvertToPalladyK(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "ka",
			input:    "ka",
			expected: "ка",
		},
		{
			name:     "kai",
			input:    "kai",
			expected: "кай",
		},
		{
			name:     "kan",
			input:    "kan",
			expected: "кань",
		},
		{
			name:     "kang",
			input:    "kang",
			expected: "кан",
		},
		{
			name:     "kao",
			input:    "kao",
			expected: "као",
		},
		{
			name:     "ke",
			input:    "ke",
			expected: "кэ",
		},
		{
			name:     "kei",
			input:    "kei",
			expected: "кэй",
		},
		{
			name:     "ken",
			input:    "ken",
			expected: "кэнь",
		},
		{
			name:     "keng",
			input:    "keng",
			expected: "кэн",
		},
		{
			name:     "kong",
			input:    "kong",
			expected: "кун",
		},
		{
			name:     "kou",
			input:    "kou",
			expected: "коу",
		},
		{
			name:     "ku",
			input:    "ku",
			expected: "ку",
		},
		{
			name:     "kua",
			input:    "kua",
			expected: "куа",
		},
		{
			name:     "kuai",
			input:    "kuai",
			expected: "куай",
		},
		{
			name:     "kuan",
			input:    "kuan",
			expected: "куань",
		},
		{
			name:     "kuang",
			input:    "kuang",
			expected: "куан",
		},
		{
			name:     "kui",
			input:    "kui",
			expected: "куй",
		},
		{
			name:     "kun",
			input:    "kun",
			expected: "кунь",
		},
		{
			name:     "kuo",
			input:    "kuo",
			expected: "ко",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := tt.input
			got := ConvertToPallady(input)
			if got != tt.expected {
				t.Errorf("got %s, expected %s", got, tt.expected)
			}

			// convert first letter to uppercase
			ttRunes := []rune(tt.input)
			ttRunes[0] = unicode.ToUpper(ttRunes[0])
			upInput := string(ttRunes)

			// convert first letter to uppercase
			upRunes := []rune(tt.expected)
			upRunes[0] = unicode.ToUpper(upRunes[0])
			upExp := string(upRunes)

			got = ConvertToPallady(upInput)

			if got != upExp {
				t.Errorf("got %s, expected %s", got, upExp)
			}
		})
	}
}

func TestConvertToPalladyL(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "la",
			input:    "la",
			expected: "ла",
		},
		{
			name:     "lai",
			input:    "lai",
			expected: "лай",
		},
		{
			name:     "lan",
			input:    "lan",
			expected: "лань",
		},
		{
			name:     "lang",
			input:    "lang",
			expected: "лан",
		},
		{
			name:     "lao",
			input:    "lao",
			expected: "лао",
		},
		{
			name:     "le",
			input:    "le",
			expected: "лэ",
		},
		{
			name:     "lei",
			input:    "lei",
			expected: "лэй",
		},
		{
			name:     "leng",
			input:    "leng",
			expected: "лэн",
		},
		{
			name:     "li",
			input:    "li",
			expected: "ли",
		},
		{
			name:     "lia",
			input:    "lia",
			expected: "ля",
		},
		{
			name:     "lian",
			input:    "lian",
			expected: "лянь",
		},
		{
			name:     "liang",
			input:    "liang",
			expected: "лян",
		},
		{
			name:     "liao",
			input:    "liao",
			expected: "ляо",
		},
		{
			name:     "lie",
			input:    "lie",
			expected: "ле",
		},
		{
			name:     "lin",
			input:    "lin",
			expected: "линь",
		},
		{
			name:     "ling",
			input:    "ling",
			expected: "лин",
		},
		{
			name:     "liu",
			input:    "liu",
			expected: "лю",
		},
		{
			name:     "lo",
			input:    "lo",
			expected: "ло",
		},
		{
			name:     "long",
			input:    "long",
			expected: "лун",
		},
		{
			name:     "lou",
			input:    "lou",
			expected: "лоу",
		},
		{
			name:     "lu",
			input:    "lu",
			expected: "лу",
		},
		{
			name:     "lü",
			input:    "lü",
			expected: "люй",
		},
		{
			name:     "lv",
			input:    "lv",
			expected: "люй",
		},
		{
			name:     "luan",
			input:    "luan",
			expected: "луань",
		},
		{
			name:     "lüan",
			input:    "lüan",
			expected: "люань",
		},
		{
			name:     "lvan",
			input:    "lvan",
			expected: "люань",
		},
		{
			name:     "lüe",
			input:    "lüe",
			expected: "люэ",
		},
		{
			name:     "lve",
			input:    "lve",
			expected: "люэ",
		},
		{
			name:     "lun",
			input:    "lun",
			expected: "лунь",
		},
		{
			name:     "lün",
			input:    "lün",
			expected: "люнь",
		},
		{
			name:     "lvn",
			input:    "lvn",
			expected: "люнь",
		},
		{
			name:     "luo",
			input:    "luo",
			expected: "ло",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := tt.input
			got := ConvertToPallady(input)
			if got != tt.expected {
				t.Errorf("got %s, expected %s", got, tt.expected)
			}

			// convert first letter to uppercase
			ttRunes := []rune(tt.input)
			ttRunes[0] = unicode.ToUpper(ttRunes[0])
			upInput := string(ttRunes)

			// convert first letter to uppercase
			upRunes := []rune(tt.expected)
			upRunes[0] = unicode.ToUpper(upRunes[0])
			upExp := string(upRunes)

			got = ConvertToPallady(upInput)

			if got != upExp {
				t.Errorf("got %s, expected %s", got, upExp)
			}
		})
	}
}

func TestConvertToPalladyM(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "m",
			input:    "m",
			expected: "м",
		},
		{
			name:     "ma",
			input:    "ma",
			expected: "ма",
		},
		{
			name:     "mai",
			input:    "mai",
			expected: "май",
		},
		{
			name:     "man",
			input:    "man",
			expected: "мань",
		},
		{
			name:     "mang",
			input:    "mang",
			expected: "ман",
		},
		{
			name:     "mao",
			input:    "mao",
			expected: "мао",
		},
		{
			name:     "me",
			input:    "me",
			expected: "мэ",
		},
		{
			name:     "mei",
			input:    "mei",
			expected: "мэй",
		},
		{
			name:     "men",
			input:    "men",
			expected: "мэнь",
		},
		{
			name:     "meng",
			input:    "meng",
			expected: "мэн",
		},
		{
			name:     "mi",
			input:    "mi",
			expected: "ми",
		},
		{
			name:     "mian",
			input:    "mian",
			expected: "мянь",
		},
		{
			name:     "miao",
			input:    "miao",
			expected: "мяо",
		},
		{
			name:     "mie",
			input:    "mie",
			expected: "ме",
		},
		{
			name:     "min",
			input:    "min",
			expected: "минь",
		},
		{
			name:     "ming",
			input:    "ming",
			expected: "мин",
		},
		{
			name:     "miu",
			input:    "miu",
			expected: "мю",
		},
		{
			name:     "mm",
			input:    "mm",
			expected: "мм",
		},
		{
			name:     "mo",
			input:    "mo",
			expected: "мо",
		},
		{
			name:     "mou",
			input:    "mou",
			expected: "моу",
		},
		{
			name:     "mu",
			input:    "mu",
			expected: "му",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := tt.input
			got := ConvertToPallady(input)
			if got != tt.expected {
				t.Errorf("got %s, expected %s", got, tt.expected)
			}

			// convert first letter to uppercase
			ttRunes := []rune(tt.input)
			ttRunes[0] = unicode.ToUpper(ttRunes[0])
			upInput := string(ttRunes)

			// convert first letter to uppercase
			upRunes := []rune(tt.expected)
			upRunes[0] = unicode.ToUpper(upRunes[0])
			upExp := string(upRunes)

			got = ConvertToPallady(upInput)

			if got != upExp {
				t.Errorf("got %s, expected %s", got, upExp)
			}
		})
	}
}

func TestConvertToPalladyN(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "n",
			input:    "n",
			expected: "нь",
		},
		{
			name:     "na",
			input:    "na",
			expected: "на",
		},
		{
			name:     "nai",
			input:    "nai",
			expected: "най",
		},
		{
			name:     "nan",
			input:    "nan",
			expected: "нань",
		},
		{
			name:     "nang",
			input:    "nang",
			expected: "нан",
		},
		{
			name:     "nao",
			input:    "nao",
			expected: "нао",
		},
		{
			name:     "ne",
			input:    "ne",
			expected: "нэ",
		},
		{
			name:     "nei",
			input:    "nei",
			expected: "нэй",
		},
		{
			name:     "nen",
			input:    "nen",
			expected: "нэнь",
		},
		{
			name:     "neng",
			input:    "neng",
			expected: "нэн",
		},
		{
			name:     "ng",
			input:    "ng",
			expected: "н",
		},
		{
			name:     "ni",
			input:    "ni",
			expected: "ни",
		},
		{
			name:     "nia",
			input:    "nia",
			expected: "ня",
		},
		{
			name:     "nian",
			input:    "nian",
			expected: "нянь",
		},
		{
			name:     "niang",
			input:    "niang",
			expected: "нян",
		},
		{
			name:     "niao",
			input:    "niao",
			expected: "няо",
		},
		{
			name:     "nie",
			input:    "nie",
			expected: "не",
		},
		{
			name:     "nin",
			input:    "nin",
			expected: "нинь",
		},
		{
			name:     "ning",
			input:    "ning",
			expected: "нин",
		},
		{
			name:     "niu",
			input:    "niu",
			expected: "ню",
		},
		{
			name:     "nong",
			input:    "nong",
			expected: "нун",
		},
		{
			name:     "nou",
			input:    "nou",
			expected: "ноу",
		},
		{
			name:     "nu",
			input:    "nu",
			expected: "ну",
		},
		{
			name:     "nun",
			input:    "nun",
			expected: "нунь",
		},
		{
			name:     "nü",
			input:    "nü",
			expected: "нюй",
		},
		{
			name:     "nv",
			input:    "nv",
			expected: "нюй",
		},
		{
			name:     "nuan",
			input:    "nuan",
			expected: "нуань",
		},
		{
			name:     "nüe",
			input:    "nüe",
			expected: "нюэ",
		},
		{
			name:     "nve",
			input:    "nve",
			expected: "нюэ",
		},
		{
			name:     "nuo",
			input:    "nuo",
			expected: "но",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := tt.input
			got := ConvertToPallady(input)
			if got != tt.expected {
				t.Errorf("got %s, expected %s", got, tt.expected)
			}

			// convert first letter to uppercase
			ttRunes := []rune(tt.input)
			ttRunes[0] = unicode.ToUpper(ttRunes[0])
			upInput := string(ttRunes)

			// convert first letter to uppercase
			upRunes := []rune(tt.expected)
			upRunes[0] = unicode.ToUpper(upRunes[0])
			upExp := string(upRunes)

			got = ConvertToPallady(upInput)

			if got != upExp {
				t.Errorf("got %s, expected %s", got, upExp)
			}
		})
	}
}

func TestConvertToPalladyO(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "o",
			input:    "o",
			expected: "о",
		},
		{
			name:     "ou",
			input:    "ou",
			expected: "оу",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := tt.input
			got := ConvertToPallady(input)
			if got != tt.expected {
				t.Errorf("got %s, expected %s", got, tt.expected)
			}

			// convert first letter to uppercase
			ttRunes := []rune(tt.input)
			ttRunes[0] = unicode.ToUpper(ttRunes[0])
			upInput := string(ttRunes)

			// convert first letter to uppercase
			upRunes := []rune(tt.expected)
			upRunes[0] = unicode.ToUpper(upRunes[0])
			upExp := string(upRunes)

			got = ConvertToPallady(upInput)

			if got != upExp {
				t.Errorf("got %s, expected %s", got, upExp)
			}
		})
	}
}

func TestConvertToPalladyP(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "pa",
			input:    "pa",
			expected: "па",
		},
		{
			name:     "pai",
			input:    "pai",
			expected: "пай",
		},
		{
			name:     "pan",
			input:    "pan",
			expected: "пань",
		},
		{
			name:     "pang",
			input:    "pang",
			expected: "пан",
		},
		{
			name:     "pao",
			input:    "pao",
			expected: "пао",
		},
		{
			name:     "pei",
			input:    "pei",
			expected: "пэй",
		},
		{
			name:     "pen",
			input:    "pen",
			expected: "пэнь",
		},
		{
			name:     "peng",
			input:    "peng",
			expected: "пэн",
		},
		{
			name:     "pi",
			input:    "pi",
			expected: "пи",
		},
		{
			name:     "pian",
			input:    "pian",
			expected: "пянь",
		},
		{
			name:     "piang",
			input:    "piang",
			expected: "пян",
		},
		{
			name:     "piao",
			input:    "piao",
			expected: "пяо",
		},
		{
			name:     "pie",
			input:    "pie",
			expected: "пе",
		},
		{
			name:     "pin",
			input:    "pin",
			expected: "пинь",
		},
		{
			name:     "ping",
			input:    "ping",
			expected: "пин",
		},
		{
			name:     "po",
			input:    "po",
			expected: "по",
		},
		{
			name:     "pou",
			input:    "pou",
			expected: "поу",
		},
		{
			name:     "pu",
			input:    "pu",
			expected: "пу",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := tt.input
			got := ConvertToPallady(input)
			if got != tt.expected {
				t.Errorf("got %s, expected %s", got, tt.expected)
			}

			// convert first letter to uppercase
			ttRunes := []rune(tt.input)
			ttRunes[0] = unicode.ToUpper(ttRunes[0])
			upInput := string(ttRunes)

			// convert first letter to uppercase
			upRunes := []rune(tt.expected)
			upRunes[0] = unicode.ToUpper(upRunes[0])
			upExp := string(upRunes)

			got = ConvertToPallady(upInput)

			if got != upExp {
				t.Errorf("got %s, expected %s", got, upExp)
			}
		})
	}
}

func TestConvertToPalladyQ(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "qi",
			input:    "qi",
			expected: "ци",
		},
		{
			name:     "qia",
			input:    "qia",
			expected: "ця",
		},
		{
			name:     "qian",
			input:    "qian",
			expected: "цянь",
		},
		{
			name:     "qiang",
			input:    "qiang",
			expected: "цян",
		},
		{
			name:     "qiao",
			input:    "qiao",
			expected: "цяо",
		},
		{
			name:     "qie",
			input:    "qie",
			expected: "це",
		},
		{
			name:     "qin",
			input:    "qin",
			expected: "цинь",
		},
		{
			name:     "qing",
			input:    "qing",
			expected: "цин",
		},
		{
			name:     "qiong",
			input:    "qiong",
			expected: "цюн",
		},
		{
			name:     "qiu",
			input:    "qiu",
			expected: "цю",
		},
		{
			name:     "qu",
			input:    "qu",
			expected: "цюй",
		},
		{
			name:     "quan",
			input:    "quan",
			expected: "цюань",
		},
		{
			name:     "que",
			input:    "que",
			expected: "цюэ",
		},
		{
			name:     "qun",
			input:    "qun",
			expected: "цюнь",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := tt.input
			got := ConvertToPallady(input)
			if got != tt.expected {
				t.Errorf("got %s, expected %s", got, tt.expected)
			}

			// convert first letter to uppercase
			ttRunes := []rune(tt.input)
			ttRunes[0] = unicode.ToUpper(ttRunes[0])
			upInput := string(ttRunes)

			// convert first letter to uppercase
			upRunes := []rune(tt.expected)
			upRunes[0] = unicode.ToUpper(upRunes[0])
			upExp := string(upRunes)

			got = ConvertToPallady(upInput)

			if got != upExp {
				t.Errorf("got %s, expected %s", got, upExp)
			}
		})
	}
}

func TestConvertToPalladyR(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "ran",
			input:    "ran",
			expected: "жань",
		},
		{
			name:     "rang",
			input:    "rang",
			expected: "жан",
		},
		{
			name:     "rao",
			input:    "rao",
			expected: "жао",
		},
		{
			name:     "re",
			input:    "re",
			expected: "жэ",
		},
		{
			name:     "rem",
			input:    "rem",
			expected: "жэм",
		},
		{
			name:     "ren",
			input:    "ren",
			expected: "жэнь",
		},
		{
			name:     "reng",
			input:    "reng",
			expected: "жэн",
		},
		{
			name:     "ri",
			input:    "ri",
			expected: "жи",
		},
		{
			name:     "rong",
			input:    "rong",
			expected: "жун",
		},
		{
			name:     "rou",
			input:    "rou",
			expected: "жоу",
		},
		{
			name:     "ru",
			input:    "ru",
			expected: "жу",
		},
		{
			name:     "rua",
			input:    "rua",
			expected: "жуа",
		},
		{
			name:     "ruan",
			input:    "ruan",
			expected: "жуань",
		},
		{
			name:     "rui",
			input:    "rui",
			expected: "жуй",
		},
		{
			name:     "run",
			input:    "run",
			expected: "жунь",
		},
		{
			name:     "ruo",
			input:    "ruo",
			expected: "жо",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := tt.input
			got := ConvertToPallady(input)
			if got != tt.expected {
				t.Errorf("got %s, expected %s", got, tt.expected)
			}

			// convert first letter to uppercase
			ttRunes := []rune(tt.input)
			ttRunes[0] = unicode.ToUpper(ttRunes[0])
			upInput := string(ttRunes)

			// convert first letter to uppercase
			upRunes := []rune(tt.expected)
			upRunes[0] = unicode.ToUpper(upRunes[0])
			upExp := string(upRunes)

			got = ConvertToPallady(upInput)

			if got != upExp {
				t.Errorf("got %s, expected %s", got, upExp)
			}
		})
	}
}

func TestConvertToPalladyS(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "sa",
			input:    "sa",
			expected: "са",
		},
		{
			name:     "sai",
			input:    "sai",
			expected: "сай",
		},
		{
			name:     "san",
			input:    "san",
			expected: "сань",
		},
		{
			name:     "sang",
			input:    "sang",
			expected: "сан",
		},
		{
			name:     "sao",
			input:    "sao",
			expected: "сао",
		},
		{
			name:     "se",
			input:    "se",
			expected: "сэ",
		},
		{
			name:     "sei",
			input:    "sei",
			expected: "сэй",
		},
		{
			name:     "sen",
			input:    "sen",
			expected: "сэнь",
		},
		{
			name:     "seng",
			input:    "seng",
			expected: "сэн",
		},
		{
			name:     "si",
			input:    "si",
			expected: "сы",
		},
		{
			name:     "song",
			input:    "song",
			expected: "сун",
		},
		{
			name:     "sou",
			input:    "sou",
			expected: "соу",
		},
		{
			name:     "su",
			input:    "su",
			expected: "су",
		},
		{
			name:     "suan",
			input:    "suan",
			expected: "суань",
		},
		{
			name:     "sui",
			input:    "sui",
			expected: "суй",
		},
		{
			name:     "sun",
			input:    "sun",
			expected: "сунь",
		},
		{
			name:     "suo",
			input:    "suo",
			expected: "со",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := tt.input
			got := ConvertToPallady(input)
			if got != tt.expected {
				t.Errorf("got %s, expected %s", got, tt.expected)
			}

			// convert first letter to uppercase
			ttRunes := []rune(tt.input)
			ttRunes[0] = unicode.ToUpper(ttRunes[0])
			upInput := string(ttRunes)

			// convert first letter to uppercase
			upRunes := []rune(tt.expected)
			upRunes[0] = unicode.ToUpper(upRunes[0])
			upExp := string(upRunes)

			got = ConvertToPallady(upInput)

			if got != upExp {
				t.Errorf("got %s, expected %s", got, upExp)
			}
		})
	}
}

func TestConvertToPalladySH(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "sha",
			input:    "sha",
			expected: "ша",
		},
		{
			name:     "shai",
			input:    "shai",
			expected: "шай",
		},
		{
			name:     "shan",
			input:    "shan",
			expected: "шань",
		},
		{
			name:     "shang",
			input:    "shang",
			expected: "шан",
		},
		{
			name:     "shao",
			input:    "shao",
			expected: "шао",
		},
		{
			name:     "she",
			input:    "she",
			expected: "шэ",
		},
		{
			name:     "shei",
			input:    "shei",
			expected: "шэй",
		},
		{
			name:     "shen",
			input:    "shen",
			expected: "шэнь",
		},
		{
			name:     "sheng",
			input:    "sheng",
			expected: "шэн",
		},
		{
			name:     "shi",
			input:    "shi",
			expected: "ши",
		},
		{
			name:     "shou",
			input:    "shou",
			expected: "шоу",
		},
		{
			name:     "shu",
			input:    "shu",
			expected: "шу",
		},
		{
			name:     "shua",
			input:    "shua",
			expected: "шуа",
		},
		{
			name:     "shuai",
			input:    "shuai",
			expected: "шуай",
		},
		{
			name:     "shuan",
			input:    "shuan",
			expected: "шуань",
		},
		{
			name:     "shuang",
			input:    "shuang",
			expected: "шуан",
		},
		{
			name:     "shui",
			input:    "shui",
			expected: "шуй",
		},
		{
			name:     "shun",
			input:    "shun",
			expected: "шунь",
		},
		{
			name:     "shuo",
			input:    "shuo",
			expected: "шо",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := tt.input
			got := ConvertToPallady(input)
			if got != tt.expected {
				t.Errorf("got %s, expected %s", got, tt.expected)
			}

			// convert first letter to uppercase
			ttRunes := []rune(tt.input)
			ttRunes[0] = unicode.ToUpper(ttRunes[0])
			upInput := string(ttRunes)

			// convert first letter to uppercase
			upRunes := []rune(tt.expected)
			upRunes[0] = unicode.ToUpper(upRunes[0])
			upExp := string(upRunes)

			got = ConvertToPallady(upInput)

			if got != upExp {
				t.Errorf("got %s, expected %s", got, upExp)
			}
		})
	}
}

func TestConvertToPalladyT(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "ta",
			input:    "ta",
			expected: "та",
		},
		{
			name:     "tai",
			input:    "tai",
			expected: "тай",
		},
		{
			name:     "tan",
			input:    "tan",
			expected: "тань",
		},
		{
			name:     "tang",
			input:    "tang",
			expected: "тан",
		},
		{
			name:     "tao",
			input:    "tao",
			expected: "тао",
		},
		{
			name:     "te",
			input:    "te",
			expected: "тэ",
		},
		{
			name:     "tei",
			input:    "tei",
			expected: "тэй",
		},
		{
			name:     "ten",
			input:    "ten",
			expected: "тэнь",
		},
		{
			name:     "teng",
			input:    "teng",
			expected: "тэн",
		},
		{
			name:     "ti",
			input:    "ti",
			expected: "ти",
		},
		{
			name:     "tian",
			input:    "tian",
			expected: "тянь",
		},
		{
			name:     "tiang",
			input:    "tiang",
			expected: "тян",
		},
		{
			name:     "tiao",
			input:    "tiao",
			expected: "тяо",
		},
		{
			name:     "tie",
			input:    "tie",
			expected: "те",
		},
		{
			name:     "ting",
			input:    "ting",
			expected: "тин",
		},
		{
			name:     "tong",
			input:    "tong",
			expected: "тун",
		},
		{
			name:     "tou",
			input:    "tou",
			expected: "тоу",
		},
		{
			name:     "tu",
			input:    "tu",
			expected: "ту",
		},
		{
			name:     "tuan",
			input:    "tuan",
			expected: "туань",
		},
		{
			name:     "tui",
			input:    "tui",
			expected: "туй",
		},
		{
			name:     "tun",
			input:    "tun",
			expected: "тунь",
		},
		{
			name:     "tuo",
			input:    "tuo",
			expected: "то",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := tt.input
			got := ConvertToPallady(input)
			if got != tt.expected {
				t.Errorf("got %s, expected %s", got, tt.expected)
			}

			// convert first letter to uppercase
			ttRunes := []rune(tt.input)
			ttRunes[0] = unicode.ToUpper(ttRunes[0])
			upInput := string(ttRunes)

			// convert first letter to uppercase
			upRunes := []rune(tt.expected)
			upRunes[0] = unicode.ToUpper(upRunes[0])
			upExp := string(upRunes)

			got = ConvertToPallady(upInput)

			if got != upExp {
				t.Errorf("got %s, expected %s", got, upExp)
			}
		})
	}
}

func TestConvertToPalladyW(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "wa",
			input:    "wa",
			expected: "ва",
		},
		{
			name:     "wai",
			input:    "wai",
			expected: "вай",
		},
		{
			name:     "wan",
			input:    "wan",
			expected: "вань",
		},
		{
			name:     "wang",
			input:    "wang",
			expected: "ван",
		},
		{
			name:     "wao",
			input:    "wao",
			expected: "вао",
		},
		{
			name:     "wei",
			input:    "wei",
			expected: "вэй",
		},
		{
			name:     "wen",
			input:    "wen",
			expected: "вэнь",
		},
		{
			name:     "weng",
			input:    "weng",
			expected: "вэн",
		},
		{
			name:     "wo",
			input:    "wo",
			expected: "во",
		},
		{
			name:     "wu",
			input:    "wu",
			expected: "у",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := tt.input
			got := ConvertToPallady(input)
			if got != tt.expected {
				t.Errorf("got %s, expected %s", got, tt.expected)
			}

			// convert first letter to uppercase
			ttRunes := []rune(tt.input)
			ttRunes[0] = unicode.ToUpper(ttRunes[0])
			upInput := string(ttRunes)

			// convert first letter to uppercase
			upRunes := []rune(tt.expected)
			upRunes[0] = unicode.ToUpper(upRunes[0])
			upExp := string(upRunes)

			got = ConvertToPallady(upInput)

			if got != upExp {
				t.Errorf("got %s, expected %s", got, upExp)
			}
		})
	}
}

func TestConvertToPalladyX(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "xi",
			input:    "xi",
			expected: "си",
		},
		{
			name:     "xia",
			input:    "xia",
			expected: "ся",
		},
		{
			name:     "xian",
			input:    "xian",
			expected: "сянь",
		},
		{
			name:     "xiang",
			input:    "xiang",
			expected: "сян",
		},
		{
			name:     "xiao",
			input:    "xiao",
			expected: "сяо",
		},
		{
			name:     "xie",
			input:    "xie",
			expected: "се",
		},
		{
			name:     "xin",
			input:    "xin",
			expected: "синь",
		},
		{
			name:     "xing",
			input:    "xing",
			expected: "син",
		},
		{
			name:     "xiong",
			input:    "xiong",
			expected: "сюн",
		},
		{
			name:     "xiu",
			input:    "xiu",
			expected: "сю",
		},
		{
			name:     "xu",
			input:    "xu",
			expected: "сюй",
		},
		{
			name:     "xuan",
			input:    "xuan",
			expected: "сюань",
		},
		{
			name:     "xue",
			input:    "xue",
			expected: "сюэ",
		},
		{
			name:     "xun",
			input:    "xun",
			expected: "сюнь",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := tt.input
			got := ConvertToPallady(input)
			if got != tt.expected {
				t.Errorf("got %s, expected %s", got, tt.expected)
			}

			// convert first letter to uppercase
			ttRunes := []rune(tt.input)
			ttRunes[0] = unicode.ToUpper(ttRunes[0])
			upInput := string(ttRunes)

			// convert first letter to uppercase
			upRunes := []rune(tt.expected)
			upRunes[0] = unicode.ToUpper(upRunes[0])
			upExp := string(upRunes)

			got = ConvertToPallady(upInput)

			if got != upExp {
				t.Errorf("got %s, expected %s", got, upExp)
			}
		})
	}
}

func TestConvertToPalladyY(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "ya",
			input:    "ya",
			expected: "я",
		},
		{
			name:     "yai",
			input:    "yai",
			expected: "яй",
		},
		{
			name:     "yan",
			input:    "yan",
			expected: "янь",
		},
		{
			name:     "yang",
			input:    "yang",
			expected: "ян",
		},
		{
			name:     "yao",
			input:    "yao",
			expected: "яо",
		},
		{
			name:     "ye",
			input:    "ye",
			expected: "е",
		},
		{
			name:     "yi",
			input:    "yi",
			expected: "и",
		},
		{
			name:     "yin",
			input:    "yin",
			expected: "инь",
		},
		{
			name:     "ying",
			input:    "ying",
			expected: "ин",
		},
		{
			name:     "yo",
			input:    "yo",
			expected: "йо",
		},
		{
			name:     "yong",
			input:    "yong",
			expected: "юн",
		},
		{
			name:     "you",
			input:    "you",
			expected: "ю",
		},
		{
			name:     "yu",
			input:    "yu",
			expected: "юй",
		},
		{
			name:     "yuan",
			input:    "yuan",
			expected: "юань",
		},
		{
			name:     "yue",
			input:    "yue",
			expected: "юэ",
		},
		{
			name:     "yun",
			input:    "yun",
			expected: "юнь",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := tt.input
			got := ConvertToPallady(input)
			if got != tt.expected {
				t.Errorf("got %s, expected %s", got, tt.expected)
			}

			// convert first letter to uppercase
			ttRunes := []rune(tt.input)
			ttRunes[0] = unicode.ToUpper(ttRunes[0])
			upInput := string(ttRunes)

			// convert first letter to uppercase
			upRunes := []rune(tt.expected)
			upRunes[0] = unicode.ToUpper(upRunes[0])
			upExp := string(upRunes)

			got = ConvertToPallady(upInput)

			if got != upExp {
				t.Errorf("got %s, expected %s", got, upExp)
			}
		})
	}
}

func TestConvertToPalladyZ(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "za",
			input:    "za",
			expected: "цза",
		},
		{
			name:     "zai",
			input:    "zai",
			expected: "цзай",
		},
		{
			name:     "zan",
			input:    "zan",
			expected: "цзань",
		},
		{
			name:     "zang",
			input:    "zang",
			expected: "цзан",
		},
		{
			name:     "zao",
			input:    "zao",
			expected: "цзао",
		},
		{
			name:     "ze",
			input:    "ze",
			expected: "цзэ",
		},
		{
			name:     "zei",
			input:    "zei",
			expected: "цзэй",
		},
		{
			name:     "zem",
			input:    "zem",
			expected: "цзэм",
		},
		{
			name:     "zen",
			input:    "zen",
			expected: "цзэнь",
		},
		{
			name:     "zeng",
			input:    "zeng",
			expected: "цзэн",
		},
		{
			name:     "zi",
			input:    "zi",
			expected: "цзы",
		},
		{
			name:     "zong",
			input:    "zong",
			expected: "цзун",
		},
		{
			name:     "zou",
			input:    "zou",
			expected: "цзоу",
		},
		{
			name:     "zu",
			input:    "zu",
			expected: "цзу",
		},
		{
			name:     "zuan",
			input:    "zuan",
			expected: "цзуань",
		},
		{
			name:     "zui",
			input:    "zui",
			expected: "цзуй",
		},
		{
			name:     "zun",
			input:    "zun",
			expected: "цзунь",
		},
		{
			name:     "zuo",
			input:    "zuo",
			expected: "цзо",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := tt.input
			got := ConvertToPallady(input)
			if got != tt.expected {
				t.Errorf("got %s, expected %s", got, tt.expected)
			}

			// convert first letter to uppercase
			ttRunes := []rune(tt.input)
			ttRunes[0] = unicode.ToUpper(ttRunes[0])
			upInput := string(ttRunes)

			// convert first letter to uppercase
			upRunes := []rune(tt.expected)
			upRunes[0] = unicode.ToUpper(upRunes[0])
			upExp := string(upRunes)

			got = ConvertToPallady(upInput)

			if got != upExp {
				t.Errorf("got %s, expected %s", got, upExp)
			}
		})
	}
}

func TestConvertToPalladyZH(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "zha",
			input:    "zha",
			expected: "чжа",
		},
		{
			name:     "zhai",
			input:    "zhai",
			expected: "чжай",
		},
		{
			name:     "zhan",
			input:    "zhan",
			expected: "чжань",
		},
		{
			name:     "zhang",
			input:    "zhang",
			expected: "чжан",
		},
		{
			name:     "zhao",
			input:    "zhao",
			expected: "чжао",
		},
		{
			name:     "zhe",
			input:    "zhe",
			expected: "чжэ",
		},
		{
			name:     "zhei",
			input:    "zhei",
			expected: "чжэй",
		},
		{
			name:     "zhen",
			input:    "zhen",
			expected: "чжэнь",
		},
		{
			name:     "zheng",
			input:    "zheng",
			expected: "чжэн",
		},
		{
			name:     "zhi",
			input:    "zhi",
			expected: "чжи",
		},
		{
			name:     "zhong",
			input:    "zhong",
			expected: "чжун",
		},
		{
			name:     "zhou",
			input:    "zhou",
			expected: "чжоу",
		},
		{
			name:     "zhu",
			input:    "zhu",
			expected: "чжу",
		},
		{
			name:     "zhua",
			input:    "zhua",
			expected: "чжуа",
		},
		{
			name:     "zhuai",
			input:    "zhuai",
			expected: "чжуай",
		},
		{
			name:     "zhuan",
			input:    "zhuan",
			expected: "чжуань",
		},
		{
			name:     "zhuang",
			input:    "zhuang",
			expected: "чжуан",
		},
		{
			name:     "zhui",
			input:    "zhui",
			expected: "чжуй",
		},
		{
			name:     "zhun",
			input:    "zhun",
			expected: "чжунь",
		},
		{
			name:     "zhuo",
			input:    "zhuo",
			expected: "чжо",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := tt.input
			got := ConvertToPallady(input)
			if got != tt.expected {
				t.Errorf("got %s, expected %s", got, tt.expected)
			}

			// convert first letter to uppercase
			ttRunes := []rune(tt.input)
			ttRunes[0] = unicode.ToUpper(ttRunes[0])
			upInput := string(ttRunes)

			// convert first letter to uppercase
			upRunes := []rune(tt.expected)
			upRunes[0] = unicode.ToUpper(upRunes[0])
			upExp := string(upRunes)

			got = ConvertToPallady(upInput)

			if got != upExp {
				t.Errorf("got %s, expected %s", got, upExp)
			}
		})
	}
}
