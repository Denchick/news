import React from "react";
import { OverlayTrigger, Tooltip } from "react-bootstrap";

interface INewsItemProps {
    title: string;
    description: string;
    howLong: string;
    source: string
}

const NewsItem = ({title, description, howLong, source}: INewsItemProps) => (
    <OverlayTrigger
        placement="right"
        overlay={
            <Tooltip id="tooltip">
                <div className="text-left">
                    <p className="mb-1">{description}</p>
                    <small>{howLong} @ {source}</small>
                </div>
            </Tooltip>
        }
    >
        <p>{title}</p>
    </OverlayTrigger>
)

export default NewsItem;