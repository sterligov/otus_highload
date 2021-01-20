import React from "react";
import axios from "axios";
import Error from "./Error";
import {Link, withRouter} from "react-router-dom";
import ListPlaceholder from "./ListPlaceholder";
import Menu from "./Menu";
import {authService} from "../authService";

class UserList extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            users: [],
            friends: {},
        };
        this.subscribeHandler = this.subscribeHandler.bind(this);
    }

    subscribeHandler(e) {
        let friends = {...this.state.friends};
        let url = `v1/friends/${e.target.dataset.user}`;

        if (e.target.dataset.friend == 1) {
            delete friends[e.target.dataset.user];
            axios.delete(url)
                .then(
                    () => {
                        if (this._isMounted) {
                            this.setState({
                                friends: friends,
                            });
                        }
                    },
                    () => {
                        if (this._isMounted) {
                            alert("Произошла ошибка при отписке");
                        }
                    }
                )
        } else {
            friends[e.target.dataset.user] = true;
            axios.post(url)
                .then(
                    () => {
                        if (this._isMounted) {
                            this.setState({
                                friends: friends,
                            });
                        }
                    },
                    () => {
                        if (this._isMounted) {
                            alert("Произошла ошибка при подписке");
                        }
                    }
                )
        }
    }

    componentDidMount() {
        this.cancelSource = axios.CancelToken.source();
        this._isMounted = true;
        document.title = "Список пользователей";

        axios.get("v1/friends", {cancelToken: this.cancelSource.token})
            .then(
                result => {
                    if (this._isMounted) {
                        let friends = {};
                        for (let i = 0; i < result.data.length; i++) {
                            let u = result.data[i];
                            friends[u["id"]] = true;
                        }
                        this.setState({friends: friends,});
                    }
                },
                () => {
                    if (this._isMounted) {
                        alert("Произошла ошибка при обращении к серверу");
                    }
                }
            );

        axios.get("v1/users", {cancelToken: this.cancelSource.token})
            .then(
                result => {
                    if (this._isMounted) {
                        this.setState({
                            users: result.data,
                        });
                    }
                },
                () => {
                    if (this._isMounted) {
                        alert("Произошла ошибка при обращении к серверу");
                    }
                }
            );
    }

    componentWillUnmount() {
        this._isMounted = false;
        this.cancelSource.cancel();
    }

    render() {
        if (this.state.error) {
            return <Error err={this.state.error}/>
        }

        const friends = this.state.friends;

        return (
            <>
                <Menu/>
                <div>
                    <h5>Список пользователей</h5>
                    {this.state.users.length === 0 &&
                    <ListPlaceholder number={3}/>
                    }
                    {this.state.users.map(user =>
                        <div className="mb-3" key={`user_${user["id"]}`}>
                            <Link to={`/profile/${user["id"]}`} rel="noopener noreferrer">
                                {`${user["first_name"]} ${user["last_name"]} (${user["city"]["name"]})`}
                            </Link>
                            {!authService.isUser(user["email"]) &&
                                <button
                                    className="btn btn-sm btn-info ml-2"
                                    data-user={user["id"]}
                                    data-friend={friends[user["id"]] === undefined ? 0 : 1}
                                    onClick={this.subscribeHandler}>
                                    {friends[user["id"]] === undefined ? "Подписаться" : "Отписаться"}
                                </button>
                            }
                        </div>
                    )}
                </div>
            </>
        );
    }
}

export default withRouter(UserList);
