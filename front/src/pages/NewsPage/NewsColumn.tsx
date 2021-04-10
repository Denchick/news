import React from "react";
import { INewsColumnContent } from "../../data/NewsColumnContent";
import NewsIcon from "./NewsIcon";
import NewsItem from "./NewsItem";

interface INewsColumnProps {
    columnContent: INewsColumnContent;
}

const NewsColumn = ({ columnContent }: INewsColumnProps) => {
    const { source, articles } = columnContent;
    const { title, icon, url } = source;
    return (
        <div>
            <div className="d-flex w-100 justify-content-start">
                <NewsIcon src={icon} />
                <h2 className="h4 ml-2">{title}</h2>
            </div>
            <div>
                {articles.map(article =>
                    <NewsItem article={article} source={url} />
                )}
            </div>
        </div>
    );
}

export default NewsColumn;