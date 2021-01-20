import React from "react";
import {Link, Redirect, withRouter} from "react-router-dom";
import {authService} from "../authService";

class Login extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            isFormDisabled: false,
            error: null,
            user: {
                email: "",
                password: ""
            }
        };

        this.handleChange = this.handleChange.bind(this);
        this.handleSignIn = this.handleSignIn.bind(this);
    }

    handleChange(e) {
        let user = {...this.state.user};
        const inputName = e.target.getAttribute("name");
        user[inputName] = e.target.value;
        this.setState({user: user, error: null});
    }

    handleSignIn(e) {
        e.preventDefault();

        this.setState({
            isFormDisabled: true,
            error: null,
        });

        authService.login(this.state.user.email, this.state.user.password)
            .then(
                () => {
                    this.setState({isFormDisabled: false});
                },
                err => {
                    this.setState({
                        isFormDisabled: false,
                        error: err.response.data.message
                    });
                }
            )
    }

    render() {
        if (authService.isAuthorized()) {
            return <Redirect to="/"/>;
        }

        const user = this.state.user;

        return (
            <div className="d-flex justify-content-center">
                <fieldset disabled={this.state.isFormDisabled} className="p-2">
                    <h2>Авторизация</h2>
                    <form id="login" className="form-group">
                        <label htmlFor="email">Email</label>
                        <input type="text"
                               className="form-control"
                               name="email"
                               value={user.email}
                               onChange={this.handleChange}/>

                        <label htmlFor="password" className="mt-2">Пароль</label>
                        <input type="password"
                               className="form-control"
                               name="password"
                               value={user.password}
                               onChange={this.handleChange}/>

                        <label className="error d-block text-danger">
                            {this.state.error ? this.state.error : ""}
                        </label>

                        {/*{this.state.error &&*/}
                        {/*<ValidationError error={"Неверный логин или пароль"}/>*/}
                        {/*}*/}

                        <input className="btn-info btn mt-2"
                               type="submit"
                               value="Войти"
                               onClick={this.handleSignIn}/>
                        <Link to={`/sign-up`} className="font-weight-bold">
                            <button className="btn-info btn mt-2 ml-3">
                                Регистрация
                            </button>
                        </Link>
                    </form>
                </fieldset>
            </div>
        );
    }
}

export default withRouter(Login);