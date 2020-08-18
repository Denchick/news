package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"server/articles"
)

func StartParse(db *gorm.DB) {
	for url, rule := range getRules() {
		articlesArr := parseSite(url, rule)
		articles.CreateOrUpdate(db, articlesArr)
	}

	for _, url := range getFeeds() {
		articlesArr := parseFeed(url)
		articles.CreateOrUpdate(db, articlesArr)
	}
}

func main() {
	articlesList := []articles.Article{
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
	dsn := "host=rogue.db.elephantsql.com port=5432 user=acumlegw dbname=acumlegw password=2ZbAnECs3Lfdc9k9sbpRWzcCOHd1mo97"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.Migrator().DropTable(&articles.Article{})
	db.AutoMigrate(&articles.Article{})


	articles.CreateOrUpdate(db, articlesList)



}