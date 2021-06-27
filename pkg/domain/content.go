package domain

type Content struct {
	contentType string
	text        string
}

func NewContent(contentType string, text string) Content {
	return Content{
		contentType: contentType,
		text:        text,
	}
}
