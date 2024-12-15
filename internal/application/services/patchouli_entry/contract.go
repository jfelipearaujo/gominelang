package patchouli_entry

// Service responsible to translate patchouli entries
//
// Example:
//   - inputFolder: ./my/mod/patchouli_books/book/en_us/entries
//   - outputFolder: ./my/mod/patchouli_books/book/pt_br/entries
type Service interface {
	// Define the language to translate from and to
	//
	// Example:
	//   - fromLang: en_us
	//   - toLang: pt_br
	SetLang(fromLang string, toLang string)

	// Receives the input folder and the output folder and translates all the entries files
	// inside the input folder to the output folder
	//
	// Example:
	//   - inputFolder: ./my/mod/patchouli_books/book/en_us/entries
	//	 - outputFolder: ./my/mod/patchouli_books/book/pt_br/entries
	Translate(inputFolder string, outputFolder string) error
}
