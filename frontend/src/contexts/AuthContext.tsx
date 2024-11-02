import React, { createContext } from 'react';

interface AuthContextType {
  isAuth: boolean;
  setIsAuth: React.Dispatch<React.SetStateAction<boolean>>;
  checkAuth: () => void;
}

const defaultAuthContext: AuthContextType = {
  isAuth: false,
  setIsAuth: () => {},
  checkAuth: () => {}
};

export const AuthContext = createContext<AuthContextType>(defaultAuthContext);
