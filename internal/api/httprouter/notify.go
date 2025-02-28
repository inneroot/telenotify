package httpRouter

type NotifyPostRequestBody struct {
	Message string `json:"message"`
	Sign    string `json:"sign"`
}
