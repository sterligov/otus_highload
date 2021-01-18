import React from "react";
import {Redirect, withRouter} from "react-router-dom";
import {authService} from "../authService";
import Menu from "./Menu";

class Home extends React.Component {
    render() {
        if (!authService.isAuthorized()) {
            return <Redirect to="/sign-in"/>;
        }

        return (
            <>
                <Menu/>
                <h2>Вы дома</h2>
            </>
        );
    }
}

export default withRouter(Home);