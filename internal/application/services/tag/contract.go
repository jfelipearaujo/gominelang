package tag

// Service responsible to handle the translation of tags
type Service interface {
	// Define the language to translate from and to
	//
	// Example:
	//   - fromLang: en_us
	//   - toLang: pt_br
	SetLang(fromLang string, toLang string)

	// Receives the input struct and translates all the fields that have the "translate" tag
	//
	// Example:
	//
	//		type myStruct struct {
	//		  Title string `json:"title" translate:"true"`
	//		  Text  string `json:"text"`
	//		}
	//
	//	Output: The value of the "title" field will be translated to the desired language
	HandleTranslation(input interface{}) error

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
