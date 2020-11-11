import React from "react";
import { Col, Button } from "react-bootstrap";

interface ICategoryButtonProps {
    title: string;
    disabled?: boolean;
}

const CategoryButton = ({title, disabled}: ICategoryButtonProps) => (
        <Col>
            <Button variant="secondary" size="lg" disabled={disabled} block>
                {title}
            </Button>
        </Col>
)

export default CategoryButton;