package response

// respBody is the struct holding response body.
type respBody struct {
	Status int `json:"status"`
	Data   any `json:"data"`
}
