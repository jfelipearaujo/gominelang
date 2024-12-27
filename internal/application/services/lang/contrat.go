package lang

// Service responsible to merge lang files and translate each one
//
// Example:
//   - input: en_us.json
//   - output: pt_br.json
type Service interface {
	// Define the language to translate from and to
	//
	// Example:
	//   - fromLang: en_us
	//   - toLang: pt_br
	SetLang(fromLang string, toLang string)

	// Translate the lang files
	//
	// Example:
	//   - inputFolder: ./my/mod/lang
	//   - outputFolder: ./my/mod/lang
	Translate(inputFolder string, outputFolder string) error
}
