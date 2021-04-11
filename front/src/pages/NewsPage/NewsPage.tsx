import React, { useEffect, useState } from "react";
import { Container, Col, Row, Jumbotron } from "react-bootstrap";
import Footer from "../../components/Footer/Footer";
import Header from "../../components/Header/Header";
import NewsColumn from "./NewsColumn";
import { CategoryName, categoryPaths } from "../../data/Category";
import { INewsSubcategory } from "../../data/News";

interface INewsPageState {
    newsSubcategories: INewsSubcategory[],
    loading: boolean,
    error?: any;
}

interface INewsPageProps {
    categoryName: CategoryName,
}

const NewsPage = ({ categoryName }: INewsPageProps) => {
    const [newsPageState, setNewsPageState] = useState<INewsPageState>({
        newsSubcategories: [],
        loading: true,
    });

    useEffect(() => {
        fetch(`/api/v1/news?category=${categoryPaths[categoryName]}`)
            .then(response => response.json())
            .then(response => setNewsPageState({
                newsSubcategories: response,
                loading: false,
                error: false
            }))
            .catch(error => setNewsPageState({
                newsSubcategories: [],
                loading: false,
                error: error
            })
        );
    }, [categoryName]);

    const { newsSubcategories } = newsPageState;
    return (
        <Container>
            <Header title={categoryName} />
            {newsSubcategories.map(subcategory => (
                <>
                    <h2>{subcategory.subcategoryName}</h2>
                    <Jumbotron className="p-3" fluid>
                        <Row>
                            {subcategory.columns.map(column => (
                                <Col>
                                    <NewsColumn columnContent={column} />
                                </Col>
                            ))}
                        </Row>
                    </Jumbotron>
                </>
            ))}
            <Footer />
        </Container>
    );

}

export default NewsPage;