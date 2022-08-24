import React from "react";
import { useLocation, useNavigate } from "react-router-dom";
import Toolbar from "@mui/material/Toolbar";
import AppBar from "@mui/material/AppBar";
import Button from "@mui/material/Button";
import ExitToAppIcon from "@mui/icons-material/ExitToApp";
import { makeStyles } from "tss-react/mui";
import Tabs from "@mui/material/Tabs";
import Tab from "@mui/material/Tab";

import logo from "../assets/gobank-logo.svg";
import { logout } from "./AuthPage";

const useStyles = makeStyles()(theme => ({
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
    color: "#DC1B37",
  },
  appBar: {
    backgroundColor: "white",
    minHeight: 64,
    padding: "0 80px",
    [theme.breakpoints.down("md")]: {
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
  const { classes } = useStyles();
  const { pathname } = useLocation();
  const navigate = useNavigate();
  const activeTab =
    ((pathname === "/" || pathname.includes("/accounts/")) && "/") ||
    (pathname === "/providers" && "/providers") ||
    null;

  return (
    <AppBar
      position="fixed"
      variant="outlined"
      className={classes.appBar}
      elevation={0}
    >
      <Toolbar>
        <img src={logo} alt="logo" style={{ width: 160 }} />
        <div style={{ flex: 1 }} />
        <div className={classes.tabsContainer}>
          <Tabs
            value={activeTab}
            onChange={(_, newValue) => navigate(newValue)}
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
        <Button className={classes.signOutButton} onClick={logout}>
          <ExitToAppIcon style={{ marginRight: 10 }} />
          Sign out
        </Button>
      </Toolbar>
    </AppBar>
  );
}
