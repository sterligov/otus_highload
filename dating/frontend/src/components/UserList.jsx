import React from "react";
import axios from "axios";
import Error from "./Error";
import {Link, withRouter} from "react-router-dom";
import ListPlaceholder from "./ListPlaceholder";

class UserList extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            error: null,
            users: [],
        };

        this.handleSubscribe = this.handleSubscribe.bind(this);
    }

    componentDidMount() {
        this.cancelSource = axios.CancelToken.source();
        this._isMounted = true;
        document.title = "Список пользователей";

        axios.get("v1/users", {cancelToken: this.cancelSource.token})
            .then(
                result => {
                    if (this._isMounted) {
                        this.setState({
                            users: result.data,
                        });
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

    handleSubscribe() {
        const id = 1;
        axios.post(`v1/friends/${id}`);
    }

    render() {
        if (this.state.error) {
            return <Error err={this.state.error}/>
        }

        return (
            <div>
                <h5>Список пользователей</h5>

                {this.state.users.length === 0 &&
                <ListPlaceholder number={3}/>
                }

                {this.state.users.map(user =>
                    <div className="row mb-3" key={`user_${user.id}`}>
                        <Link to={`/users/${user["id"]}`} target="_blank" rel="noopener noreferrer">
                            {`${user["first_name"]} ${user["last_name"]} (${user["city"]["name"]})`}
                        </Link>
                        <button className="btn btn-info" onClick={this.handleSubscribe}>
                            Подписаться
                        </button>
                    </div>
                )}
            </div>
        );
    }
}

export default withRouter(UserList);
