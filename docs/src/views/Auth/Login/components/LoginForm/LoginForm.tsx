import React, { useState, FormEvent, ChangeEvent } from 'react';
import axios from 'axios';
import { StyledLoginForm } from './LoginForm.styled';
import { useRouter } from 'next/router';
import { LoginResponse } from './LoginForm.types';
import { useAuth } from '@/context/Auth';
import FormModal from '@/components/FormModal/FormModal';

function LoginForm() {
  const [username, setUsername] = useState<string>('');
  const [password, setPassword] = useState<string>('');
  const [usernameError, setUsernameError] = useState<string>('');
  const [passwordError, setPasswordError] = useState<string>('');
  const [modalActive, setModalActive] = useState<boolean>(false);
  const [loading, setLoading] = useState<boolean>(false);
  const { login } = useAuth();
  const router = useRouter();

  const handleUsernameChange = (e: ChangeEvent<HTMLInputElement>) => {
    setUsername(e.target.value);
    setUsernameError('');
  };

  const handlePasswordChange = (e: ChangeEvent<HTMLInputElement>) => {
    setPassword(e.target.value);
    setPasswordError('');
  };

  const handleSubmit = async (e: FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    let hasError = false;

    if (username.length < 8) {
      setUsernameError('Username must be at least 8 characters long.');
      hasError = true;
    }

    if (password.length < 8) {
      setPasswordError('Password must be at least 8 characters long.');
      hasError = true;
    }

    if (hasError) return;

    setLoading(true);

    try {
      const formData = new FormData();
      formData.append('username', username);
      formData.append('password', password);

      const response = await axios.post<LoginResponse>('http://localhost:8080/login', formData, {
        headers: {
          'Content-Type': 'multipart/form-data',
        },
        withCredentials: true,
      });

      setUsernameError('');
      setPasswordError('');
      login(response.data);
      setModalActive(true);
      setTimeout(() => {
        router.push('/auth/profile');
      }, 1000);
    } catch (error) {
      if (axios.isAxiosError(error)) {
        if (error.response) {
          const status = error.response.status;
          if (status === 401 || status === 403) {
            setPasswordError('Invalid username or password.');
          } else {
            setPasswordError('An error occurred. Please try again later.');
          }
        } else {
          setPasswordError('Network error. Please check your connection.');
        }
      } else {
        setPasswordError('An unexpected error occurred.');
      }
      console.error('Error during login:', error);
    } finally {
      setLoading(false);
    }
  };

  return (
    <>
      <FormModal className={modalActive ? 'active' : ''} />
      <StyledLoginForm>
        <form onSubmit={handleSubmit}>
          <h1>Login</h1>
          <div className="form-control">
            <label htmlFor="username">Username</label>
            <input
              type="text"
              id="username"
              placeholder="Enter your username"
              value={username}
              onChange={handleUsernameChange}
            />
            {usernameError && <p className="error-message">{usernameError}</p>}
          </div>
          <div className="form-control">
            <label htmlFor="password">Password</label>
            <input
              type="password"
              id="password"
              placeholder="Enter your password"
              value={password}
              onChange={handlePasswordChange}
            />
            {passwordError && <p className="error-message">{passwordError}</p>}
          </div>

          <button type="submit" disabled={loading}>
            {loading ? 'Carregando...' : 'Submit'}
          </button>
        </form>
      </StyledLoginForm>
    </>
  );
}

export default LoginForm;
