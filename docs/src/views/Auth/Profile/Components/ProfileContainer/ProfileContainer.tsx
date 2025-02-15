import React, { useState } from 'react';
import { useAuth } from '@/context/Auth';
import StyledProfileContainer from './ProfileContainer.styled';

function ProfileContainer() {
  const { authData } = useAuth();
  const [isCopied, setIsCopied] = useState(false);

  const handleCopyToken = () => {
    if (authData?.token) {
      navigator.clipboard
        .writeText(authData.token)
        .then(() => {
          setIsCopied(true);
          setTimeout(() => setIsCopied(false), 2000); // Reseta o texto após 2 segundos
        })
        .catch(err => console.error('Erro ao copiar token:', err));
    }
  };

  const codeSnippet = `const fetchPost = async () => {
    try {
      const response = await axios.get('http://localhost:8080/posts/1', {
        headers: {
          Authorization:
            'Bearer ${authData?.token}',
          'Content-Type': 'application/json',
        },
      });
  
      console.log('Response:', response.data);
    } catch (error) {
      console.error('Error fetching post:', error);
    }
  };
  
  fetchPost();`;

  return (
    <StyledProfileContainer>
      <div className="container">
        <h1>Welcome to your user account!</h1>

        <h1>Access token:</h1>
        <div className="token-container">
          <button
            className="btn-copy"
            onClick={handleCopyToken}
            style={{ marginTop: '8px', cursor: 'pointer' }}
          >
            {isCopied ? 'Token Copiado' : 'Copiar Token'}
          </button>
          <code>{authData?.token}</code>
        </div>

        <div className='code-example' style={{width: '100%', overflow: 'auto'}}>
          <h1>Implementation Example:</h1>
          <pre>
            <code>{codeSnippet}</code>
          </pre>
        </div>
      </div>
    </StyledProfileContainer>
  );
}

export default ProfileContainer;
