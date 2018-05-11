import React, { Component } from 'react'
import { Link, Route } from 'react-router-dom'
import asyncComponent from '../generic/async.component';
import './home.css';

const AsyncLogin = asyncComponent(() => import("./login/login.component"))
const AsyncRegister = asyncComponent(() => import("./registration/register.component"))

class Home extends Component {
    render() {
        return (
            <article className="container home">
                <div className="columns">
                    <div className="column is-half is-offset-one-quarter">
                        <div className="card header-card">
                            <div className="card-content">
                                <h1 className="title is-2">Welcome to Hatch! <span aria-label="A sweet egg (unhatched)" role="img">🥚</span></h1>
                                <Route path="/" exact component={AsyncLogin} />
                                <Route path="/login" exact component={AsyncLogin} />
                                <Route path="/register" exact component={AsyncRegister} />
                            </div>
                            <footer className="card-footer">
                                {
                                     this.props.location.pathname === "/login" ?
                                        <p className="card-footer-item is-centered">
                                            Not in our crew? <Link to="/register">Sign up</Link>
                                        </p>
                                        :
                                        <p className="card-footer-item is-centered">
                                            Already one of the club? <Link to="/login">Log in</Link>
                                        </p>
                                }
                            </footer>
                        </div>
                    </div>
                </div>
            </article>
        )
    }
}

export default Home

