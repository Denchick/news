package main

type Rule struct {
	Exists      string
	Title       string
	Description string
}

func getRules() map[string]Rule {
	rules := make(map[string]Rule)
	rules["https://www.aviasales.ru/blog/"] = Rule{
		Title: ".article__title",
		Exists: "meta[content='article']",
		Description: "article p:nth-of-type(3)",
	}
	rules["https://vandrouki.ru/"] = Rule{
		Title: ".entry-title",
		Exists: "body[class*='single-post']",
		Description: ".entry-content",
	}
	return rules
}