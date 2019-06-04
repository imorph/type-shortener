package shrt

// Shortener can ether Shorten ether Resolve
type Shortener interface {
	Shorten(url string) string
	Resolve(url string) string
	}