import axios from 'axios';
import { FC, PropsWithChildren, useEffect, useState } from 'react';
import { API_URL } from '../api/api';
import { AuthResponse } from '../services/auth.service';
import { AuthContext } from './AuthContext';

export const AuthProvider: FC<PropsWithChildren> = ({ children }) => {
  const [isAuth, setIsAuth] = useState<boolean>(false);

  const checkAuth = async (): Promise<void> => {
    try {
      const res = await axios.get<AuthResponse>(`${API_URL}/auth/refresh`, {
        withCredentials: true
      });
      console.log('here11');
      localStorage.setItem('token', res.data.accessToken);
      setIsAuth(true);
    } catch (e) {
      console.log(e);
    }
  };

  useEffect(() => {
    checkAuth();
  }, []);

  return (
    <AuthContext.Provider value={{ isAuth, setIsAuth, checkAuth }}>{children}</AuthContext.Provider>
  );
};
