import { IArticle } from "./Article";

export interface INewsSubcategory {
    subcategoryName: string;
    columns: INewsColumn[];
}

export interface INewsColumn {
    feedName: string,
    feedUrl: string,
    articles: IArticle[],
}
