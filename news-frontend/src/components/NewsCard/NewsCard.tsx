import React from "react";
import { ListGroup } from "react-bootstrap";
import { IArticle } from "../../data/Article";
import { INewsSource } from "../../data/NewsSource";
import Icon from "../Icon/Icon";
import NewsItem from "./NewsItem";

interface INewsCardProps {
    source: INewsSource;
    articles: IArticle[],
}

const NewsCard = ({source, articles}: INewsCardProps) => {
    const {title, icon, url} = source;
    return (
        <div>
            <div className="d-flex w-100 justify-content-start">
                <Icon src={icon} />
                <h3 style={{marginLeft: 5}}>{title}</h3>
            </div>
            <ListGroup>
                {
                    articles.map(article => 
                        <NewsItem
                            title={article.title}
                            description={article.description}
                            howLong={"3 days ago"}
                            source={url}
                        />)
                }
            </ListGroup>
        </div>
    );
}

export default NewsCard;