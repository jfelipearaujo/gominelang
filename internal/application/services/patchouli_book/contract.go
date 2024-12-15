package patchouli_book

// Service responsible to translate patchouli books
//
// Example:
//   - inputFolder: ./my/mod/patchouli_books/book/en_us
//   - outputFolder: ./my/mod/patchouli_books/book/pt_br
type Service interface {
	// Define the language to translate from and to
	//
	// Example:
	//   - fromLang: en_us
	//   - toLang: pt_br
	SetLang(fromLang string, toLang string)

	// Receives the input folder and the output folder and translates all the files
	// inside the input folder to the output folder
	//
	// # Attention: The input folder must contains the 'categories' and 'entries' folders
	//
	// Example:
	//   - inputFolder: ./my/mod/patchouli_books/book/en_us
	//	 - outputFolder: ./my/mod/patchouli_books/book/pt_br
	Translate(inputFolder string, outputFolder string) error
}
