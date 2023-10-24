import axios from "axios"
import { getCookie } from "./cookie"

const instance = axios.create({
  baseURL: "http://localhost:4000/v1",
});

instance.interceptors.request.use(
  config => {
    const token = getCookie("token");
    if (token) {
      config.headers.Authorization = `Bearer ${token}`;
    }
    return config;
  },
  err => {
    return Promise.reject(err);
  }
);

export default instance;
