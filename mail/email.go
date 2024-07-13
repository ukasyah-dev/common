package mail

import (
	"github.com/caitlinelfring/go-env-default"
	"github.com/matcornic/hermes/v2"
	"github.com/samber/lo"
	"github.com/ukasyah-dev/common/log"
)

type Email struct {
	From    string
	To      string
	Subject string
	Body    Body
}

type Body struct {
	Name    string
	Intros  []string
	Actions []Action
	Outros  []string
}

type Action struct {
	Color     string
	Link      string
	Text      string
	TextColor string
}

var h = hermes.Hermes{
	Theme: new(hermes.Default),
	Product: hermes.Product{
		Name: env.GetDefault("EMAIL_PRODUCT_NAME", "My App"),
		Link: env.GetDefault("EMAIL_PRODUCT_LINK", "https://example.com"),
		Logo: env.GetDefault("EMAIL_PRODUCT_LOGO", "https://example.com/logo.svg"),
	},
}

func (e *Body) GetHTML() string {
	email := hermes.Email{
		Body: hermes.Body{
			Name:   e.Name,
			Intros: e.Intros,
			Actions: lo.Map(e.Actions, func(a Action, i int) hermes.Action {
				return hermes.Action{
					Button: hermes.Button{
						Color:     a.Color,
						Link:      a.Link,
						Text:      a.Text,
						TextColor: a.TextColor,
					},
				}
			}),
			Outros: e.Outros,
		},
	}

	result, err := h.GenerateHTML(email)
	if err != nil {
		log.Errorf("Failed to generate email html: %s", err)
		return ""
	}

	return result
}
