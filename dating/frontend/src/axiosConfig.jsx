import axios from "axios";
import {authService} from "./authService";

export const axiosConfig = {
    register
};

function register() {
    // axios.defaults.baseURL = process.env.REACT_APP_API_ENDPOINT;
    axios.defaults.baseURL = 'http://localhost:8082';
    axios.defaults.headers.common["Content-Type"] = "application/ld+json; charset=utf-8";
    axios.defaults.headers.common["Accept"] = "application/ld+json";

    if (authService.tokenHeader()) {
        axios.defaults.headers.common['Authorization'] = authService.tokenHeader();
    }
}