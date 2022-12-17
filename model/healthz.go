package model

// A HealthzResponse expresses health check message.
type HealthzResponse struct {
	Message string `json:"message"`
}

// Write implements io.Writer
func (*HealthzResponse) Write(p []byte) (n int, err error) {
	panic("unimplemented")
}
