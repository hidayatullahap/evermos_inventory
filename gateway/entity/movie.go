package entity

type MovieList struct {
	TotalResults int64   `json:"total_result"`
	Page         int64   `json:"page"`
	Result       int64   `json:"result"`
	Movies       []Movie `json:"movies"`
}

type Movie struct {
	Title  string `json:"title"`
	Year   string `json:"year"`
	ImdbID string `json:"imdb_id"`
	Type   string `json:"type"`
	Poster string `json:"poster"`
}

type MovieDetail struct {
	Title      string   `json:"title"`
	Year       string   `json:"year"`
	Rated      string   `json:"rated"`
	Released   string   `json:"released"`
	Runtime    string   `json:"runtime"`
	Genre      string   `json:"genre"`
	Director   string   `json:"director"`
	Writer     string   `json:"writer"`
	Actors     string   `json:"actors"`
	Plot       string   `json:"plot"`
	Language   string   `json:"language"`
	Country    string   `json:"country"`
	Awards     string   `json:"awards"`
	Poster     string   `json:"poster"`
	Ratings    []Rating `json:"ratings"`
	Metascore  string   `json:"metascore"`
	ImdbRating string   `json:"imdb_rating"`
	ImdbVotes  string   `json:"imdb_votes"`
	ImdbID     string   `json:"imdb_id"`
	Type       string   `json:"type"`
	DVD        string   `json:"DVD"`
	BoxOffice  string   `json:"box_office"`
	Production string   `json:"production"`
	Website    string   `json:"website"`
}

type Rating struct {
	Source string `json:"source"`
	Value  string `json:"value"`
}
