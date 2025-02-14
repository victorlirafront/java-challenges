import React, { useEffect } from "react";
import { useRouter } from "next/router";
import { useAuth } from "@/context/Auth";
import Cookies from "js-cookie";

const PrivateRoute: React.FC<{ children: React.ReactNode }> = ({ children }) => {
  const { isAuthenticated, setIsAuthenticated } = useAuth();
  const router = useRouter();

  useEffect(() => {
    const sessionToken = Cookies.get("session_token");

    if (!sessionToken) {
      setIsAuthenticated(false);
      router.replace("/auth/login");
    } else {
      setIsAuthenticated(true);
    }
  }, [router, setIsAuthenticated]);

  return isAuthenticated ? <>{children}</> : null;
};

export default PrivateRoute;
