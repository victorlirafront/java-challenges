import React, { useState, FormEvent, ChangeEvent } from 'react';
import axios from 'axios';
import { StyledLoginForm } from './LoginForm.styled';
import { LoginResponse } from './LoginForm.types';
import { BLOG_API } from '@/constants/endpoints';

function LoginForm() {
  const [username, setUsername] = useState<string>('');
  const [password, setPassword] = useState<string>('');
  const [error, setError] = useState<string>('');

  const handleUsernameChange = (e: ChangeEvent<HTMLInputElement>) => {
    setUsername(e.target.value);
  };

  const handlePasswordChange = (e: ChangeEvent<HTMLInputElement>) => {
    setPassword(e.target.value);
  };

  const handleSubmit = async (e: FormEvent<HTMLFormElement>) => {
    e.preventDefault();

    if (username.length < 8 || password.length < 8) {
      setError('Username and password must be at least 8 characters long.');
      return;
    }

    try {
      const formData = new FormData();
      formData.append('username', username);
      formData.append('password', password);

      const response = await axios.post<LoginResponse>(`${BLOG_API}/login`, formData, {
        headers: {
          'Content-Type': 'multipart/form-data',
        },
      });

      console.log('Login successful:', response.data);
      setError('');
      if (response.data.token) {
        localStorage.setItem('accessToken', response.data.token);
      }
    } catch (error) {
      if (axios.isAxiosError(error)) {
        if (error.response) {
          setError(error.response.data.error || 'Invalid username or password.');
        } else {
          setError('An error occurred during login.');
        }
      } else {
        setError('An unexpected error occurred.');
      }
      console.error('Error during login:', error);
    }
  };

  return (
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
        </div>

        {error && <p className="error-message">{error}</p>}

        <button type="submit">Submit</button>
      </form>
    </StyledLoginForm>
  );
}

export default LoginForm;
