package domain

type Content struct {
	ContentType string
	Text        string
}

func NewContent(contentType string, text string) Content {
	return Content{
		ContentType: contentType,
		Text:        text,
	}
}
