import { StyledLoginForm } from "./LoginForm.styled";

function LoginForm() {
  return (
    <StyledLoginForm>
      <form action="">
        <h1>Login</h1>
        <div className="form-control">
          <label htmlFor="username">Username</label>
          <input type="text" id="username" placeholder="Enter your username" />
        </div>
        <div className="form-control">
          <label htmlFor="password">Password</label>
          <input type="password" id="password" placeholder="Enter your password" />
        </div>

        <button>Submit</button>
      </form>
    </StyledLoginForm>
  );
}

export default LoginForm;
