package domain

type PatchouliCategory struct {
	Name        string `json:"name" translate:"true"`
	Description string `json:"description" translate:"true"`
	SortNum     *int   `json:"sortnum,omitempty"`
	Icon        string `json:"icon"`
}

type PatchouliEntry struct {
	Name     string               `json:"name" translate:"true"`
	SortNum  int                  `json:"sortnum,omitempty"`
	Category string               `json:"category"`
	Icon     string               `json:"icon"`
	Pages    []PatchouliEntryPage `json:"pages"`
}

type PatchouliEntryPage struct {
	Type   string    `json:"type"`
	Title  *string   `json:"title,omitempty" translate:"true"`
	Images *[]string `json:"images,omitempty"`
	Border *bool     `json:"border,omitempty"`
	Text   *string   `json:"text,omitempty" translate:"true"`
	Recipe *string   `json:"recipe,omitempty"`
}

func (dst *PatchouliEntry) MapFrom(src *PatchouliEntry) {
	dst.Name = src.Name
	dst.SortNum = src.SortNum
	dst.Category = src.Category
	dst.Icon = src.Icon
	dst.Pages = make([]PatchouliEntryPage, len(src.Pages))
	for i := 0; i < len(src.Pages); i++ {
		dst.Pages[i].Type = src.Pages[i].Type

		// this is necessary because we are dealing with pointers
		if src.Pages[i].Title != nil {
			out := *src.Pages[i].Title
			dst.Pages[i].Title = &out
		}

		dst.Pages[i].Images = src.Pages[i].Images
		dst.Pages[i].Border = src.Pages[i].Border

		// this is necessary because we are dealing with pointers
		if src.Pages[i].Text != nil {
			out := *src.Pages[i].Text
			dst.Pages[i].Text = &out
		}

		dst.Pages[i].Recipe = src.Pages[i].Recipe
	}
}
