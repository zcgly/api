package models

type ApiResponse struct {
	Success bool   `json:"success"`
	Msg     string `json:"msg,omitempty"`
	Data    any    `json:"data"`
	Code    int    `json:"code"`
	Title   string `json:"title,omitempty"`
}

func NewMsgResponse(msg string) *ApiResponse {
	return &ApiResponse{Success: false, Msg: msg, Code: 500, Title: "出错啦！"}
}

type TitledError struct {
	title   string
	content string
}

func (t TitledError) Error() string {
	return t.content
}

func (t TitledError) Title() string {
	return t.title
}

func NewTitledError(title, content string) *TitledError {
	return &TitledError{title: title, content: content}
}
