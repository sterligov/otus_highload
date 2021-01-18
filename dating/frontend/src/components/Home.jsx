import React from "react";
import {Redirect, withRouter} from "react-router-dom";
import {authService} from "../authService";

class Home extends React.Component {
    render() {
        if (!authService.isAuthorized()) {
            return <Redirect to="/sign-in"/>;
        }

        return (
            <>Вы дома</>
        );
    }
}

export default withRouter(Home);