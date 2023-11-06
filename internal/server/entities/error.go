package entities

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Result struct {
	Error  *Error      `json:"error,omitempty"`
	Result interface{} `json:"result,omitempty"`
}
