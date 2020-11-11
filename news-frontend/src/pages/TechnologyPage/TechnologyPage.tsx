import React from "react";
import { Container, Col, Row } from "react-bootstrap";
import Footer from "../../components/Footer/Footer";
import Header from "../../components/Header/Header";
import NewsCard from "../../components/NewsCard/NewsCard";
import CategoryName from "../../data/CategoryName";
import newsCard from "../../mocks/cards";


const TechnologyPage = () => {
    return (
        <Container>
            <Header title={CategoryName.Technology} />
            <Row>
                <Col>
                    <NewsCard source={newsCard[0].source} articles={newsCard[0].articles} />            
                </Col>
                <Col>
                    <NewsCard source={newsCard[1].source} articles={newsCard[1].articles} />            
                </Col>
            </Row>
            <Footer />
        </Container>
    );

}

export default TechnologyPage;