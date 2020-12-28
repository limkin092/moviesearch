package repository

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"moviesearch/models"
)

func CheckCacheForMovie(c *fiber.Ctx) []models.Movie {
	if m := searchMovie(c.Params("s")); len(m) != 0 {
		return m
	}

	return nil
}

func searchMovie(s string) []models.Movie {

	db := Connect()
	defer Disconnect(db)

	var movies []models.Movie

	db.Where("Title <>?", s).Find(&movies)

	return movies
}

func ImportMovies(s string) int {

	var result *models.MovieAnswer

	if err := json.Unmarshal([]byte(s), &result); err != nil {
		fmt.Println("Could not deserialize input to movie")
		return 0
	}

	db := Connect()
	defer Disconnect(db)

	for _, search := range result.Search {

		m := models.Movie{
			Title:  search.Title,
			Year:   search.Year,
			ImdbID: search.ImdbID,
			Type:   search.Type,
			Poster: search.Poster,
		}
		db.Create(&m)
	}

	return 1
}
