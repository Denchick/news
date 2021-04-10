import { IArticle } from "../data/Article";

export const mockArticleVc: IArticle = {
    title: "Отель на Мальдивах предложил гостям неограниченное проживание на весь 2021 год — услуга стоит $30 тысяч",
    description: "За эти деньги гости получат бунгало, бесплатные завтраки и трансфер от аэропорта.",
    publishedAt: new Date(2020, 11, 7, 21, 0, 0),
    createdAt: new Date(2020, 11, 8, 21, 0, 0),
    url: "https://vc.ru/offline/174608-otel-na-maldivah-predlozhil-gostyam-neogranichennoe-prozhivanie-na-ves-2021-god-usluga-stoit-30-tysyach"
}

export const mockArticleTelegram: IArticle = {
    title: "#08 - FAANG Interview (Sergii Sema, Google)",
    description: "В этом выпуске у нас в гостях Сергей Сема (Google) и мы попробуем разобраться как подготовиться и пройти техническое интервью в компании уровня FAANG.",
    createdAt: new Date(2020, 8, 21, 11, 41, 0),
    url: "https://www.youtube.com/watch?v=Jvuk3ZjYIZw"
}
