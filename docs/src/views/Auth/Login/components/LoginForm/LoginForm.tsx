import React, { useState, FormEvent, ChangeEvent } from 'react';
import axios from 'axios';
import { StyledLoginForm } from './LoginForm.styled';
import { useRouter } from 'next/router';
import { LoginResponse } from './LoginForm.types';
import { useAuth } from '@/context/Auth';

function LoginForm() {
  const [username, setUsername] = useState<string>('');
  const [password, setPassword] = useState<string>('');
  const [usernameError, setUsernameError] = useState<string>('');
  const [passwordError, setPasswordError] = useState<string>('');
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

      console.log('Login successful:', response.data);
      setUsernameError('');
      setPasswordError('');
      login(response.data);
      router.push('/auth/profile');
    } catch (error) {
      if (axios.isAxiosError(error)) {
        const errorMessage = error.response?.data.error || 'Invalid username or password.';
        setUsernameError(errorMessage.includes('username') ? errorMessage : '');
        setPasswordError(errorMessage.includes('password') ? errorMessage : '');
      } else {
        setUsernameError('An unexpected error occurred.');
        setPasswordError('An unexpected error occurred.');
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

        <button type="submit">Submit</button>
      </form>
    </StyledLoginForm>
  );
}

export default LoginForm;
