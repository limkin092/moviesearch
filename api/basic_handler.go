package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/parnurzeal/gorequest"
	"moviesearch/config"
	"moviesearch/repository"
)

func HelloWorld(c *fiber.Ctx) error {
	return c.SendString("Hello World!")
}

func BySearch(c *fiber.Ctx) error {

	if err := ValidateSearchBy(c); err != nil {
		return err
	}

	if m := repository.CheckCacheForMovie(c); len(m) != 0 {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"Result": m})
	}

	if body, err := requestBySearch(c); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error": "Encountered error while requesting data"})
	} else {
		repository.ImportMovies(body)
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"Result": body})
	}

}

func requestBySearch(c *fiber.Ctx) (string, []error) {
	request := gorequest.New()
	_, body, errors := request.Get("https://movie-database-imdb-alternative.p.rapidapi.com/").
		Set("x-rapidapi-key", config.Config("X-RAPIDAPI-KEY")).
		Set("x-rapidapi-host", "movie-database-imdb-alternative.p.rapidapi.com").
		Set("userQueryString", "true").
		Param("s", c.Query("s")).
		Param("page", c.Query("page")).
		Param("y", c.Query("y")).
		Param("r", c.Query("return")).
		Param("type", c.Query("type")).
		End()

	return body, errors
}
