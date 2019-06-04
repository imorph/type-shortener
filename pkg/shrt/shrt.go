package shrt

// Shortener can ether Shorten ether Resolve
type Shortener interface {
	Shorten(url string) string
	Resolve(url string) string
}

// URLShortener is type
type URLShortener struct {
//map is here?
}

// Shorten implements Shorten interface
func (u URLShortener) Shorten(url string) string{
	hashurl:="hash" 
	return hashurl
}

// Resolve implements Resolve interface
func (u URLShortener) Resolve(hashurl string) string{
	url:= "url"
	return url
}