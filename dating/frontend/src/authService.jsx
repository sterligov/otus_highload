import axios, {post} from 'axios';

const role = {
    admin: 'ROLE_ADMIN',
    user: 'ROLE_USER'
};

export const authService = {
    login,
    logout,
    tokenHeader,
    isLogged
};

function login(username, password) {
    return post('/v1/login', {
        username: username,
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

function isLogged() {
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