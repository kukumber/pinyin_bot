package converter

import (
	"testing"
	"unicode"
	"unicode/utf8"
)

var palladySyllableTests = []struct {
	input    string
	expected string
}{
	{"a", "а"}, {"ai", "ай"}, {"an", "ань"}, {"ang", "ан"}, {"ao", "ао"},
	{"ba", "ба"}, {"bai", "бай"}, {"ban", "бань"}, {"bang", "бан"}, {"bao", "бао"},
	{"bei", "бэй"}, {"ben", "бэнь"}, {"beng", "бэн"}, {"bi", "би"}, {"bian", "бянь"},
	{"biao", "бяо"}, {"bie", "бе"}, {"bin", "бинь"}, {"bing", "бин"}, {"bo", "бо"}, {"bu", "бу"},
	{"ca", "ца"}, {"cai", "цай"}, {"can", "цань"}, {"cang", "цан"}, {"cao", "цао"},
	{"ce", "цэ"}, {"cen", "цэнь"}, {"ceng", "цэн"}, {"ci", "цы"}, {"cong", "цун"},
	{"cou", "цоу"}, {"cu", "цу"}, {"cuan", "цуань"}, {"cui", "цуй"}, {"cun", "цунь"}, {"cuo", "цо"},
	{"cha", "ча"}, {"chai", "чай"}, {"chan", "чань"}, {"chang", "чан"}, {"chao", "чао"},
	{"che", "чэ"}, {"chen", "чэнь"}, {"cheng", "чэн"}, {"chi", "чи"}, {"chong", "чун"},
	{"chou", "чоу"}, {"chu", "чу"}, {"chua", "чуа"}, {"chuai", "чуай"}, {"chuan", "чуань"},
	{"chuang", "чуан"}, {"chui", "чуй"}, {"chun", "чунь"}, {"chuo", "чо"},
	{"da", "да"}, {"dai", "дай"}, {"dan", "дань"}, {"dang", "дан"}, {"dao", "дао"},
	{"de", "дэ"}, {"dei", "дэй"}, {"den", "дэнь"}, {"di", "ди"}, {"dia", "дя"},
	{"dian", "дянь"}, {"diang", "дян"}, {"diao", "дяо"}, {"die", "де"}, {"ding", "дин"},
	{"diu", "дю"}, {"dong", "дун"}, {"dou", "доу"}, {"du", "ду"}, {"duan", "дуань"},
	{"dui", "дуй"}, {"dun", "дунь"}, {"duo", "до"},
	{"e", "э"}, {"ei", "эй"}, {"en", "энь"}, {"eng", "эн"}, {"er", "эр"},
	{"fa", "фа"}, {"fan", "фань"}, {"fang", "фан"}, {"fei", "фэй"}, {"fen", "фэнь"},
	{"feng", "фэн"}, {"fiao", "фяо"}, {"fo", "фо"}, {"fou", "фоу"}, {"fu", "фу"},
	{"ga", "га"}, {"gai", "гай"}, {"gan", "гань"}, {"gang", "ган"}, {"gao", "гао"},
	{"ge", "гэ"}, {"gei", "гэй"}, {"gen", "гэнь"}, {"geng", "гэн"}, {"go", "го"},
	{"gong", "гун"}, {"gou", "гоу"}, {"gu", "гу"}, {"gua", "гуа"}, {"guai", "гуай"},
	{"guan", "гуань"}, {"guang", "гуан"}, {"gui", "гуй"}, {"gun", "гунь"}, {"guo", "го"},
	{"ha", "ха"}, {"hai", "хай"}, {"han", "хань"}, {"hang", "хан"}, {"hao", "хао"},
	{"he", "хэ"}, {"hei", "хэй"}, {"hen", "хэнь"}, {"heng", "хэн"}, {"hm", "хм"},
	{"hng", "хн"}, {"hong", "хун"}, {"hou", "хоу"}, {"hu", "ху"}, {"hua", "хуа"},
	{"huai", "хуай"}, {"huan", "хуань"}, {"huang", "хуан"}, {"hui", "хуэй"}, {"hun", "хунь"}, {"huo", "хо"},
	{"ji", "цзи"}, {"jia", "цзя"}, {"jian", "цзянь"}, {"jiang", "цзян"}, {"jiao", "цзяо"},
	{"jie", "цзе"}, {"jin", "цзинь"}, {"jing", "цзин"}, {"jiong", "цзюн"}, {"jiu", "цзю"},
	{"ju", "цзюй"}, {"juan", "цзюань"}, {"jue", "цзюэ"}, {"jun", "цзюнь"},
	{"ka", "ка"}, {"kai", "кай"}, {"kan", "кань"}, {"kang", "кан"}, {"kao", "као"},
	{"ke", "кэ"}, {"kei", "кэй"}, {"ken", "кэнь"}, {"keng", "кэн"}, {"kong", "кун"},
	{"kou", "коу"}, {"ku", "ку"}, {"kua", "куа"}, {"kuai", "куай"}, {"kuan", "куань"},
	{"kuang", "куан"}, {"kui", "куй"}, {"kun", "кунь"}, {"kuo", "ко"},
	{"la", "ла"}, {"lai", "лай"}, {"lan", "лань"}, {"lang", "лан"}, {"lao", "лао"},
	{"le", "лэ"}, {"lei", "лэй"}, {"leng", "лэн"}, {"li", "ли"}, {"lia", "ля"},
	{"lian", "лянь"}, {"liang", "лян"}, {"liao", "ляо"}, {"lie", "ле"}, {"lin", "линь"},
	{"ling", "лин"}, {"liu", "лю"}, {"lo", "ло"}, {"long", "лун"}, {"lou", "лоу"},
	{"lu", "лу"}, {"lü", "люй"}, {"lv", "люй"}, {"luan", "луань"}, {"lüan", "люань"},
	{"lvan", "люань"}, {"lüe", "люэ"}, {"lve", "люэ"}, {"lun", "лунь"}, {"lün", "люнь"},
	{"lvn", "люнь"}, {"luo", "ло"},
	{"m", "м"}, {"ma", "ма"}, {"mai", "май"}, {"man", "мань"}, {"mang", "ман"},
	{"mao", "мао"}, {"me", "мэ"}, {"mei", "мэй"}, {"men", "мэнь"}, {"meng", "мэн"},
	{"mi", "ми"}, {"mian", "мянь"}, {"miao", "мяо"}, {"mie", "ме"}, {"min", "минь"},
	{"ming", "мин"}, {"miu", "мю"}, {"mm", "мм"}, {"mo", "мо"}, {"mou", "моу"}, {"mu", "му"},
	{"n", "нь"}, {"na", "на"}, {"nai", "най"}, {"nan", "нань"}, {"nang", "нан"},
	{"nao", "нао"}, {"ne", "нэ"}, {"nei", "нэй"}, {"nen", "нэнь"}, {"neng", "нэн"},
	{"ng", "н"}, {"ni", "ни"}, {"nia", "ня"}, {"nian", "нянь"}, {"niang", "нян"},
	{"niao", "няо"}, {"nie", "не"}, {"nin", "нинь"}, {"ning", "нин"}, {"niu", "ню"},
	{"nong", "нун"}, {"nou", "ноу"}, {"nu", "ну"}, {"nun", "нунь"}, {"nü", "нюй"},
	{"nv", "нюй"}, {"nuan", "нуань"}, {"nüe", "нюэ"}, {"nve", "нюэ"}, {"nuo", "но"},
	{"o", "о"}, {"ou", "оу"},
	{"pa", "па"}, {"pai", "пай"}, {"pan", "пань"}, {"pang", "пан"}, {"pao", "пао"},
	{"pei", "пэй"}, {"pen", "пэнь"}, {"peng", "пэн"}, {"pi", "пи"}, {"pian", "пянь"},
	{"piang", "пян"}, {"piao", "пяо"}, {"pie", "пе"}, {"pin", "пинь"}, {"ping", "пин"},
	{"po", "по"}, {"pou", "поу"}, {"pu", "пу"},
	{"qi", "ци"}, {"qia", "ця"}, {"qian", "цянь"}, {"qiang", "цян"}, {"qiao", "цяо"},
	{"qie", "це"}, {"qin", "цинь"}, {"qing", "цин"}, {"qiong", "цюн"}, {"qiu", "цю"},
	{"qu", "цюй"}, {"quan", "цюань"}, {"que", "цюэ"}, {"qun", "цюнь"},
	{"ran", "жань"}, {"rang", "жан"}, {"rao", "жао"}, {"re", "жэ"}, {"rem", "жэм"},
	{"ren", "жэнь"}, {"reng", "жэн"}, {"ri", "жи"}, {"rong", "жун"}, {"rou", "жоу"},
	{"ru", "жу"}, {"rua", "жуа"}, {"ruan", "жуань"}, {"rui", "жуй"}, {"run", "жунь"}, {"ruo", "жо"},
	{"sa", "са"}, {"sai", "сай"}, {"san", "сань"}, {"sang", "сан"}, {"sao", "сао"},
	{"se", "сэ"}, {"sei", "сэй"}, {"sen", "сэнь"}, {"seng", "сэн"}, {"si", "сы"},
	{"song", "сун"}, {"sou", "соу"}, {"su", "су"}, {"suan", "суань"}, {"sui", "суй"},
	{"sun", "сунь"}, {"suo", "со"},
	{"sha", "ша"}, {"shai", "шай"}, {"shan", "шань"}, {"shang", "шан"}, {"shao", "шао"},
	{"she", "шэ"}, {"shei", "шэй"}, {"shen", "шэнь"}, {"sheng", "шэн"}, {"shi", "ши"},
	{"shou", "шоу"}, {"shu", "шу"}, {"shua", "шуа"}, {"shuai", "шуай"}, {"shuan", "шуань"},
	{"shuang", "шуан"}, {"shui", "шуй"}, {"shun", "шунь"}, {"shuo", "шо"},
	{"ta", "та"}, {"tai", "тай"}, {"tan", "тань"}, {"tang", "тан"}, {"tao", "тао"},
	{"te", "тэ"}, {"tei", "тэй"}, {"ten", "тэнь"}, {"teng", "тэн"}, {"ti", "ти"},
	{"tian", "тянь"}, {"tiang", "тян"}, {"tiao", "тяо"}, {"tie", "те"}, {"ting", "тин"},
	{"tong", "тун"}, {"tou", "тоу"}, {"tu", "ту"}, {"tuan", "туань"}, {"tui", "туй"},
	{"tun", "тунь"}, {"tuo", "то"},
	{"wa", "ва"}, {"wai", "вай"}, {"wan", "вань"}, {"wang", "ван"}, {"wao", "вао"},
	{"wei", "вэй"}, {"wen", "вэнь"}, {"weng", "вэн"}, {"wo", "во"}, {"wu", "у"},
	{"xi", "си"}, {"xia", "ся"}, {"xian", "сянь"}, {"xiang", "сян"}, {"xiao", "сяо"},
	{"xie", "се"}, {"xin", "синь"}, {"xing", "син"}, {"xiong", "сюн"}, {"xiu", "сю"},
	{"xu", "сюй"}, {"xuan", "сюань"}, {"xue", "сюэ"}, {"xun", "сюнь"},
	{"ya", "я"}, {"yai", "яй"}, {"yan", "янь"}, {"yang", "ян"}, {"yao", "яо"},
	{"ye", "е"}, {"yi", "и"}, {"yin", "инь"}, {"ying", "ин"}, {"yo", "йо"},
	{"yong", "юн"}, {"you", "ю"}, {"yu", "юй"}, {"yuan", "юань"}, {"yue", "юэ"}, {"yun", "юнь"},
	{"za", "цза"}, {"zai", "цзай"}, {"zan", "цзань"}, {"zang", "цзан"}, {"zao", "цзао"},
	{"ze", "цзэ"}, {"zei", "цзэй"}, {"zem", "цзэм"}, {"zen", "цзэнь"}, {"zeng", "цзэн"},
	{"zi", "цзы"}, {"zong", "цзун"}, {"zou", "цзоу"}, {"zu", "цзу"}, {"zuan", "цзуань"},
	{"zui", "цзуй"}, {"zun", "цзунь"}, {"zuo", "цзо"},
	{"zha", "чжа"}, {"zhai", "чжай"}, {"zhan", "чжань"}, {"zhang", "чжан"}, {"zhao", "чжао"},
	{"zhe", "чжэ"}, {"zhei", "чжэй"}, {"zhen", "чжэнь"}, {"zheng", "чжэн"}, {"zhi", "чжи"},
	{"zhong", "чжун"}, {"zhou", "чжоу"}, {"zhu", "чжу"}, {"zhua", "чжуа"}, {"zhuai", "чжуай"},
	{"zhuan", "чжуань"}, {"zhuang", "чжуан"}, {"zhui", "чжуй"}, {"zhun", "чжунь"}, {"zhuo", "чжо"},
}

