import React from "react";
import Toolbar from "@mui/material/Toolbar";
import AppBar from "@mui/material/AppBar";
import Button from "@mui/material/Button";
import ExitToAppIcon from "@mui/icons-material/ExitToApp";
import { makeStyles } from "tss-react/mui";

import logo from "../assets/gobank-logo.svg";
import { logout } from "./AuthPage";

const useStyles = makeStyles()(() => ({
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
  const { classes } = useStyles();

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
