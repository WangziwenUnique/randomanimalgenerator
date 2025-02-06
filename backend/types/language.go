package types

// Language 定义语言类型
type Language string

const (
	// 支持的语言
	LanguageEnglish    Language = "en"
	LanguageHaitian    Language = "ht"
	LanguageOldEnglish Language = "ang" // ISO 639-2/3 code for Old English
)

// validLanguages 存储所有有效的语言
var validLanguages = map[Language]bool{
	LanguageEnglish:    true,
	LanguageHaitian:    true,
	LanguageOldEnglish: true,
}

// ValidateLanguage 验证语言是否有效
func ValidateLanguage(lang Language) bool {
	return validLanguages[lang]
}

// GetLanguageName 获取语言的显示名称
func GetLanguageName(lang Language) string {
	switch lang {
	case LanguageEnglish:
		return "English"
	case LanguageHaitian:
		return "Haitian Creole"
	case LanguageOldEnglish:
		return "Old English"
	default:
		return string(lang)
	}
}
