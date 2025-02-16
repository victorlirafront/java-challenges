import React, { createContext, useContext, useState, useEffect } from 'react';
import Cookies from 'js-cookie';
import axios from 'axios';

interface AuthData {
  message: string;
  role: string;
  token: string;
  user: string;
}

interface AuthContextType {
  isAuthenticated: boolean;
  authData: AuthData | null;
  setIsAuthenticated: React.Dispatch<React.SetStateAction<boolean>>;
  login: (data: AuthData) => void;
  logout: () => void;
}

const AuthContext = createContext<AuthContextType | undefined>(undefined);

export const AuthProvider: React.FC<{ children: React.ReactNode }> = ({ children }) => {
  const [isAuthenticated, setIsAuthenticated] = useState<boolean>(false);
  const [authData, setAuthData] = useState<AuthData | null>(null);

  const API_BASE_URL =
  process.env.NODE_ENV === 'development'
    ? process.env.NEXT_PUBLIC_BLOG_API_DEVELOPMENT
    : process.env.NEXT_PUBLIC_BLOG_API_PRODUCTION;

  useEffect(() => {

    const storedAuthData = Cookies.get('auth_data');

    if (storedAuthData) {
      setIsAuthenticated(true);
      setAuthData(JSON.parse(storedAuthData));
    }
    
  }, []);

  const login = (data: AuthData) => {
    setIsAuthenticated(true);
    setAuthData(data);
    Cookies.set('auth_data', JSON.stringify(data));
  };

  const logout = async () => {

    const userString = Cookies.get('auth_data');

    if (userString) {
      try {
        const user = JSON.parse(userString);

        const formData = new FormData();
        formData.append('username', user.user);

        try{
          const response = await axios.post(`${API_BASE_URL}/logout`, formData, {
            headers: {
              'Content-Type': 'multipart/form-data',
            },
            withCredentials: true,
          });
  
          console.log('Logout realizado com sucesso:', response.data);
        }catch(err){
          console.log(err)
        }
      } catch (error) {
        console.error('Erro ao parsear o cookie ou ao fazer o logout:', error);
      }
    } else {
      console.warn("Nenhum cookie 'auth_data' encontrado.");
    }

    Cookies.remove('auth_data');

    setIsAuthenticated(false);
    setAuthData(null);
  };

  return (
    <AuthContext.Provider value={{ isAuthenticated, authData, setIsAuthenticated, login, logout }}>
      {children}
    </AuthContext.Provider>
  );
};

export const useAuth = () => {
  const context = useContext(AuthContext);

  if (!context) {
    throw new Error('useAuth must be used within an AuthProvider');
  }
  return context;
};
