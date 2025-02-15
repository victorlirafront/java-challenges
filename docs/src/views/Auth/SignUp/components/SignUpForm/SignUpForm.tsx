import { useState } from 'react';
import axios from 'axios';
import { useRouter } from 'next/router';
import { StyledSignUpForm } from './SignUpForm.styled';

function SignUpForm() {
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const [confirmPassword, setConfirmPassword] = useState('');
  const [usernameError, setUsernameError] = useState('');
  const [passwordError, setPasswordError] = useState('');
  const [confirmPasswordError, setConfirmPasswordError] = useState('');
  const [error, setError] = useState('');
  const router = useRouter();

  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    let hasError = false;

    setUsernameError('');
    setPasswordError('');
    setConfirmPasswordError('');

    if (username.length < 8) {
      setUsernameError('Username must be at least 8 characters long.');
      hasError = true;
    }

    if (password.length < 8) {
      setPasswordError('Password must be at least 8 characters long.');
      hasError = true;
    }

    if (password !== confirmPassword) {
      setConfirmPasswordError('Passwords do not match.');
      hasError = true;
    }

    if (hasError) return;

    try {
      const formData = new FormData();
      formData.append('username', username);
      formData.append('password', password);

      const response = await axios.post('http://localhost:8080/register', formData, {
        headers: {
          'Content-Type': 'multipart/form-data',
        },
        withCredentials: true,
      });

      console.log('Signup successful:', response.data);
      setError('');
      router.push('/auth/login');
    } catch (error) {
      if (axios.isAxiosError(error)) {
        setError(error.response?.data.error || 'An error occurred during sign-up.');
      } else {
        setError('An unexpected error occurred.');
      }
      console.error('Error during signup:', error);
    }
  };

  return (
    <StyledSignUpForm>
      <form onSubmit={handleSubmit}>
        <h1>Sign up</h1>

        <div className="form-control">
          <label htmlFor="username">Username</label>
          <input
            type="text"
            id="username"
            placeholder="Enter a unique username"
            value={username}
            onChange={(e) => setUsername(e.target.value)}
          />
          {usernameError && <p className="error">{usernameError}</p>}
        </div>

        <div className="form-control">
          <label htmlFor="password">Password</label>
          <input
            type="password"
            id="password"
            placeholder="Create a strong password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
          />
          {passwordError && <p className="error">{passwordError}</p>}
        </div>

        <div className="form-control">
          <label htmlFor="confirm-password">Repeat Password</label>
          <input
            type="password"
            id="confirm-password"
            placeholder="Confirm your password"
            value={confirmPassword}
            onChange={(e) => setConfirmPassword(e.target.value)}
          />
          {confirmPasswordError && <p className="error">{confirmPasswordError}</p>}
        </div>

        {error && <p className="error">{error}</p>}

        <button type="submit">Sign Up</button>
      </form>
    </StyledSignUpForm>
  );
}

export default SignUpForm;
