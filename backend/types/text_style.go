package types

import (
	"fmt"
)

// TextStyle 定义文本优化的风格
type TextStyle string

const (
	TextStyleStandard      TextStyle = "standard"
	TextStyleAcademic      TextStyle = "academic"
	TextStyleSimple        TextStyle = "simple"
	TextStyleFlowing       TextStyle = "flowing"
	TextStyleFormal        TextStyle = "formal"
	TextStyleInformal      TextStyle = "informal"
	TextStyleExpand        TextStyle = "expand"
	TextStyleShorten       TextStyle = "shorten"
	TextStyleShakespearean TextStyle = "shakespearean"
	TextStyleOldEnglish    TextStyle = "old_english"
)

// validStyles 存储所有有效的文本风格
var validStyles = map[TextStyle]bool{
	TextStyleStandard:      true,
	TextStyleAcademic:      true,
	TextStyleSimple:        true,
	TextStyleFlowing:       true,
	TextStyleFormal:        true,
	TextStyleInformal:      true,
	TextStyleExpand:        true,
	TextStyleShorten:       true,
	TextStyleShakespearean: true,
	TextStyleOldEnglish:    true,
}

// promptTemplates 存储所有文本风格的提示模板
var promptTemplates = map[TextStyle]string{
	TextStyleStandard: `[TextStyleStandard]:
You are a professional text optimization assistant. Please rewrite the text in a standard and clear way, improving readability while maintaining the original meaning.

Text to Rewrite:
[%s]

Return Format:[Your text here]`,
	TextStyleAcademic: `[TextStyleAcademic]:
You are an academic writing expert. Please rewrite the text using academic language, professional terminology, and formal expressions suitable for academic papers and research reports.

Text to Rewrite:
[%s]

Return Format:[Your text here]`,
	TextStyleSimple: `[TextStyleSimple]:
You are a text simplification expert. Please rewrite the text using simple, easy-to-understand language, avoiding complex vocabulary and sentence structures to make the content more accessible.

Text to Rewrite:
[%s]

Return Format:[Your text here]`,
	TextStyleFlowing: `[TextStyleFlowing]:
You are a literary writing expert. Please rewrite the text using fluid and elegant language, focusing on rhythm and coherence to make the text more expressive.

Text to Rewrite:
[%s]

Return Format:[Your text here]`,
	TextStyleFormal: `[TextStyleFormal]:
You are a business writing expert. Please rewrite the text using formal and professional language suitable for business settings and official documents.

Text to Rewrite:
[%s]

Return Format:[Your text here]`,
	TextStyleInformal: `[TextStyleInformal]:
You are a casual communication expert. Please rewrite the text using relaxed and natural language with everyday expressions to make the content more friendly and approachable.

Text to Rewrite:
[%s]

Return Format:[Your text here]`,
	TextStyleExpand: `[TextStyleExpand]:
You are a content expansion expert. Please expand the original text by adding more details and explanations to make the content richer and more comprehensive.

Text to Rewrite:
[%s]

Return Format:[Your text here]`,
	TextStyleShorten: `[TextStyleShorten]:
You are a text condensation expert. Please maintain the core information while reducing the text length and removing redundant content.

Text to Rewrite:
[%s]

Return Format:[Your text here]`,
	TextStyleShakespearean: `[TextStyleShakespearean]:
You are a Shakespearean text transformation expert. Please rewrite the text in the style of Shakespeare with the following requirements:
1. Use dramatic and theatrical language reminiscent of Shakespearean plays
2. Add exaggerated expressions and metaphors for dramatic effect
3. Incorporate comedic elements and witty wordplay
4. Use archaic English terms and sentence structures where appropriate
5. Transform mundane descriptions into grandiose theatrical scenes

Text to Rewrite:
[%s]

Return Format:[Your text here]`,
	TextStyleOldEnglish: `[TextStyleOldEnglish]:
You are an Old English language expert. Please transform the text following these rules:
1. Handle special Old English characters:
   - æ (can be written as .ae)
   - þ (can be written as .th)
   - ð (can be written as .dh)
2. Consider inflected forms and different case endings
3. Support character variations and normalize them
4. Handle dialect variations
5. Maintain the historical context and meaning

Special Instructions:
- If input contains special characters (æ, þ, ð), process them directly
- If input uses dot notation (.ae, .th, .dh), convert them to proper characters
- Consider all possible inflected forms of words
- Normalize character variations to standard forms
- Provide modern English equivalents while preserving the original meaning

Text to Process:
[%s]

Return Format:[Your text here]`,
}

// GetPrompt 获取对应风格的提示
func (s TextStyle) GetPrompt(input string, sourceLanguage Language, targetLanguage Language) string {
	// 对于古英语风格，强制使用古英语作为源语言
	if s == TextStyleOldEnglish {
		sourceLanguage = LanguageOldEnglish
	}

	// 通用规则
	commonRule := fmt.Sprintf("[System Prompt]: You are a professional text optimization assistant. "+
		"Please transform the text from %s to %s while maintaining the specified style. "+
		"If you cannot complete the style transformation as requested, return the original text without any modifications. "+
		"Only perform style transformation and language conversion of the text. "+
		"Do not perform any other actions like summarization or content modification. "+
		"Please provide the transformed text directly without including any introductory phrases or concluding remarks. ",
		GetLanguageName(sourceLanguage), GetLanguageName(targetLanguage))

	if template, ok := promptTemplates[s]; ok {
		return commonRule + "\n\n" + fmt.Sprintf(template, input)
	}
	return commonRule + "\n\n" + fmt.Sprintf(promptTemplates[TextStyleStandard], input)
}

// ValidateStyle 验证文本风格是否有效
func ValidateStyle(style TextStyle) bool {
	return validStyles[style]
}
