import React, { Component } from 'react';
import { Switch, Route } from 'react-router-dom';
import logo from './logo.svg';
import './App.css';

import asyncComponent from './generic/Async.component';

const AsyncHome = asyncComponent(() => import("./home/Home.component"))
const AsyncNotFound = asyncComponent(() => import("./generic/NotFound.component"));

class App extends Component {
  render() {
    return (
      <main className="App">
        <header className="App-header">
          <img src={logo} className="App-logo" alt="logo" />
          <h1 className="App-title">Welcome to React Boyo</h1>
        </header>
        <section>
          <Switch>
            <Route path="/" exact component={AsyncHome} />
            <Route component={AsyncNotFound} />  
          </Switch>
        </section>
        <footer>
          <p className="App-intro">
            To log out, click here:
          </p>
          <a href="/logout">Log Out</a>
        </footer>    
        </main>
    );
  }
}

export default App;
