# news

Хитро собирает новости из разных источников в одном месте.

![preview](preview.png)

## Как работает

Модуль состоит из трех пакетов:

- `database` - подключение и запросы к БД
- `server` - отображение новостей в красивом виде
- `parser` - собирает новости, парсит и сохраняет в базу

Сейчас поддерживаются 2 источника новостей:

- RSS-ленты разных форматов, парсинг происходит с помощью [gofeed](github.com/mmcdole/gofeed).
- Парсинг сайтов через указывание правил а-ля запросы jQuery (работает с [goquery](https://github.com/PuerkitoBio/goquery), вдохновлено [Telegram Instant View](https://instantview.telegram.org/)).

## Запуск локально

```
go get github.com/denchick/news
# запуск сервера
cd $GOPATH/src/github.com/denchick/news/server
go build
go run main.go
# запуск парсера
cd $GOPATH/src/github.com/denchick/news/parser
go build
go run main.go
```



