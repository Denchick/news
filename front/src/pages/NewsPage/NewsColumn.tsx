import React from "react";
import { INewsColumn } from "../../data/News";
import NewsItem from "./NewsItem";

interface INewsColumnProps {
    columnContent: INewsColumn;
}

const NewsColumn = ({ columnContent }: INewsColumnProps) => {
    const {feedName, feedUrl, articles} = columnContent;
    return (
        <div>
            <h2 className="h4">{feedName}</h2>
            <div>
                {articles.map(article =>
                    <NewsItem article={article} source={feedUrl} />
                )}
            </div>
        </div>
    );
}

export default NewsColumn;