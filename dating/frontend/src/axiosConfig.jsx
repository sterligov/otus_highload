import axios from "axios";
import {authService} from "./authService";

export const axiosConfig = {
    register
};

function register() {
    // axios.defaults.baseURL = 'http://dating-api:8082';
    axios.defaults.headers.common["Content-Type"] = "application/json; charset=utf-8";
    axios.defaults.headers.common["Accept"] = "application/json";

    if (authService.tokenHeader()) {
        axios.defaults.headers.common['Authorization'] = authService.tokenHeader();
    }
}