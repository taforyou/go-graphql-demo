// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package go_graphql_demo

type NewReview struct {
	VideoID     int64  `json:"videoId"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Rating      int    `json:"rating"`
	UserID      int64  `json:"userId"`
}

type NewScreenshot struct {
	VideoID     int64   `json:"videoId"`
	URL         string  `json:"url"`
	Description *string `json:"description"`
}

type NewVideo struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	UserID      int64  `json:"userId"`
	URL         string `json:"url"`
}