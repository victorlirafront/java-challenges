import { StyledSignUpForm } from './SignUpForm.styled';

function SignUpForm() {
  return (
    <StyledSignUpForm>
      <form action="">
        <h1>Sign up</h1>
        <div className="form-control">
          <label htmlFor="username">Username</label>
          <input type="text" id="username" placeholder="Enter a unique username" />
        </div>
        <div className="form-control">
          <label htmlFor="password">Password</label>
          <input type="password" id="password" placeholder="Create a strong password" />
        </div>

        <div className="form-control">
          <label htmlFor="confirm-password">Repeat Password</label>
          <input type="password" id="confirm-password" placeholder="Confirm your password" />
        </div>

        <button>Submit</button>
      </form>
    </StyledSignUpForm>
  );
}

export default SignUpForm;
