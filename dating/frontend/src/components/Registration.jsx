import React from "react";
import {Link, withRouter} from "react-router-dom";
import axios from "axios";
import ReactFormInputValidation from "react-form-input-validation";

class Registration extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            isFormDisabled: false,
            errors: {},
            cities: [],
            fields: {
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

        this.form = new ReactFormInputValidation(this);
        this.form.useRules({
            name: "required",
            city_id: "required",
            password: "required",
            first_name: "required",
            last_name: "required",
            sex: "required",
            interests: "required",
            birthday: "required",
            email: "required|email",
        });

        this.handleSubmit = this.handleSubmit.bind(this);

        this.form.onformsubmit = (fields) => {
            console.log(fields);
        }
    }

    handleSubmit(e) {
        this.form.
        e.preventDefault();

        this.setState({errors: {}});
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

    // handleChange(e) {
    //     let user = {...this.state.user};
    //     const inputName = e.target.getAttribute("name");
    //     fields[inputName] = e.target.value;
    //     this.setState({user: user, error: null});
    // }

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
                          onSubmit={this.form.handleSubmit}>

                        <label htmlFor="first_name" className="mt-2 font-weight-bold">Имя</label>
                        <input className="form-control"
                               type="text"
                               name="first_name"
                               onBlur={this.form.handleBlurEvent}
                               onChange={this.form.handleChangeEvent}
                               value={this.state.fields["first_name"]}/>
                        <label className="error d-block text-danger">
                            {this.state.errors["first_name"] ? this.state.errors["first_name"] : ""}
                        </label>

                        <label htmlFor="last_name" className="font-weight-bold">Фамилия</label>
                        <input className="form-control"
                               type="text"
                               name="last_name"
                               onBlur={this.form.handleBlurEvent}
                               onChange={this.form.handleChangeEvent}
                               value={this.state.fields["last_name"]}/>
                        <label className="error d-block text-danger">
                            {this.state.errors["last_name"] ? this.state.errors["last_name"] : ""}
                        </label>

                        <label htmlFor="email" className="mt-2 font-weight-bold">Email</label>
                        <input className="form-control"
                               type="text"
                               name="email"
                               onBlur={this.form.handleBlurEvent}
                               onChange={this.form.handleChangeEvent}
                               value={this.state.fields["email"]}/>
                        <label className="error d-block text-danger">
                            {this.state.errors["email"] ? this.state.errors["email"] : ""}
                        </label>

                        <label htmlFor="password" className="font-weight-bold">Пароль</label>
                        <input className="form-control"
                               type="password"
                               name="password"
                               onBlur={this.form.handleBlurEvent}
                               onChange={this.form.handleChangeEvent}
                               value={this.state.fields["password"]}/>
                        <label className="error d-block text-danger">
                            {this.state.errors["password"] ? this.state.errors["password"] : ""}
                        </label>

                        <label htmlFor="birthday" className="font-weight-bold">Дата рождения</label>
                        <input className="form-control"
                               type="date"
                               name="birthday"
                               onBlur={this.form.handleBlurEvent}
                               onChange={this.form.handleChangeEvent}
                               value={this.state.fields["birthday"]}/>
                        <label className="error d-block text-danger">
                            {this.state.errors["birthday"] ? this.state.errors["birthday"] : ""}
                        </label>

                        <label htmlFor="city" className="font-weight-bold">Город</label>
                        <select name="city_id" className="form-control"
                                onBlur={this.form.handleBlurEvent}
                                onChange={this.form.handleChangeEvent}
                                value={this.state.fields["city_id"]}>
                            {this.state.cities.map(c => {
                                return <option value={c.id}>{c.value}</option>;
                            })}
                        </select>

                        <label htmlFor="sex" className="font-weight-bold">Пол</label>
                        <select name="sex" className="form-control"
                                onBlur={this.form.handleBlurEvent}
                                onChange={this.form.handleChangeEvent}
                                value={this.state.fields["sex"]}>
                            <option value="M">Мужской</option>;
                            <option value="F">Женский</option>;
                        </select>

                        <label htmlFor="interests" className="font-weight-bold">Интересы</label>
                        <textarea className="form-control"
                               name="interests"
                               onBlur={this.form.handleBlurEvent}
                               onChange={this.form.handleChangeEvent}
                               value={this.state.fields["interests"]}/>
                        <label className="error d-block text-danger">
                            {this.state.errors["interests"] ? this.state.errors["interests"] : ""}
                        </label>

                        <button className="btn btn-info mt-2" type="submit">
                            Зарегистрироваться
                        </button>

                        <Link to="/sign-in"><button className="btn btn-info mt-2 ml-2">Назад</button></Link>
                    </form>
                </fieldset>
            </div>
        );
    }
}

export default withRouter(Registration);