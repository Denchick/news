import React from "react";
import { OverlayTrigger, Tooltip } from "react-bootstrap";
import { IArticle } from "../../data/Article";

interface INewsItemProps {
    source: string;
    article: IArticle;
}

const NewsItem = ({ article, source }: INewsItemProps) => {
    const { url, title, description, publishedAt } = article;
    return (
        <OverlayTrigger
            placement="right"
            overlay={
                <Tooltip id="tooltip">
                    <div className="text-left">
                        <p className="mb-1">{description}</p>
                        <small>{publishedAt} @ {source}</small>
                    </div>
                </Tooltip>
            }
        >
            <a href={url}>
                <p>{title}</p>
            </a>
        </OverlayTrigger>
    );
}

export default NewsItem;