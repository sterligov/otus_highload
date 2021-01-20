import React from "react";
import axios from "axios";
import {Switch, withRouter} from "react-router-dom";
import Menu from "./Menu";
import {authService} from "../authService";

class Profile extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            user: null,
        }

        this.subscribeHandler = this.subscribeHandler.bind(this);
    }

    subscribeHandler(e) {
        let user = {...this.state.user};
        let url = `v1/friends/${e.target.dataset.user}`;

        if (e.target.dataset.friend == 1) {
            axios.delete(url)
                .then(
                    () => {
                        if (this._isMounted) {
                            user["is_friend"] = 0;
                            this.setState({
                                user: user,
                            });
                        }
                    },
                    () => {
                        if (this._isMounted) {
                            alert('Произошла ошибка при отписке');
                        }
                    }
                )
        } else {
            axios.post(url)
                .then(
                    () => {
                        if (this._isMounted) {
                            user["is_friend"] = 1;
                            this.setState({
                                user: user,
                            });
                        }
                    },
                    () => {
                        if (this._isMounted) {
                            alert('Произошла ошибка при подписке');
                        }
                    }
                )
        }
    }

    componentDidUpdate(prevProps, prevState, snapshot) {
        if (this.props.location.pathname !== prevProps.location.pathname) {
            this.cancelSource.cancel();
            this.fetchData();
        }
    }

    componentDidMount() {
        this.cancelSource = axios.CancelToken.source();
        this._isMounted = true;
        document.title = "Профиль";

        this._isMounted = true;
        this.fetchData();
    }

    componentWillUnmount() {
        this._isMounted = false;
        this.cancelSource.cancel();
    }

    fetchData() {
        let url = `v1/profile`;
        if (this.props.match.params.id) {
            url = `v1/users/${this.props.match.params.id}`;
        }

        axios.get(url)
            .then(
                result => {
                    if (this._isMounted) {
                        this.setState({
                            user: result.data,
                        });
                    }
                },
            );
    }

    render() {
        const user = this.state.user;
        const sex = {
            "M": "Мужской",
            "F": "Женский"
        }

        return (
            <>
                <Menu/>
                <h5>Профиль</h5>
                {user &&
                <div>
                    <p><span className="font-weight-bold">Имя:</span> {user["first_name"]}</p>
                    <p><span className="font-weight-bold">Фамилия:</span>: {user["last_name"]}</p>
                    <p><span className="font-weight-bold">Дата рождения:</span> {user["birthday"]}</p>
                    <p><span className="font-weight-bold">Email:</span> {user["email"]}</p>
                    <p><span className="font-weight-bold">Пол:</span> {sex[user["sex"]]}</p>
                    <p><span className="font-weight-bold">Город:</span> {user["city"]["name"]}</p>
                    <p><span className="font-weight-bold">Интересы:</span> {user["interests"]}</p>
                    {!authService.isUser(user["email"]) &&
                        <button
                            className="btn btn-sm btn-info ml-2"
                            data-user={user["id"]}
                            data-friend={user["is_friend"]}
                            onClick={this.subscribeHandler}>
                            {user["is_friend"] ? "Отписаться" : "Подписаться"}
                        </button>
                    }
                </div>
                }
            </>

        );
    }
}

export default withRouter(Profile);