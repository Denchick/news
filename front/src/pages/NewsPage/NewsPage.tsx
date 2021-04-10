import React from "react";
import { Container, Col, Row, Jumbotron } from "react-bootstrap";
import Footer from "../../components/Footer/Footer";
import Header from "../../components/Header/Header";
import NewsColumn from "./NewsColumn";
import CategoryName from "../../data/CategoryName";
import { INewsColumnContent } from "../../data/NewsColumnContent";

interface INewsPageProps {
    title: CategoryName,
    columns: INewsColumnContent[][]
}

const NewsPage = ({ title, columns }: INewsPageProps) => {
    return (
        <Container>
            <Header title={title} />
            {columns.map(row => {
                <Jumbotron className="p-3" fluid>
                    <Row>
                        {row.map(column => {
                            <Col>
                                <NewsColumn columnContent={column} />
                            </Col>
                        })}
                    </Row>
                </Jumbotron>
            })}
            <Footer />
        </Container>
    );

}

export default NewsPage;