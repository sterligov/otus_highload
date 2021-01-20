import React from "react";
import axios from "axios";
import Error from "./Error";
import {Link, withRouter} from "react-router-dom";
import ListPlaceholder from "./ListPlaceholder";
import Menu from "./Menu";

class FriendsList extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            error: null,
            users: [],
        };

        this.unsubscribe = this.unsubscribe.bind(this);
    }

    unsubscribe(e) {
        let users = [...this.state.users];

        axios.delete(`v1/friends/${e.target.dataset.user}`)
            .then(
                () => {
                    if (this._isMounted) {
                        users.splice(e.target.dataset.idx, 1);
                        this.setState({
                            users: users,
                        });
                    }
                },
                () => {
                    if (this._isMounted) {
                        alert('Произошла ошибка при удалении');
                    }
                }
            )
    }

    componentDidMount() {
        this.cancelSource = axios.CancelToken.source();
        this._isMounted = true;
        document.title = "Друзья";

        axios.get("v1/friends", {cancelToken: this.cancelSource.token})
            .then(
                result => {
                    if (this._isMounted) {
                        this.setState({users: result.data});
                    }
                },
                err => {
                    if (this._isMounted) {
                        this.setState({
                            items: [],
                            error: err
                        });
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

        return (
            <>
                <Menu/>
                <div>
                    <h5>Друзья</h5>
                    {this.state.users.size === 0 &&
                    <ListPlaceholder number={3}/>
                    }
                    {this.state.users.map((user, idx) =>
                        <div className="mb-3" key={`user_${user["id"]}`}>
                            <Link to={`/profile/${user["id"]}`} rel="noopener noreferrer">
                                {`${user["first_name"]} ${user["last_name"]} (${user["city"]["name"]})`}
                            </Link>
                            <button
                                className="btn btn-sm btn-info ml-2"
                                data-user={user["id"]}
                                data-idx={idx}
                                onClick={this.unsubscribe}>
                                Отписаться
                            </button>
                        </div>
                    )}
                </div>
            </>
        );
    }
}

export default withRouter(FriendsList);
