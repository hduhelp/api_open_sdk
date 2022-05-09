package faqv1

import (
	"fmt"
	"strings"
)

type AnswerRichtexts []*AnswerRichtext

func (x AnswerRichtexts) ToMarkdown() string {
	contentArray := make([]string, 0)
	for _, content := range x {
		switch content.Type {
		case "text":
			if content.Content == "" {
				continue
			}
			contentArray = append(contentArray, content.Content)
		case "hyperlink":
			contentArray = append(contentArray, fmt.Sprintf("[%s](%s)", content.Text, content.Url))
		case "img":
			contentArray = append(contentArray, fmt.Sprintf("![%s](%s)", content.Key, content.Src))
		}
	}
	return strings.Join(contentArray, ",")
}
