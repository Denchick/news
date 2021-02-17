import { IArticle } from "./Article";
import { INewsSource } from "./NewsSource";

export interface INewsCard {
    source: INewsSource;
    articles: IArticle[],
}