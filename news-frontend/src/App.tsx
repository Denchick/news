import React from 'react';
import { BrowserRouter, Route, Switch } from 'react-router-dom';
import MainPage from './pages/MainPage/MainPage';
import TechnologyPage from './pages/TechnologyPage/TechnologyPage';

const App = () => {
    return (
        <BrowserRouter>
            <Switch>
                <Route path="/" exact={true}>
                    <MainPage />
                </Route>
                <Route path="/tech" exact={true}>
                    <TechnologyPage />
                </Route>
            </Switch>
      </BrowserRouter>
    );
}

export default App;
