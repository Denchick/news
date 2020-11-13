import React from "react";
import { Container, Jumbotron, Row } from "react-bootstrap";
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
                    <CategoryLink title={CategoryName.News} to="/news" />
                    <CategoryLink title={CategoryName.Technology} to="/tech" />
                    <CategoryLink title={CategoryName.PeoplesBlogs} to="/blogs" />
                </Row>
                <Row className="mb-4">
                    <CategoryLink title={CategoryName.German} />
                    <CategoryLink title={CategoryName.Polish} />
                </Row>
                <Row className="mb-4">
                    <CategoryLink title={CategoryName.Frontend} />
                    <CategoryLink title={CategoryName.MachineLearning} />
                    <CategoryLink title={CategoryName.Backend} />
                </Row>
            </Container>
        </Jumbotron>
        <Footer />
    </>
);

export default MainPage;