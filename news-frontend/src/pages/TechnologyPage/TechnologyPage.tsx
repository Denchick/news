import React from "react";
import { Container, Col, Row, Jumbotron } from "react-bootstrap";
import Footer from "../../components/Footer/Footer";
import Header from "../../components/Header/Header";
import NewsCard from "../../components/NewsCard/NewsCard";
import CategoryName from "../../data/CategoryName";
import newsCard from "../../mocks/cards";


const TechnologyPage = () => {
    return (
        <Container>
            <Header title={CategoryName.Technology} />
            <Jumbotron className="p-3" fluid>
                <Row>
                    <Col>
                        <NewsCard source={newsCard[0].source} articles={newsCard[0].articles} />            
                    </Col>
                    <Col>
                        <NewsCard source={newsCard[1].source} articles={newsCard[1].articles} />            
                    </Col>
                    <Col>
                        <NewsCard source={newsCard[1].source} articles={newsCard[1].articles} />            
                    </Col>
                </Row>
            </Jumbotron>

            <Jumbotron className="p-3" fluid>
                <Row className="mb-5">
                    <Col>
                        <NewsCard source={newsCard[0].source} articles={newsCard[0].articles} />            
                    </Col>
                    <Col>
                        <NewsCard source={newsCard[1].source} articles={newsCard[1].articles} />            
                    </Col>
                    <Col>
                        <NewsCard source={newsCard[1].source} articles={newsCard[1].articles} />            
                    </Col>
                </Row>
                <Row className="mb-5">
                    <Col>
                        <NewsCard source={newsCard[1].source} articles={newsCard[1].articles} />            
                    </Col>
                    <Col>
                        <NewsCard source={newsCard[0].source} articles={newsCard[0].articles} />            
                    </Col>
                    <Col>
                        <NewsCard source={newsCard[1].source} articles={newsCard[1].articles} />            
                    </Col>
                </Row>
            </Jumbotron>

            <Jumbotron className="p-3" fluid>
                <Row>
                    <Col>
                        <NewsCard source={newsCard[0].source} articles={newsCard[0].articles} />            
                    </Col>
                    <Col>
                        <NewsCard source={newsCard[1].source} articles={newsCard[1].articles} />            
                    </Col>
                    <Col>
                        <NewsCard source={newsCard[1].source} articles={newsCard[1].articles} />            
                    </Col>
                </Row>
            </Jumbotron>
            
            <Footer />
        </Container>
    );

}

export default TechnologyPage;