package types

type ResponseEnvelope[T any] struct {
	Data   T        `json:"data,omitempty"`
	Errors []string `json:"errors,omitempty"`
}

func NewResponseEnvelope[T any](data T) ResponseEnvelope[T] {
	return ResponseEnvelope[T]{Data: data}
}
