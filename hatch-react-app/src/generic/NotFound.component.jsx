import React from 'react'
import { Link } from 'react-router-dom'

const NotFound = () => {
    return (
        <article>
            <header>
                <h1>Here there be dragons</h1>
                <h3>Or probably something just as scary</h3>
            </header>
            <section>
                <p>Looks like you're a little lost - maybe go <Link to="/">home</Link>?</p>
            </section>
        </article>
    )
};

export default NotFound;