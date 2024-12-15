package translate_tag

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
}
