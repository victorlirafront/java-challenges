import React, { useState, useEffect } from 'react';
import axios, { AxiosError } from 'axios';
import { useAuth } from '@/context/Auth';
import { useRouter } from 'next/router';
import LoadingSpinner from '@/components/LoadingSpinner/LoadingSpinner';

const ProtectedRoute: React.FC<{ children: React.ReactNode }> = ({ children }) => {
  const { isAuthenticated, setIsAuthenticated } = useAuth();
  const router = useRouter();
  const [loading, setLoading] = useState<boolean>(true);

  const API_BASE_URL =
    process.env.NODE_ENV === 'development'
      ? process.env.NEXT_PUBLIC_BLOG_API_DEVELOPMENT
      : process.env.NEXT_PUBLIC_BLOG_API_PRODUCTION;

  useEffect(() => {
    const fetchProtectedRoute = async () => {
      try {
        const response = await axios.post(
          `${API_BASE_URL}/protected`,
          {},
          {
            withCredentials: true,
          },
        );

        console.log('✅ Acesso autorizado:', response.data);
        setIsAuthenticated(true);
      } catch (err) {
        if (err instanceof AxiosError) {
          console.log(err);
          if (err.response?.status === 401) {
            setIsAuthenticated(false);
            router.replace('/auth/login');
          }
        } else {
          console.error('Erro desconhecido', err);
        }
      } finally {
        setLoading(false);
      }
    };

    fetchProtectedRoute();
  }, [setIsAuthenticated, router]);

  if (loading) {
    return <LoadingSpinner />;
  }

  return isAuthenticated ? <>{children}</> : null;
};

export default ProtectedRoute;
