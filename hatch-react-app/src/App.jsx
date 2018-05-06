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
      <main className="App has-background-primary">
        <header className="App-header">
          <h1 className="App-title">Welcome to Hatch</h1>
        </header>
        <section className="section">
          <Switch>
            <Route path="/" exact component={AsyncHome} />
            <Route component={AsyncNotFound} />  
          </Switch>
        </section>
        <footer className="footer">
          <div className="container">
            <div className="content has-text-centered">
              <p className="App-intro">
                To log out, click here: <a href="/logout">Log Out</a>
              </p>
            </div>
          </div>
        </footer>
        </main>
    );
  }
}

export default App;
