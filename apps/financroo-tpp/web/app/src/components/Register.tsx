import React from "react";
import { makeStyles } from "tss-react/mui";
import backgroundLogin from "../assets/background-login.png";
import financrooLogo from "../assets/financroo-logo.svg";
import Grid from "@mui/material/Grid";
import Button from "@mui/material/Button";

import { Alert } from "@mui/material";
import TextField from "@mui/material/TextField";

const useStyles = makeStyles()(theme => ({
  root: {
    height: "100vh",
  },
  image: {
    width: "100%",
    height: "100%",
    backgroundImage: `url(${backgroundLogin})`,
    backgroundRepeat: "no-repeat",
    backgroundPosition: "center",
    backgroundSize: "cover",
  },
  formContainerRoot: {
    height: "100%",
    display: "flex",
    flexDirection: "column",
    alignItems: "flex-start",
    justifyContent: "center",

    [theme.breakpoints.down("md")]: {
      padding: 16,
    },
    [theme.breakpoints.up("sm")]: {
      padding: 48,
    },
    [theme.breakpoints.up("lg")]: {
      padding: 82,
    },
  },
}));

export default function Register({ onLogin }) {
  const { classes } = useStyles();

  const [state, setState] = React.useState({
    login: "test",
    password: "p@ssw0rd!",
    error: false,
    processing: false,
  });;

  const loginWrapper = fn => {
    return async e => {
      e.preventDefault();
      setState({ ...state, processing: true });

      try {
        await fn(state.login, state.password);
      } catch (e) {
        setState({ ...state, processing: false, password: "", error: true });
        return;
      }

      setState({ ...state, processing: false, password: "", error: false });
    };
  };

  const handleLoginChange = e => {
    setState({ ...state, login: e.target.value });
  };
  const handleLoginChange = e => {
    setState({ ...state, login: e.target.value });
  };

  const handlePasswordChange = e => {
    setState({ ...state, password: e.target.value });
  };
  const handlePasswordChange = e => {
    setState({ ...state, password: e.target.value });
  };

  return (
    <div className={classes.root}>
      <Grid container style={{ height: "100%" }}>
        <Grid item sm={6} lg={7}>
          <div className={classes.image}  />
        </Grid>
        <Grid item xs={12} sm={6} lg={5}>
          <div className={classes.formContainerRoot}>
            <img
              alt="financroo logo"
              src={financrooLogo}
              style={{ marginBottom: 44 }}
            />
            <form onSubmit={loginWrapper(onLogin)}>
              <TextField
               
                margin="normal"
                id="standard-login-input"
                label="Login"
                type="text"
                autoComplete="current-login"
                value={state.login}
                onChange={handleLoginChange}
                style={{ width: "100%" }}
              />
              <TextField
               
                margin="normal"
                id="standard-password-input"
                label="Password"
                type="password"
                autoComplete="current-password"
                value={state.password}
                onChange={handlePasswordChange}
                style={{ width: "100%" }}
              />

              <Button
                data-testid="login-button"
                disabled={state.processing}
                id="login-button"
                type="submit"
                className="login-button"
                color="secondary"
                style={{ width: "100%", minHeight: 50, marginTop: 24 }}
                variant="contained"
              >
                Login
              </Button>
              <Alert
                severity="error"
                variant="outlined"
                style={{
                  width: "100%",
                  marginTop: 24,
                  visibility: state.error ? "visible" : "hidden",
                }}
              >
                Invalid login or password
              </Alert>
            </form>
          </div>
        </Grid>
      </Grid>
    </div>
  );
}
