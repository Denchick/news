import { INewsCard } from "../data/NewsCard";
import vcIcon from "../data/icons/vc-ru.png"
import telegramIcon from "../data/icons/telegram.png"
import { mockArticleVc, mockArticleTelegram } from "./articles";

const newsCard: INewsCard[] = [
    {
        source: {
            title: "vc.ru",
            icon: vcIcon,
            url: "vc.ru"
        },
        articles: [
            mockArticleVc,
            mockArticleVc,
            mockArticleVc,
        ]
    },
    {
        source: {
            title: "platforma podcast",
            icon: telegramIcon,
            url: "t.me/radioplatforma"
        },
        articles: [
            mockArticleTelegram,
            mockArticleTelegram,
            mockArticleTelegram
        ]
    }
]

export default newsCard;