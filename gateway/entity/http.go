package entity

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	FullName string `json:"full_name"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type MovieSearchRequest struct {
	Search string `json:"search"`
	ImdbID string `json:"imdb_id"`
	Page   int64  `json:"page"`
}