func TestConvertToPallady(t *testing.T) {
	for _, tt := range palladySyllableTests {
		t.Run(tt.input, func(t *testing.T) {
			checkPallady(t, tt.input, tt.expected)
		})
	}
}

func TestConvertToPalladyText(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"empty input", "", ""},
		{"two syllables", "ni hao", "ни хао"},
		{"capitalized words", "Zhang San", "Чжан Сань"},
		{"tone marks", "nǐ hǎo", "ни хао"},
		{"non-pinyin text unchanged", "Hello, world!", "Hello, world!"},
		{"mixed text", "hello ni hao", "hello ни хао"},
		{"punctuation preserved", "hao, hao, hao!", "хао, хао, хао!"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConvertToPallady(tt.input); got != tt.expected {
				t.Errorf("ConvertToPallady(%s): got %s, expected %s", tt.input, got, tt.expected)
			}
		})
	}
}

// checkPallady checks the conversion of input and its capitalized form
func checkPallady(t *testing.T, input, expected string) {
	t.Helper()

	if got := ConvertToPallady(input); got != expected {
		t.Errorf("ConvertToPallady(%s): got %s, expected %s", input, got, expected)
	}

	upInput := upperFirst(input)
	upExpected := upperFirst(expected)

	if got := ConvertToPallady(upInput); got != upExpected {
		t.Errorf("ConvertToPallady(%s): got %s, expected %s", upInput, got, upExpected)
	}
}

func upperFirst(s string) string {
	r, size := utf8.DecodeRuneInString(s)
	return string(unicode.ToUpper(r)) + s[size:]
}

// TestSyllableValuesAreCyrillic guards against stray Latin letters
// sneaking into the Cyrillic replacement values
func TestSyllableValuesAreCyrillic(t *testing.T) {
	for k, v := range syllables {
		for _, r := range v {
			if unicode.IsLetter(r) && !unicode.Is(unicode.Cyrillic, r) {
				t.Errorf("syllables[%q] = %q contains non-Cyrillic letter %q", k, v, r)
			}
		}
	}
}
