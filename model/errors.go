package model

//BadRequest swagger:response BadRequest
type BadRequest struct {
	// Error message
	// in: string
	Message string
	// Error status
	// in: int
	StatusCode int
}
