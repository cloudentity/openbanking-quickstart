import React from "react";
import Toolbar from "@material-ui/core/Toolbar";
import AppBar from "@material-ui/core/AppBar";
import Button from "@material-ui/core/Button";
import { Theme } from "@material-ui/core";
import ExitToAppIcon from "@material-ui/icons/ExitToApp";
import { makeStyles } from "@material-ui/core/styles";

import logo from "../assets/gobank-logo.svg";
import { logout } from "./AuthPage";

const useStyles = makeStyles((theme: Theme) => ({
  indicator: {
    backgroundColor: "#fff",
  },
  expandIcon: {
    position: "absolute",
    right: 32,
    top: 24,
    color: "#006580",
  },
  signOutButton: {
 //   ...theme.custom.button,
    color: "#DC1B37",
  },
}));

export default function PageToolbar({
  authorizationServerURL,
  authorizationServerId,
  tenantId,
}) {
  const classes = useStyles();

  return (
    <AppBar
      position="fixed"
      variant="outlined"
      style={{ backgroundColor: "white" }}
    >
      <Toolbar>
        <img src={logo} alt="logo" style={{ width: 160 }} />
        <div style={{ flex: 1 }} />
        <Button
          className={classes.signOutButton}
          onClick={() =>
            logout(authorizationServerURL, tenantId, authorizationServerId)
          }
        >
          <ExitToAppIcon style={{ marginRight: 10 }} />
          Sign out
        </Button>
      </Toolbar>
    </AppBar>
  );
}
