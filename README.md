# news

Pet-projects just for fun. Collects news from different sources in one place.

![preview](preview.png)

## Whats next?

- add news by categories
- replace gorm to github.com/jackc/pgx
- add feed source management (replace `getFeed()`)
- rewrite frontend to TS/React
- deploy to cloud

## How to run?

```
docker-compose up
```

# API News

## `GET /news/technology`

response:

```
[
    {
        "sourceName": "vas3k.ru",
        "items": [
            {
                // newsItem1
            },
            {
                // newsItem2
            },
            ...
        ]
    },
    {
        ...
    }
]
```json