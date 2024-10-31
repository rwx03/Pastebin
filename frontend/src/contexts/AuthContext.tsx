import React, {
  createContext,
  FC,
  PropsWithChildren,
  useState
} from 'react';

interface IUser {
  id: number;
  name: string;
}

interface AuthContextType {
  user: IUser | null;
  isAuth: boolean;
  setUser: React.Dispatch<
    React.SetStateAction<IUser | null>
  >;
  setIsAuth: React.Dispatch<React.SetStateAction<boolean>>;
}

const defaultAuthContext: AuthContextType = {
  user: null,
  isAuth: false,
  setUser: () => {},
  setIsAuth: () => {}
};

export const AuthContext = createContext<AuthContextType>(
  defaultAuthContext
);

export const AuthProvider: FC<PropsWithChildren> = ({
  children
}) => {
  const [user, setUser] = useState<IUser | null>(null);
  const [isAuth, setIsAuth] = useState<boolean>(false);

  return (
    <AuthContext.Provider
      value={{
        user,
        isAuth,
        setUser,
        setIsAuth
      }}>
      {children}
    </AuthContext.Provider>
  );
};
