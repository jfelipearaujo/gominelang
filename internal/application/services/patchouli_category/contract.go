package patchouli_category

// Service responsible to translate patchouli categories
//
// Example:
//   - inputFolder: ./my/mod/patchouli_books/book/en_us/categories
//   - outputFolder: ./my/mod/patchouli_books/book/pt_br/categories
type Service interface {
	// Define the language to translate from and to
	//
	// Example:
	//   - fromLang: en_us
	//   - toLang: pt_br
	SetLang(fromLang string, toLang string)

	// Receives the input folder and the output folder and translates all the categories files
	// inside the input folder to the output folder
	//
	// Example:
	//   - inputFolder: ./my/mod/patchouli_books/book/en_us/categories
	//	 - outputFolder: ./my/mod/patchouli_books/book/pt_br/categories
	Translate(inputFolder string, outputFolder string) error
}
