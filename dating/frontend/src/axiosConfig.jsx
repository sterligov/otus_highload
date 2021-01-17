import axios from "axios";
import {authService} from "./authService";

export const axiosConfig = {
    register
};

function register() {
    axios.defaults.baseURL = process.env.REACT_APP_API_ENDPOINT;
    axios.defaults.headers.common["Content-Type"] = "application/ld+json; charset=utf-8";
    axios.defaults.headers.common["Accept"] = "application/ld+json";

    if (authService.tokenHeader()) {
        axios.defaults.headers.common['Authorization'] = authService.tokenHeader();
    }

    axios.interceptors.response.use(
        res => res,
        err => {
            if (err.response && (err.response.status === 403 || err.response.status === 401)) {
                if (err.response.status === 401) {
                    authService.logout();
                }

                if (window.location.pathname.indexOf("/login") === -1) {
                    window.location = "/login";
                }
            }

            throw err;
        });
}