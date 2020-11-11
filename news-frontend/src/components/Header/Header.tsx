import React from "react";

interface IHeaderProps {
    emodji?: string;
}

const Header = ({emodji}: IHeaderProps) => (
    <header>
        <h1 className="display-1 text-center">
            {emodji || "📰"} news
        </h1>
    </header>
)

export default Header;