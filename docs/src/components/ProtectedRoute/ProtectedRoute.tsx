import React, { useState, useEffect } from 'react';
import axios from 'axios';
import { useAuth } from '@/context/Auth';
import { useRouter } from 'next/router';

const ProtectedRoute: React.FC<{ children: React.ReactNode }> = ({ children }) => {
  const { isAuthenticated, setIsAuthenticated } = useAuth();
  const router = useRouter();
  const [loading, setLoading] = useState<boolean>(true);

  useEffect(() => {
    const fetchProtectedRoute = async () => {
      try {
        try{
          const response = await axios.post('http://localhost:8080/protected', {}, { 
            withCredentials: true,
          });

          console.log('✅ Acesso autorizado:', response.data);
          setIsAuthenticated(true);
        }catch(err){
          console.log(err)
          router.replace('/auth/login');
        }
      } catch (error: any) {
        console.error('🔴 Erro ao acessar rota protegida:', error);

        if (error.response?.status === 401) {
          console.warn('🔴 Usuário não autenticado. Redirecionando para login...');
          setIsAuthenticated(false);
          router.replace('/auth/login');
        }
      } finally {
        setLoading(false);
      }
    };

    fetchProtectedRoute();
  }, [setIsAuthenticated, router]);

  if (loading) {
    return <div>Loading...</div>;
  }

  return isAuthenticated ? <>{children}</> : null;
};

export default ProtectedRoute;
