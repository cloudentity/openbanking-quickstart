import React from "react";
import { Redirect } from "react-router";
import {Theme} from "@material-ui/core";
import {makeStyles} from "@material-ui/core/styles";
import backgroundLogin from "../assets/background-login.png";
import financrooLogo from "../assets/financroo-logo.svg";
import Grid from "@material-ui/core/Grid";
import Button from "@material-ui/core/Button";
import { ThemeProvider, createMuiTheme } from "@material-ui/core/styles";
import { Alert } from '@material-ui/lab';
import TextField from "@material-ui/core/TextField";

import {
  getTokenFromStore,
  isTokenInStore,
  removeAllAuthDataFromStore,
} from "./auth.utils";


const useStyles = makeStyles(() => ({
  container: {
    background: "#FFFFFF",
    boxShadow:
      "0px 1px 1px rgba(0, 0, 0, 0.08), 0px 0px 1px rgba(0, 0, 0, 0.31)",
    borderRadius: 4,
    maxWidth: 650,
    margin: "0 auto 24px auto",
    boxSizing: "border-box",
    padding: 48,
    marginTop: 24,
  },
  header: {
    borderBottom: "1px solid #ECECEC",
    padding: "0 32px",
  },
}));

export const logout = () => {
  removeAllAuthDataFromStore();
  window.location.href = `/`;
};


const AuthPage = ({
  login
}) => {

  const [state, setState] = React.useState({
    login: 'admin',
    password: 'p@ssw0rd!',
    error: false,
    processing: false,
  })

  const classes = useStyles();

  if (isTokenInStore()) {
    login({ token: getTokenFromStore() });
    return <Redirect to={"/"} />;
  }

  const onSubmit = (e) => {
    e.preventDefault();
    setState({...state, processing: true})

    if (state.login === "admin" && state.password === "p@ssw0rd!") {
      login({token: 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.XbPfbIHMI6arZ3Y922BhjWgQzWXcXNrz0ogtVhfEd2o'})
      return
    }

    setState({...state, processing:false, password: '', error: true})
  }

  const handleLoginChange = (e) => {
    setState({...state, login: e.target.value});
  }

  const handlePasswordChange = (e) =>{
    setState({...state, password: e.target.value});
  }


  return (
    <div style={{
      position:"fixed",
      bottom: "0px",
      top: "0px",
      left: "0px",
      right: "0px",
      backgroundColor: "#002D4C",

      backgroundImage: "url(\"/static/media/background.f7f97194.svg\")",
      backgroundPosition: "left center",
      backgroundRepeat: "no-repeat",
      backgroundSize: "contain"
    }}>
      <div className={classes.container}>
        <center>
          <img src="/static/media/gobank-logo.042fc889.svg" alt="logo" />
        </center>

        <form onSubmit={onSubmit}>
          <TextField margin="normal"
            id="standard-login-input"
            label="Login"
            type="text"
            autoComplete="current-login"
            color="primary"
            value={state.login}
            onChange={handleLoginChange}
            style={{width: '100%'}}
          />
          <TextField margin="normal"
            id="standard-password-input"
            label="Password"
            type="password"
            autoComplete="current-password"
            color="primary"
            value={state.password}
            onChange={handlePasswordChange}
            style={{width: '100%', color: 'white',}}
          />

          <Button data-testid="login-button" disabled={state.processing} id="login-button" type="submit" className={'login-button'}  style={{width: '100%', minHeight: 50, marginTop: 24,}} variant={'contained'}>Login</Button>
          <Alert severity="error" variant="outlined" style={{marginTop: 24, display: state.error? '': 'none',}}>Invalid login or password</Alert>
        </form>
      </div>
    </div>
  )
};

export default AuthPage;
