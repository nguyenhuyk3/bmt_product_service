package request

type filmChange struct {
	ChangedBy string
}

type otherFilmInformation struct {
	PosterUrl  string `json:"poster_url" binding:"required"`
	TrailerUrl string `json:"trailer_url" binding:"required"`
}

type filmInformation struct {
	Title       string   `json:"title" binding:"required"`
	Description string   `json:"description" binding:"required"`
	ReleaseDate string   `json:"release_date" binding:"required"`
	Genres      []string `json:"genres" binding:"required"`
	// This prop will have format as hh:mm:ss
	// When using api, we will use 2h39m
	// When stroring at databse then this will be at 02:39:00
	Duration string `json:"duration" binding:"required"`
}
