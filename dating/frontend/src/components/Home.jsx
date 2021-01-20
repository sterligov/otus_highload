import React from "react";
import {Redirect, Route, Switch, withRouter} from "react-router-dom";
import {authService} from "../authService";
import Menu from "./Menu";

class Home extends React.Component {
    render() {
        if (!authService.isAuthorized()) {
            return <Redirect to="/sign-in"/>;
        }

        return (
            <div className="w-50">
                <Menu/>
            </div>
        );
    }
}

export default withRouter(Home);