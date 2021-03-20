package fake

const (
	NoMockPresent    = "no fake request present for the url"
	IncorrectRequest = "request does not match the expectation received: %v, expected: %v"
	IncorrectHeader  = "header does not match the expectation received: %v, expected: %v"
)
