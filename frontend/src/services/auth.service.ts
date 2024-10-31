import api from '../api/api';

interface ILoginResponse {
  refreshToken: string;
  accessToken: string;
  email: string;
}

class AuthService {
  async login(email: string, password: string) {
    const res = await api.post<ILoginResponse>('/auth/login', {
      email: email,
      password: password
    });

    localStorage.setItem('token', res.data.accessToken);
  }
}

export const authService = new AuthService();
