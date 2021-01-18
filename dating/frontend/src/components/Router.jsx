import React from "react";
import Login from "./Login";
import Home from "./Home";
import Error from "./Error";
import Logout from "./Logout";
import {Route, Switch} from "react-router-dom";
import Profile from "./Profile";
import UserList from "./UserList";
import Registration from "./Registration";

export default class Router extends React.Component {
    render() {
        return (
            <>
                <Switch>
                    <Route exact path="/sign-in">
                        <Login/>
                    </Route>
                    <Route exact path="/sign-out">
                        <Logout/>
                    </Route>
                    <Route exact path="/sign-up">
                        <Registration/>
                    </Route>
                    <Route exact path="/">
                        <Home/>
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