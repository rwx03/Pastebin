import axios, { AxiosResponse, InternalAxiosRequestConfig } from 'axios';
import { AuthResponse } from '../services/auth.service';

export const API_URL = 'http://localhost:8080/api';

const instance = axios.create({
  withCredentials: true,
  baseURL: API_URL
});

instance.interceptors.request.use((config: InternalAxiosRequestConfig) => {
  config.headers!.Authorization = `Bearer ${localStorage.getItem('token')}`;
  return config;
});

instance.interceptors.response.use(
  (config: AxiosResponse) => {
    return config;
  },
  async (error) => {
    const originalRequest = error.config;
    if (error.response.status == 401 && error.config && !error.config._isRetry) {
      originalRequest._isRetry = true;
      try {
        const res = await axios.get<AuthResponse>(`${API_URL}/auth/refresh`, {
          withCredentials: true
        });
        localStorage.setItem('token', res.data.accessToken);

        return instance.request(originalRequest);
      } catch {
        window.location.reload();
        console.log('Unauthorized');
      }
    }
    throw error;
  }
);

export default instance;
