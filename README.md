# type-shortener

This is homework #1*

## What to do

Написать тип, который реализует интерфейс:

```go
type Shortener interface {
	Shorten(url string) string
	Resolve(url string) string
}
```

Метод Shorten - возвращать "короткую" ссылку (выбор алгоритма - за студентом), например otus.ru/some-long-link -> otus.ru/jhg34 и сохранять соответствие короткой и исходной ссылок в памяти (не используя БД, а использовать, например, map).
При вызове метода Resolve - отдавать "длинную ссылку" или пустую строку, если ссылка не найдена.

