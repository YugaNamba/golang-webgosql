package controllers


type HttpStatusCode int

const (
	OK HttpStatusCode = 200
	Found HttpStatusCode = 302
	NotFound HttpStatusCode = 404
	Unauthorized HttpStatusCode = 401
	InternalServerError HttpStatusCode = 500
)

