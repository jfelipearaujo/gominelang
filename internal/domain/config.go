package domain

type Config struct {
	Version   string               `yaml:"version" validate:"required"`
	Translate map[string]Translate `yaml:"translate" validate:"required,dive"`
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
