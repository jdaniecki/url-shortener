package shortener

type Shortener struct {
	id int
}

func NewShortener() *Shortener {
	return &Shortener{id: 0}
}

func (s *Shortener) Shorten(url string) string {
	shortUrl := toBase62(s.id)
	s.id++
	return shortUrl
}

func toBase62(num int) string {
	const base62Chars = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	if num == 0 {
		return string(base62Chars[0])
	}

	result := ""
	for num > 0 {
		remainder := num % 62
		result = string(base62Chars[remainder]) + result
		num = num / 62
	}

	return result
}
