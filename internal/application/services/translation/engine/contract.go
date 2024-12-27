package engine

// Service responsible to translate text
type Service interface {
	// Translate the text from the 'from' language to the 'to' language
	//
	// Example:
	//   - from: en_us
	//   - to: pt_br
	//   - text: Hello world!
	Translate(from string, to string, text string) (string, error)
}
