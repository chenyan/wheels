package protocol

type Response struct {
	Code  int            `json:"code"`
	Error string         `json:"error"`
	Data  map[string]any `json:"data,omitempty"`
}

type M map[string]any

func (m M) R() Response {
	return Response{Data: m}
}

func (m M) E(code int, err string) Response {
	return Response{Error: err, Code: code, Data: m}
}
