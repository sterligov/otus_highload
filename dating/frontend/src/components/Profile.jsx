import React from "react";
import axios from "axios";
import {Link} from "react-router-dom";

class Profile extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            user: null,
            error: null,
        }
    }

    componentDidMount() {
        this.cancelSource = axios.CancelToken.source();
        this._isMounted = true;
        document.title = "Профиль";

        axios.get(`v1/users/${this.props.id}`, {cancelToken: this.cancelSource.token})
            .then(
                result => {
                    if (this._isMounted) {
                        this.setState({
                            user: result.data,
                        });
                    }
                },
                err => {
                    if (this._isMounted) {
                        this.setState({
                            user: null,
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
        const user = this.state.user;
        const sex = {
            "M": "Мужской",
            "F": "Женский"
        }

        return (
            <div>
                <p>Имя: {user["first_name"]}</p>
                <p>Фамилия: {user["last_name"]}</p>
                <p>Дата рождения: {user["birthday"]}</p>
                <p>Email: {user["email"]}</p>
                <p>Пол: {sex[user["sex"]]}</p>
                <p>Город: {user["city"]["name"]}</p>
                <p>Интересы: {user["interests"]}</p>
                <p><Link to="/subscribers">Подписчики</Link></p>
            </div>
        );
    }
}

export default Profile;