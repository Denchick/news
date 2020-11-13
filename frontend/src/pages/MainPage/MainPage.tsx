import React from "react";
import { Col, Container, Jumbotron, Row } from "react-bootstrap";
import CategoryLink from "../../components/CategoryLink/CategoryLink";
import Footer from "../../components/Footer/Footer";
import Header from "../../components/Header/Header";
import CategoryName from "../../data/CategoryName";
import Emoji from "../../data/Emoji";

const MainPage = () => (
    <>
        <Jumbotron>
            <Container>
                <Header title={Emoji.Newspaper + "news"} description="No more annoying notifications" />
                <Row className="mb-4">
                    <Col>
                        <CategoryLink title={CategoryName.News} to="/news" />
                    </Col>
                    <Col>
                        <CategoryLink title={CategoryName.Technology} to="/tech" />
                    </Col>
                    <Col>
                        <CategoryLink title={CategoryName.PeoplesBlogs} to="/blogs" />
                    </Col>
                </Row>
                <Row className="mb-4">
                    <Col>
                        <CategoryLink title={CategoryName.German} />
                    </Col>
                    <Col>
                        <CategoryLink title={CategoryName.Polish} />
                    </Col>
                </Row>
                <Row className="mb-4">
                    <Col>
                        <CategoryLink title={CategoryName.Frontend} />
                    </Col>
                    <Col>
                        <CategoryLink title={CategoryName.MachineLearning} />
                    </Col>
                    <Col>
                        <CategoryLink title={CategoryName.Backend} />
                    </Col>
                </Row>
            </Container>
        </Jumbotron>
        <Footer />
    </>
);

export default MainPage;