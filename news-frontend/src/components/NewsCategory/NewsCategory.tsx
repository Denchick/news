import React from "react";
import { Row, Col } from "react-bootstrap";

interface INewsCategoryProps {
    title?: string;
    children?: React.ReactNode;
}

const NewsCategory = ({title, children}: INewsCategoryProps) => {
    return (
        <div style={{marginBottom: 50}}>
            <Row>
                <Col>
                    <hr className="my-4" />
                    <h2 className="text-center">{title}</h2>
                    <hr className="my-4" />
                </Col>
            </Row>
            <Row>
                {children}
            </Row>
        </div>
    );
}

export default NewsCategory;