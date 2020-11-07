import React from 'react';
import Container from 'react-bootstrap/esm/Container';
import Header from './components/Header/Header';
import Footer from './components/Footer/Footer'
import TechnologyCategory from './components/NewsCategory/TechnologyCategory';

const App = () => {
    return (
        <Container>
            <Header />
            <TechnologyCategory />
            <Footer />
        </Container>
    );
}

export default App;
