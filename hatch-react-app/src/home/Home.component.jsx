import React, { Component } from 'react'
import { Link } from 'react-router-dom'

class Home extends Component {
    render() {
        return (
            <article className="container">
                <div className="columns">
                    <div className="column is-half is-offset-one-quarter">
                        <div className="card">
                            <div className="card-content">
                                <h1 className="title is-2">Welcome to Hatch! 🥚</h1>
                                <h3 className="subtitle is-4">How did you get here?</h3>
                            </div>
                            <footer className="card-footer">
                                <p className="card-footer-item">
                                    <Link to="/login">
                                        Login
                                    </Link>
                                </p>
                                <p className="card-footer-item">
                                    <Link to="/register">
                                        Register
                                    </Link>
                                </p>
                            </footer>
                        </div>
                    </div>
                </div>
            </article>
        )
    }
}

export default Home

