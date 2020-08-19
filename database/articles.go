package database

import (
	"gorm.io/gorm"
	"strings"
	"time"
)

type Article struct {
	Link        string `gorm:"column:id"`
	Title       string `gorm:"index"`
	Description string
	CreatedAt time.Time
}

func (a *Article) BeforeCreate(tx *gorm.DB) error {
	a.CreatedAt = time.Now()
	return nil
}

func ReadAllArticles(db *gorm.DB) []Article {
	var articles []Article
	db.Find(&articles)
	return articles
}

func CreateOrUpdateArticles(db *gorm.DB, articles []Article) {
	addedLinks := readLinks(db)
	added, notAdded := splitArticles(articles, addedLinks)
	db.Create(&notAdded)
	db.Save(&added)
}

func FindArticleByTitle(db *gorm.DB, query string) (result []Article) {
	articles := ReadAllArticles(db)
	for _, article := range articles {
		if strings.Contains(article.Title, query) {
			result = append(result, article)
		}
	}
	return
}

func readLinks(db *gorm.DB) []string {
	var links []string
	db.Select("id").Find(&links)
	return links
}

func splitArticles(articles []Article, addedLinks []string) (notAdded []Article, added []Article) {
	for _, article := range articles {
		if contains(addedLinks, article.Link) {
			added = append(added, article)
		} else {
			notAdded = append(notAdded, article)
		}
	}
	return
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

func GetMockArticles() []Article {
	return []Article{
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
