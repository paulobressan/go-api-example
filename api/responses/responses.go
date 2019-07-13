package responses

import (
	"encoding/json"
	"net/http"
)

type meta struct {
	Self string `json:"self"`
}

type response struct {
	Meta    meta        `json:"meta"`
	Records interface{} `json:"records"`
}

// Response : Criar resposta http
type Response struct {
	W http.ResponseWriter
	R *http.Request
}

func (r *Response) getResponsePattern(self string, values interface{}) response {
	return response{
		Meta: meta{
			Self: self,
		},
		Records: values,
	}
}

// Ok : Requisição feita com sucesso
func (r *Response) Ok(values interface{}) {
	r.W.WriteHeader(200)
	json.NewEncoder(r.W).Encode(r.getResponsePattern(r.R.RequestURI, values))
}

// BadRequest : Requisição mal formatada
func (r *Response) BadRequest(message string) {
	r.W.WriteHeader(404)
	json.NewEncoder(r.W).Encode(message)
}

// NotFound : Requisição não encontrada ou item não encontrado
func (r *Response) NotFound(message string) {
	r.W.WriteHeader(400)
	json.NewEncoder(r.W).Encode(message)
}
