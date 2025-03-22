package request

type AddProductReq struct {
	FilmInformation      filmInformation `json:"film_information" binding:"required"`
	FilmChanges          filmChange
	OtherFilmInformation otherFilmInformation `json:"other_film_informations" binding:"required"`
}

type GetFilmByIdReq struct {
	FilmId int `json:"film_id" binding:"required"`
}
