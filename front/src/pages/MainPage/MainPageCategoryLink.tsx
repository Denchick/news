import React from "react";
import { Button } from "react-bootstrap";
import { Link } from "react-router-dom";
import { IMainPageCategory } from "./MainPageCategories";

interface IMainPageCategoryLinkProps {
    category: IMainPageCategory;
}

const MainPageCategoryLink = ({ category }: IMainPageCategoryLinkProps) => {
    const { title, path } = category;
    const disabled = !path;
    const button = (
        <Button variant="secondary" size="lg" disabled={disabled} block>
            {title}
        </Button>
    );

    return disabled
        ? button
        : <Link to={path as string} style={{ textDecoration: "none" }}>{button}</Link>
}

export default MainPageCategoryLink;