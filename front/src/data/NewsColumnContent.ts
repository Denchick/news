import { IArticle } from "./Article";

export interface ISourceMeta {
    title: string
    url: string
    icon: string
}
export interface INewsColumnContent {
    source: ISourceMeta;
    articles: IArticle[],
}