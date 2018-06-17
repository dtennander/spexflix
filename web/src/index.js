import React from 'react';
import ReactDOM from 'react-dom';
import { BrowserRouter, Switch, Route } from 'react-router-dom';

import './index.css';
import registerServiceWorker from './registerServiceWorker';

import Home from "./pages/Home";
import Year from "./pages/Year";

ReactDOM.render(
    <BrowserRouter>
        <Switch>
            <Route exact path='/' component={Home} />
            <Route path='/:year/' component={Year}/>
        </Switch>
    </BrowserRouter>,
    document.getElementById('root'));
registerServiceWorker();
