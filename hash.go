package hash

// Hash hashing string
type Hash interface {
	Make(string) string
	Check(string) error
}
