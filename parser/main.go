package main

import (
	"github.com/denchick/news/database"
)

func parseArticles() []database.Article { // TODO maybe should refactor
	articlesMap := make(map[string]database.Article)
	for url, rule := range getRules() {
		for _, article := range parseSite(url, rule) {
			if len(article.Link) == 0 {
				continue
			}
			articlesMap[article.Link] = article
		}
	}
	for _, url := range getFeeds() {
		for _, article := range parseFeed(url) {
			if len(article.Link) == 0 {
				continue
			}
			articlesMap[article.Link] = article
		}
	}

	result := make([]database.Article, len(articlesMap))
	i := 0
	for _, article := range articlesMap {
		result[i] = article
		i++
	}
	return result
}

func getMockArticles() []database.Article {
	return []database.Article{
		{
			Title:       "Путеводитель по Тюмени: примерный бюджет и программа",
			Description: "«Тюмень — столица деревень» — миф и бред. Это уже давно не так: город развивается быстро и может удивить даже тех, кого не удивляет раф с попкорном блёстками. Для активного отдыха, как писали выше, подходят биатлонный центр, Кулига-парк, в черте города — пруд «Лесной», эко-парк Затюменский, Гилёвская роща, созданы локации на набережной.",
			Link:        "https://www.aviasales.ru/blog/chto-delat-v-tumeni",
		},
		{
			Title: "Давай не в Стамбул: бюджет и маршрут на 5 дней и 4 города в Турции",
			Description: "Город, который, как мы помним, не является столицей Турции, принято или любить до беспамятства, невзирая на недостатки, или ругать до безумия, забывая об очевидных ништяках. Одни взахлёб, попивая чаёк из стаканчиков-тюльпанов, рассказывают о вкусной кухне, приемлемых ценах и базарах. Другие, дожёвывая рахат-лукум, бухтят о замусоренных улицах, навязчивости, орущих чайках. Обожать его модно, ругать его просто. Стамбул такой, Стамбул сякой... Хватит. Пора отправиться вглубь мятежной страны, где всю жизнь воюют, торгуют и молятся.",
			Link: "https://www.aviasales.ru/blog/davai-ne-v-stambul",
		},
		{
			Title: "Пил в маршрутке и спал с волками: шальная поездка в Киргизию",
			Description: "Я жуткий раздолбай, но мне везёт. Путешествия, которые не планирую, удаются гораздо круче любых заранее обдуманных. В этот раз со мной случилась Киргизия: палатка, волки, альпинисты и вкусная еда.",
			Link: "https://www.aviasales.ru/blog/shalnaya-kirgiziya",
		},
	}
}

func main() {
	db, err := database.Connect()
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&database.Article{})
	database.CreateOrUpdateArticles(db, parseArticles())
}