import React from "react";
import { Container, Jumbotron, Row } from "react-bootstrap";
import CategoryButton from "../../components/CategoryButton/CategoryButton";
import Footer from "../../components/Footer/Footer";
import Header from "../../components/Header/Header";

const MainPage = () => (
    <>
        <Jumbotron>
            <Container>
                <Header />
                <p className="lead text-center">No more annoying notifications</p>
                <Row className="mb-4">
                    <CategoryButton title="🌎 News" />
                    <CategoryButton title="👨‍💻 Technology" />
                    <CategoryButton title="🙂 People's blogs" />
                </Row>
                <Row className="mb-4">
                    <CategoryButton title="🇩🇪 Deutsche Sprache" disabled />
                    <CategoryButton title="🇵🇱 Język polski" disabled />
                </Row>
                <Row className="mb-4">
                    <CategoryButton title="🌈 Front-End" disabled />
                    <CategoryButton title="🤡 Machine Learning" disabled />
                    <CategoryButton title="⚙️ Back-End" disabled />
                </Row>
            </Container>
        </Jumbotron>
        <Footer />
    </>
);

export default MainPage;