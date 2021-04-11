import React from "react";
import { Button } from "react-bootstrap";
import { Link } from "react-router-dom";
import { CategoryName, categoryPaths } from "../../data/Category";

interface IMainPageCategoryLinkProps {
    category: CategoryName;
}

const MainPageCategoryLink = ({ category }: IMainPageCategoryLinkProps) => {
    const path =  categoryPaths[category];
    const disabled = !path;
    const button = (
        <Button variant="secondary" size="lg" disabled={disabled} block>
            {category}
        </Button>
    );

    return disabled
        ? button
        : <Link to={path as string} style={{ textDecoration: "none" }}>{button}</Link>
}

export default MainPageCategoryLink;