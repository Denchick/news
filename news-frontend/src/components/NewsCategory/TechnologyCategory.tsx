import React from "react";
import { Col } from "react-bootstrap";
import newsCard from "../../mocks/cards";
import NewsCard from "../NewsCard/NewsCard";
import NewsCategory from "./NewsCategory";

const TechnologyCategory = () => {
    return (
        <NewsCategory title="Технологии">
            <Col>
                <NewsCard source={newsCard[0].source} articles={newsCard[0].articles} />            
            </Col>
            <Col>
                <NewsCard source={newsCard[1].source} articles={newsCard[1].articles} />            
            </Col>
        </NewsCategory>
    )
}

export default TechnologyCategory