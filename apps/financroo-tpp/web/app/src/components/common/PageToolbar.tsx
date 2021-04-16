import React from "react";
import { useHistory } from "react-router";
import Toolbar from "@material-ui/core/Toolbar";
import AppBar from "@material-ui/core/AppBar";
import MenuIcon from "@material-ui/icons/Menu";
import IconButton from "@material-ui/core/IconButton";
import Tab from "@material-ui/core/Tab";
import Hidden from "@material-ui/core/Hidden";
import Tabs from "@material-ui/core/Tabs";
import Button from "@material-ui/core/Button";
import { makeStyles } from "@material-ui/core/styles";

import financrooLogo from "../../assets/financroo-logo.svg";
import { logout } from "../AuthPage";

export const subHeaderHeight = 116;

const useStyles = (withSubheader: boolean, mode: string) =>
  makeStyles((theme) => ({
    appBar: {
      ...(withSubheader
        ? {
            border: "none",
          }
        : {}),
    },
    toolBar: {
      ...(withSubheader
        ? {
            border: "1px solid transparent",
            borderBottom: "none",
          }
        : {}),
      ...(mode === "onlySubheader"
        ? {
            display: "none",
          }
        : {}),
    },
    subHeaderContainer: {
      height: subHeaderHeight,
      backgroundColor: theme.palette.secondary.main,
      color: theme.palette.primary.main,
      fontSize: 28,
      lineHeight: "40px",
      padding: "0 80px",
      display: "flex",
      alignItems: "center",
      justifyContent: "space-between",
    },
    button: {
      color: "white",
      fontSize: 16,
      lineHeight: "24px",
      textTransform: "none",
      padding: "8px 24px",
      "&:hover": {
        backgroundColor: theme.palette.primary.main,
      },
    },
  }));

type CommonProps = {
  tab?: "accounts" | "investments";
  subHeaderTitle?: string | React.ReactNode;
  subHeaderButton?: { title: string; onClick: () => void };
};

type Props =
  | ({
      mode: "dialog";
      children: React.ReactNode;
      authorizationServerURL?: string;
      authorizationServerId?: string;
      tenantId?: string;
    } & CommonProps)
  | ({
      mode: "main";
      children?: React.ReactNode;
      authorizationServerURL: string;
      authorizationServerId: string;
      tenantId: string;
    } & CommonProps)
  | ({
      mode: "onlySubheader";
      children?: React.ReactNode;
      authorizationServerURL?: string;
      authorizationServerId?: string;
      tenantId?: string;
    } & CommonProps);

export default function PageToolbar({
  mode,
  children,
  authorizationServerURL,
  authorizationServerId,
  tenantId,
  tab,
  subHeaderTitle,
  subHeaderButton,
}: Props) {
  const history = useHistory();
  const classes = useStyles(!!subHeaderTitle, mode)();

  return (
    <AppBar
      position="fixed"
      color="inherit"
      variant="outlined"
      className={classes.appBar}
    >
      <Toolbar className={classes.toolBar}>
        <img alt="financroo logo" src={financrooLogo} />
        <div style={{ flex: 1 }} />

        {mode === "dialog" && children}
        {mode === "main" && (
          <>
            <Hidden mdUp>
              <IconButton edge="start" color="inherit" aria-label="menu">
                <MenuIcon />
              </IconButton>
            </Hidden>
            <Hidden smDown>
              <Tabs
                value={tab || "accounts"}
                indicatorColor="primary"
                aria-label="menu tabs"
                style={{ height: 64 }}
              >
                <Tab
                  label="Accounts"
                  value="accounts"
                  style={{ height: 64 }}
                  onClick={() => history.push("/")}
                />
                <Tab
                  label="Investments"
                  value="investments"
                  style={{ height: 64 }}
                  onClick={() => history.push("/investments")}
                />
                <Tab label="Spending" value="spending" style={{ height: 64 }} />
                <Tab label="Settings" value="settings" style={{ height: 64 }} />
              </Tabs>
            </Hidden>
            <Button
              variant="outlined"
              onClick={() =>
                logout(authorizationServerURL, tenantId, authorizationServerId)
              }
            >
              Logout
            </Button>
          </>
        )}
      </Toolbar>
      {subHeaderTitle && (
        <div className={classes.subHeaderContainer}>
          <div>{subHeaderTitle}</div>
          {subHeaderButton && (
            <Button
              onClick={subHeaderButton.onClick}
              variant="contained"
              color="primary"
              className={classes.button}
            >
              {subHeaderButton.title}
            </Button>
          )}
        </div>
      )}
    </AppBar>
  );
}
