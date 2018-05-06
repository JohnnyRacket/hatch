import React from 'react'
import { Link } from 'react-router-dom'

const NotFound = () => {
    return (
        <article className="container">
            <div className="columns">
                <div className="column is-half is-offset-one-quarter">
                    <div className="card">
                        <div className="card-content">
                            <h1 className="title is-2">Here there be dragons 🐉 </h1>
                            <h3 className="subtitle is-4">(Or probably something just as scary)</h3>
                            <p> 
                                You've wandered off the beaten path and we don't know what you're looking for!
                            </p>
                        </div>
                        <footer className="card-footer">
                            <p className="card-footer-item">
                                <span>
                                    <Link to="/">Go Home</Link>
                                </span>
                            </p>
                        </footer>
                    </div>
                </div>
            </div>
        </article>
    )
};

export default NotFound;

