import React from "react";
import {Link, withRouter} from "react-router-dom";
import Dropdown from "react-dropdown";
import axios from "axios";

class Registration extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            isFormDisabled: false,
            error: null,
            cities: [],
            user: {
                email: "",
                password: "",
                first_name: "",
                last_name: "",
                sex: "M",
                interests: "",
                birthday: "",
                city_id: 1,
            }
        };

        this.handleSubmit = this.handleSubmit.bind(this);
        this.handleChange = this.handleChange.bind(this);
    }

    handleSubmit(e) {
        e.preventDefault();

        this.setState({validationErrors: {}});
        const user = {...this.state.user};
        console.log(user);

        axios.post("v1/sign-up", user)
            .then(
                () => {
                    if (this._isMounted) {
                        this.setState({isFormDisabled: false});
                    }
                },
                () => {
                    if (this._isMounted) {
                        this.setState({
                            isFormDisabled: false,
                        });
                    }
                }
            );
    }

    handleChange(e) {
        let user = {...this.state.user};
        const inputName = e.target.getAttribute("name");
        user[inputName] = e.target.value;
        this.setState({user: user, error: null});
    }

    componentDidMount() {
        this.cancelSource = axios.CancelToken.source();
        this._isMounted = true;
        document.title = "Регистрация";

        axios.get("v1/cities", {cancelToken: this.cancelSource.token})
            .then(
                result => {
                    if (this._isMounted) {
                        let cities = [];
                        for (let i = 0; i < result.data.length; i++) {
                            const city = result.data[i];
                            cities.push({
                                id: city["id"],
                                value: `${city["name"]} (${city["country"]["name"]})`
                            })
                        }
                        this.setState({
                            cities: cities,
                        });
                    }
                },
                err => {
                    if (this._isMounted) {
                        this.setState({
                            cities: [],
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
        return (
            <div>
                <fieldset disabled={this.state.isFormDisabled}>
                    <form name="article" className="form-group" style={{width: "300px"}}
                          onSubmit={this.handleSubmit}>

                        <label htmlFor="first_name" className="mt-2 font-weight-bold">
                            Имя
                        </label>
                        <input className="form-control"
                               type="text"
                               name="first_name"
                               onChange={this.handleChange}/>

                        <label htmlFor="second_name" className="font-weight-bold">Фамилия</label>
                        <input className="form-control"
                               type="text"
                               name="second_name"
                               onChange={this.handleChange}/>

                        <label htmlFor="email" className="mt-2 font-weight-bold">Email</label>
                        <input className="form-control"
                               type="text"
                               name="email"
                               onChange={this.handleChange}/>

                        <label htmlFor="password" className="font-weight-bold">Пароль</label>
                        <input className="form-control"
                               type="text"
                               name="password"
                               onChange={this.handleChange}/>

                        <label htmlFor="birthday" className="font-weight-bold">Дата рождения</label>
                        <input className="form-control"
                               type="date"
                               name="birthday"
                               onChange={this.handleChange}/>

                        <label htmlFor="city" className="font-weight-bold">Город</label>
                        <select name="city_id" className="form-control" onChange={this.handleChange}>
                            {this.state.cities.map(c => {
                                return <option value={c.id}>{c.value}</option>;
                            })}
                        </select>

                        <label htmlFor="sex" className="font-weight-bold">Пол</label>
                        <select name="sex" className="form-control" onChange={this.handleChange}>
                            <option value="M">Мужской</option>;
                            <option value="F">Женский</option>;
                        </select>

                        <label htmlFor="interests" className="font-weight-bold">Интересы</label>
                        <textarea className="form-control"
                               name="interests"
                               onChange={this.handleChange}/>

                        <input className="btn btn-info mt-2"
                               type="submit"
                               value="Зарегистрироваться"
                               onClick={this.handleSubmit}
                        />

                        <Link to="/sign-in"><button className="btn btn-info mt-2 ml-2">Назад</button></Link>
                    </form>
                </fieldset>
            </div>
        );
    }
}

export default withRouter(Registration);