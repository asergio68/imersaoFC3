import axios from "axios";

const http = axios.create({
  baseURL: "http://host.docker.internal:3000",
});

export default http;
