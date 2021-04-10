import React from 'react';
import { BrowserRouter, Route, Switch } from 'react-router-dom';
import MainPage from './pages/MainPage/MainPage';
import NewsPage from './pages/NewsPage/NewsPage';

const App = () => {
    return (
        <BrowserRouter>
            <Switch>
                <Route path="/" exact={true}>
                    <MainPage />
                </Route>
                <Route path="/dev" exact={true}>
                    <NewsPage />
                </Route>
            </Switch>
      </BrowserRouter>
    );
}

export default App;
