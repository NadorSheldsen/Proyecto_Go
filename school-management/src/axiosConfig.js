// src/axiosConfig.js
import axios from 'axios';

const apiUrl = 'http://localhost:8000';

const axiosInstance = axios.create({
  baseURL: apiUrl,
});

export default axiosInstance;
