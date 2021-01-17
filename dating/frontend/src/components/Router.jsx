import React from "react";
import Login from "./Login";
import Home from "./Home";
import Error from "./Error";
import Logout from "./Logout";
import {Route, Switch} from "react-router-dom";
import Profile from "./Profile";
import UserList from "./UserList";

export default class Router extends React.Component {
    render() {
        return (
            <>
                <Switch>
                    <Route exact path="/login">
                        <Login/>
                    </Route>
                    <Route exact path="/">
                        <Home/>
                    </Route>
                    <Route exact path="/logout">
                        <Logout/>
                    </Route>
                    <Route exact path="/users/:id">
                        <Profile/>
                    </Route>
                    <Route exact path="/users">
                        <UserList/>
                    </Route>
                    <Route >
                        <Error error={{status: 404}}/>
                    </Route>
                </Switch>
            </>
        );
    }
}