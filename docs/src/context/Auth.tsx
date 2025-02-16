import React, { createContext, useContext, useState, useEffect } from "react";
import Cookies from "js-cookie";

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

  useEffect(() => {
    const storedAuthData = Cookies.get("auth_data");
    
    if (storedAuthData) {
      setIsAuthenticated(true);
      setAuthData(JSON.parse(storedAuthData));
    }
  }, []);

  const login = (data: AuthData) => {
    setIsAuthenticated(true);
    setAuthData(data);
    Cookies.set("auth_data", JSON.stringify(data));
  };

  const logout = () => {
    Cookies.remove("auth_data");
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
    throw new Error("useAuth must be used within an AuthProvider");
  }
  return context;
};
