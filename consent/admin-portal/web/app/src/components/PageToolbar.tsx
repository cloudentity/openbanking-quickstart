import React from "react";
import { useHistory, useLocation } from "react-router";
import Toolbar from "@material-ui/core/Toolbar";
import AppBar from "@material-ui/core/AppBar";
import Button from "@material-ui/core/Button";
import { Theme } from "@material-ui/core";
import ExitToAppIcon from "@material-ui/icons/ExitToApp";
import { makeStyles } from "@material-ui/core/styles";
import Tabs from "@material-ui/core/Tabs";
import Tab from "@material-ui/core/Tab";

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
    ...theme.custom.button,
    color: "#DC1B37",
  },
  appBar: {
    backgroundColor: "white",
    minHeight: 64,
    padding: "0 80px",
    [theme.breakpoints.down("sm")]: {
      padding: 0,
    },
  },
  tabsContainer: {
    height: 64,
    position: "relative",
    top: 1,
  },
  tabsRoot: {
    textTransform: "none",
    color: "#002D4C",
    height: "100%",
    marginRight: 16,
    "& .MuiTabs-flexContainer": {
      height: "100%",
    },
  },
  tabsIndicator: {
    backgroundColor: "#DC1B37",
  },
  tab: {
    ...theme.custom.body2,
    textTransform: "none",
  },
  content: {
    color: "black",
  },
  tabRoot: {
    minWidth: "unset",
  },
}));

export default function PageToolbar({
  authorizationServerURL,
  authorizationServerId,
  tenantId,
}) {
  const classes = useStyles();
  const { pathname } = useLocation();
  const history = useHistory();
  const activeTab =
    ((pathname === "/" || pathname.includes("/accounts/")) && "/") ||
    (pathname === "/providers" && "/providers") ||
    null;

  return (
    <AppBar
      position="fixed"
      variant="outlined"
      classes={{
        root: classes.appBar,
      }}
    >
      <Toolbar>
        <img src={logo} alt="logo" style={{ width: 160 }} />
        <div style={{ flex: 1 }} />
        <div className={classes.tabsContainer}>
          <Tabs
            value={activeTab}
            onChange={(_, newValue) => history.push(newValue)}
            classes={{
              root: classes.tabsRoot,
              indicator: classes.tabsIndicator,
            }}
          >
            <Tab
              id="consent-management-tab"
              key="consent"
              value="/"
              label="Consent management"
              className={classes.tab}
              classes={{
                root: classes.tabRoot,
              }}
              style={activeTab !== "/" ? { opacity: 0.5 } : {}}
            />
            <Tab
              id="third-party-providers-tab"
              key="access"
              value="/providers"
              label="Third party providers"
              className={classes.tab}
              classes={{
                root: classes.tabRoot,
              }}
              style={activeTab !== "/providers" ? { opacity: 0.5 } : {}}
            />
          </Tabs>
        </div>
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
