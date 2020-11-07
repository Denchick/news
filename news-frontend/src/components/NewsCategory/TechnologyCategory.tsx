import React from "react";
import { Row, Col } from "react-bootstrap";
import newsCard from "../../mocks/cards";
import NewsCard from "../NewsCard/NewsCard";
import NewsCategory from "./NewsCategory";

const TechnologyCategory = () => {
    return (
        <NewsCategory title="Технологии">
            <Col>
                <NewsCard source={newsCard.source} articles={newsCard.articles} />            
            </Col>
            <Col>
                <NewsCard source={newsCard.source} articles={newsCard.articles} />            
            </Col>
        </NewsCategory>
    )
}

export default TechnologyCategory