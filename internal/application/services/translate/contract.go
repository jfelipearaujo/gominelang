package translate

// Service responsible to translate text
type Service interface {
	// Translate the text from the 'from' language to the 'to' language
	//
	// Example:
	//   - from: en_us
	//   - to: pt_br
	//   - text: Hello world!
	Translate(from string, to string, text string) (string, error)

	// Fix wrong translations given the original and translated text against the tags
	// that are not translated correctly
	//
	// The tag is a string that starts with '$(l:' and ends with ')', for example:
	// $(l:minecraft:item.pickaxe.diamond)
	//
	// If in the translation process the tag is translated incorrectly, the original
	// tag will be replaced with the translated tag
	//
	// Example:
	//   - original: This is a $(l:minecraft:item.pickaxe.diamond) item
	//   - translated: Este é um $(l:minecraft:item.picatera.diamond) item
	//   - output: This is a $(l:minecraft:item.pickaxe.diamond) item
	FixWrongTranslation(original string, translated string) string
}
