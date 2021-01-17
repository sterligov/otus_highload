import React from "react";
import {authService} from "../authService";
import {Redirect} from "react-router-dom";

export default class Logout extends React.Component {
    render() {
        authService.logout();
        return <Redirect to="/login"/>;
    }
}