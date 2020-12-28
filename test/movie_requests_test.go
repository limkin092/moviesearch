package test

import (
	"moviesearch/repository"
	"testing"
)

func TestImportMovies(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"JSON_NDeserializtedWithStruct_Return1",
			args{s: `{
				"Search": [{
				"Title": "Tesla",
				"Year": "2020",
				"imdbID": "tt5259822",
				"Type": "movie",
				"Poster": "https://m.media-amazon.com/images/M/MV5BYzg0MjQ0ODUtYTgyNC00Y2Y5LWE5NDctODY3ZTFkYmZkNGFiXkEyXkFqcGdeQXVyMTE1MzI2NzIz._V1_SX300.jpg"
			}, {
				"Title": "The Secret Life of Nikola Tesla",
				"Year": "1980",
				"imdbID": "tt0079985",
				"Type": "movie",
				"Poster": "https://m.media-amazon.com/images/M/MV5BMjE3NjU4Nzk1Nl5BMl5BanBnXkFtZTcwNjk0MzcxMQ@@._V1_SX300.jpg"
			}],
				"totalResults": "65",
				"Response": "True"
			}`},
			1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := repository.ImportMovies(tt.args.s); got != tt.want {
				t.Errorf("ImportMovies() = %v, want %v", got, tt.want)
			}
		})
	}
}
