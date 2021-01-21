import React from "react";
import {Link, Redirect, withRouter} from "react-router-dom";
import {authService} from "../authService";

class Menu extends React.Component {
    render() {
        if (!authService.isAuthorized()) {
            return <Redirect to="/sign-in"/>;
        }

        return (
            <nav className="navbar navbar-expand-lg navbar-light bg-light mb-2">
                <div className="collapse navbar-collapse" id="navbarSupportedContent">
                    <ul className="navbar-nav mr-auto">
                        <li className="nav-item">
                            <Link to="/" className="nav-link">Профиль</Link>
                        </li>
                        <li className="nav-item">
                            <Link to="/friends" className="nav-link">Друзья</Link>
                        </li>
                        <li className="nav-item">
                            <Link to="/users" className="nav-link">Все пользователи</Link>
                        </li>
                        <li className="nav-item">
                            <Link to="/sign-out" className="nav-link">Выйти</Link>
                        </li>
                    </ul>
                    {/*<form className="form-inline my-2 my-lg-0">*/}
                    {/*    <input className="form-control mr-sm-2" type="search" placeholder="Search" aria-label="Search"/>*/}
                    {/*    <button className="btn btn-outline-success my-2 my-sm-0" type="submit">Поиск</button>*/}
                    {/*</form>*/}
                </div>
            </nav>
        );
    }
}

export default withRouter(Menu);