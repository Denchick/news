import React from "react";
import { Col, Container, Jumbotron, Row } from "react-bootstrap";
import MainPageCategoryLink from "./MainPageCategoryLink";
import Footer from "../../components/Footer/Footer";
import Header from "../../components/Header/Header";
import { rows } from "./MainPageCategories";

const MainPage = () => (
    <>
        <Jumbotron>
            <Container>
                <Header title="📰 news" description="No more annoying notifications" /> {/* TODO get from config */}
                {rows.map(row => (
                    <Row className="mb-4">
                        {row.map(col => (
                            <Col>
                                <MainPageCategoryLink category={col} />
                            </Col>
                        ))}
                    </Row>
                ))}
            </Container>
        </Jumbotron>
        <Footer />
    </>
);

export default MainPage;