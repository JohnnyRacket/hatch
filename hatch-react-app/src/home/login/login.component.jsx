import React from 'react';

const Login = () => {
    return (
        <form>
            <div className="field">
                <p className="control has-icons-left has-icons-right">
                    <input className="input is-large" type="email" placeholder="Email" />
                    <span className="icon is-small is-left">
                        <i className="far fa-envelope"></i>
                    </span>
                </p>
            </div>
            <div className="field">
                <p className="control">
                    <button type="submit" className="button is-primary is-fullwidth is-large">Log In</button>
                </p>
            </div>
        </form>
    );
}

export default Login;