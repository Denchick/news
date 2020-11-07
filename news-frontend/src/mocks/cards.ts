import { INewsCard } from "../data/NewsCard";
import vcruLogo from "../data/logo/vc-ru.png"
import mockArticle from "./article";

const newsCard: INewsCard = {
    source: {
        title: "vc.ru",
        icon: vcruLogo
    },
    articles: [
        mockArticle,
        mockArticle,
        mockArticle,
    ]
}

export default newsCard;