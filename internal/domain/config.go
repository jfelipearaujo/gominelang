package domain

type Config struct {
	Version   string               `yaml:"version" validate:"required"`
	Engine    *Engine              `yaml:"engine" validate:"required"`
	Translate map[string]Translate `yaml:"translate" validate:"required,dive"`
}

type Engine struct {
	GoogleTranslate *GoogleTranslate `yaml:"google_translate"`
	OpenAI          *OpenAI          `yaml:"open_ai"`
}

type GoogleTranslate struct {
	Enabled bool `yaml:"enabled"`
}

type OpenAI struct {
	Enabled bool   `yaml:"enabled"`
	APIKey  string `yaml:"api_key" validate:"required_if=Enabled true"`
}

type Translate struct {
	From          string         `yaml:"from" validate:"required"`
	To            string         `yaml:"to" validate:"required"`
	Lang          *Lang          `yaml:"lang"`
	PatchouliBook *PatchouliBook `yaml:"patchouli_books"`
}

type Lang struct {
	Input  string `yaml:"input" validate:"required"`
	Output string `yaml:"output" validate:"required"`
}

type PatchouliBook struct {
	Input  string `yaml:"input" validate:"required"`
	Output string `yaml:"output" validate:"required"`
}
