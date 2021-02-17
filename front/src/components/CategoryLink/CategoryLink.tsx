import React from "react";
import {  Button } from "react-bootstrap";
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

    return disabled
        ? button
        : <Link to={to as string} style={{ textDecoration: "none" }}>{button}</Link>
}

export default CategoryLink;