import React from "react";

const Header = () => (
    <header>
        <h1 className="display-2" style={{ display: "inline" }}>
            news
        </h1>
        <a className="text-decoration-none lead" href="/search">
            🔎 пошуршать
        </a>
    </header>
)

export default Header;