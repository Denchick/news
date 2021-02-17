import React from "react";
import CategoryName from "../../data/CategoryName";

interface IHeaderProps {
    title: string | CategoryName;
    description?: string;
}

const Header = ({title, description}: IHeaderProps) => (
    <header className="m-5">
        <h1 className="display-1 text-center">
            {title}
        </h1>
        {description && <p className="lead text-center">{description}</p>}
    </header>
)

export default Header;