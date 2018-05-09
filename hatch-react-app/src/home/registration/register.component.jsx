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
            <div className="columns">
                <div className="column">
                    <div className="field">
                        <p className="control has-icons-left has-icons-right">
                            <input className="input is-large" type="text" placeholder="First Name" />
                            <span className="icon is-small is-left">
                                <i className="far fa-user"></i>
                            </span>
                        </p>
                    </div>
                </div>
                <div className="column">
                    <div className="field">
                        <p className="control has-icons-left has-icons-right">
                            <input className="input is-large" type="text" placeholder="Last Name" />
                            <span className="icon is-small is-left">
                                <i className="far fa-user"></i>
                            </span>
                        </p>
                    </div>
                </div>
            </div>
            <div className="field">
                <p className="control">
                    <button type="submit" className="button is-primary is-fullwidth is-large">Register</button>
                </p>
            </div>
        </form>
    );
}

export default Login;