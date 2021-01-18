import axios, {post} from 'axios';

export const authService = {
    login,
    logout,
    tokenHeader,
    isAuthorized
};

function login(email, password) {
    return post('/v1/sign-in', {
        email: email,
        password: password
    })
        .then(
            res => {
                localStorage.setItem("token", res.data.token);
                axios.defaults.headers.common['Authorization'] = `Bearer ${res.data.token}`;

                return res.data.token;
            }
        );
}

function logout() {
    delete axios.defaults.headers.common['Authorization'];
    localStorage.removeItem("token");
}

function isAuthorized() {
    return Boolean(localStorage.getItem("token"));
}

function tokenHeader() {
    const token = getToken();
    return token ? `Bearer ${token}` : "";
}

function getToken() {
    const token = localStorage.getItem("token");
    return token ? token : '';
}