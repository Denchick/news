import React from "react";
import { Col, Button } from "react-bootstrap";
import { Link } from "react-router-dom";
import CategoryName from "../../data/CategoryName";

interface ICategoryButtonProps {
    title: CategoryName;
    to?: string;
    disabled?: boolean;
}

const CategoryButton = ({title, to, disabled}: ICategoryButtonProps) => (
    <Col>
        <Button variant="secondary" size="lg" disabled={disabled} block>
            {title}
        </Button>
    </Col>
)

export default CategoryButton;