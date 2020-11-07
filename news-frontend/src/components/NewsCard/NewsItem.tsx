import React from "react";
import { ListGroup } from "react-bootstrap"

interface INewsItemProps {
    title: string;
    description: string;
    howLong: string;
    source: string
}

const NewsItem = ({title, description, howLong, source}: INewsItemProps) => (
    <ListGroup.Item>
        <h4>{title}</h4>
        <p className="mb-1">{description}</p>
        <small>{howLong} @ {source}</small>
    </ListGroup.Item>
)

export default NewsItem;