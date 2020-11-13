import React from "react";
import { Col, Button } from "react-bootstrap";
import { Link } from "react-router-dom";
import CategoryName from "../../data/CategoryName";

interface ICategoryLinkProps {
    title: CategoryName;
    to?: string;
}

const CategoryLink = ({ title, to }: ICategoryLinkProps) => {
    const disabled = !to;
    const button = (
        <Button variant="secondary" size="lg" disabled={disabled} block>
            {title}
        </Button>
    );

    return (
        <Col>
            {disabled ? button : <Link to={to as string} style={{textDecoration: "none"}}>{button}</Link>}
        </Col>
    )
}

export default CategoryLink;