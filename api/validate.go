package api

import (
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/gofiber/fiber/v2"
	"regexp"
)

func ValidateSearchBy(c *fiber.Ctx) error {
	type BySearchRequestValues struct {
		Search string
		Page   string
		Year   string
		Return string
		Type   string
	}

	b := BySearchRequestValues{
		Search: c.Query("s"),
		Page:   c.Query("page"),
		Year:   c.Query("y"),
		Return: c.Query("r"),
		Type:   c.Query("type"),
	}

	err := validation.ValidateStruct(&b,
		validation.Field(&b.Search, validation.Required),
		validation.Field(&b.Page, validation.Match(regexp.MustCompile("(0*(?:[1-9][0-9]?|100))"))), //1-100
		validation.Field(&b.Year, validation.Match(regexp.MustCompile("^(?:19|20)\\d{2}$"))),       // 1900-2099
		validation.Field(&b.Return, validation.In("json", "xml")),
		validation.Field(&b.Type, validation.In("movie", "series", "episode")))

	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
