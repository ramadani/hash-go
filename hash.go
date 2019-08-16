package hash

// Hash hashing string
type Hash interface {
	Make(s string) (string, error)
	Check(s, hash string) bool
}
