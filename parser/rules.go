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
	rules["https://meduza.io/"] = Rule{
		Title: "main h1",
		Exists: "meta[content='article']",
		Description: "main .GeneralMaterial-article p:nth-of-type(1)",
	}
	return rules
}