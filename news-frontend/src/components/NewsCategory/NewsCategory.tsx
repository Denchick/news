import React from "react";
import { Jumbotron, Row, Col } from "react-bootstrap";

interface INewsCategoryProps {
    title?: string;
    children?: React.ReactNode;
}

const NewsCategory = ({title, children}: INewsCategoryProps) => {
    return (
        <Jumbotron>
            <Row>
                <Col>
                    <h2 className="text-center">{title}</h2>
                    <hr className="my-4" />
                </Col>
            </Row>
            <Row>
                {children}
            </Row>
        </Jumbotron>
    );
}

export default NewsCategory;