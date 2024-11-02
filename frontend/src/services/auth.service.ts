import axios from 'axios';
import api, { API_URL } from '../api/api';

export interface AuthResponse {
  refreshToken: string;
  accessToken: string;
  email: string;
}

class AuthService {
  async login(email: string, password: string) {
    try {
      const res = await api.post<AuthResponse>('/auth/login', {
        email: email,
        password: password
      });

      localStorage.setItem('token', res.data.accessToken);
    } catch (e) {
      console.log(e);
      throw e;
    }
  }

  async register(email: string, password: string) {
    try {
      const res = await api.post<AuthResponse>('/auth/register', {
        email: email,
        password: password
      });

      localStorage.setItem('token', res.data.accessToken);
    } catch (e) {
      console.log(e);
      throw e;
    }
  }

  async checkAuth(): Promise<boolean> {
    try {
      const res = await axios.get<AuthResponse>(`${API_URL}/refresh`, { withCredentials: true });

      localStorage.setItem('token', res.data.accessToken);

      return true;
    } catch (e) {
      console.log(e);
      return false;
    }
  }
}

export const authService = new AuthService();
