import React from 'react';
import { useRouter } from 'next/router';
import { useAuth } from '@/context/Auth';
import StyledProfileContainer from './ProfileContainer.styled';

function ProfileContainer() {
  const { logout, authData } = useAuth();
  const router = useRouter();

  const handleLogout = () => {
    logout();
    router.push('/login');
  };

  const handleCopyToken = () => {
    if (authData?.token) {
      navigator.clipboard
        .writeText(authData.token)
        .then(() => alert('Token copiado para a área de transferência!'))
        .catch(err => console.error('Erro ao copiar token:', err));
    }
  };

  return (
    <StyledProfileContainer>
      <div className='container'>
        <h1>Welcome to your user account!</h1>
        <button onClick={handleLogout}>Logout</button>

        <h1>Access token:</h1>
        <div>
          <button onClick={handleCopyToken} style={{ marginTop: '8px', cursor: 'pointer' }}>
            Copiar Token
          </button>
          <code
            style={{
              wordBreak: 'break-all',
              whiteSpace: 'pre-wrap',
              overflowWrap: 'break-word',
              display: 'block',
              padding: '8px',
              backgroundColor: '#f4f4f4',
              borderRadius: '4px',
              overflowX: 'auto',
            }}
          >
            {authData?.token}
          </code>
        </div>
      </div>
    </StyledProfileContainer>
  );
}

export default ProfileContainer;
