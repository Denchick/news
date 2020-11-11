import React from "react";
import { Container, Jumbotron, Row } from "react-bootstrap";
import CategoryButton from "../../components/CategoryButton/CategoryButton";
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
                    <CategoryButton title={CategoryName.News} />
                    <CategoryButton title={CategoryName.Technology} />
                    <CategoryButton title={CategoryName.PeoplesBlogs} />
                </Row>
                <Row className="mb-4">
                    <CategoryButton title={CategoryName.German} disabled />
                    <CategoryButton title={CategoryName.Polish} disabled />
                </Row>
                <Row className="mb-4">
                    <CategoryButton title={CategoryName.Frontend} disabled />
                    <CategoryButton title={CategoryName.MachineLearning} disabled />
                    <CategoryButton title={CategoryName.Backend} disabled />
                </Row>
            </Container>
        </Jumbotron>
        <Footer />
    </>
);

export default MainPage;