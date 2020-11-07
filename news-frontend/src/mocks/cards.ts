import { INewsCard } from "../data/NewsCard";
import vcIcon from "../data/logo/vc-ru.png"
import telegramIcon from "../data/logo/telegram.png"
import { mockArticleVc, mockArticleTelegram } from "./articles";

const newsCard: INewsCard[] = [
    {
        source: {
            title: "vc.ru",
            icon: vcIcon
        },
        articles: [
            mockArticleVc,
            mockArticleVc,
            mockArticleVc,
        ]
    },
    {
        source: {
            title: "Подкаст Радио Платформа",
            icon: telegramIcon
        },
        articles: [
            mockArticleTelegram,
            mockArticleTelegram,
            mockArticleTelegram
        ]
    }
]

export default newsCard;