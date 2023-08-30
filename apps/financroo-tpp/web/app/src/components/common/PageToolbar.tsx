import React from "react";
import Toolbar from "@mui/material/Toolbar";
import AppBar from "@mui/material/AppBar";
import MenuIcon from "@mui/icons-material/Menu";
import IconButton from "@mui/material/IconButton";
import Tab from "@mui/material/Tab";
import Hidden from "@mui/material/Hidden";
import Tabs from "@mui/material/Tabs";
import Button from "@mui/material/Button";
import { makeStyles } from "tss-react/mui";

import financrooLogo from "../../assets/financroo-logo.svg";
import { logout } from "../AuthPage";
import { useNavigate } from "react-router-dom";

export const subHeaderHeight = 116;

const useStyles = (withSubheader: boolean, mode: string) =>
  makeStyles()(theme => ({
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
  subHeaderButton?: { title: string; onClick: () => void; id?: string };
};

type Props =
  | ({
      mode: "dialog";
      children: React.ReactNode;
    } & CommonProps)
  | ({
      mode: "main";
      children?: React.ReactNode;
    } & CommonProps)
  | ({
      mode: "onlySubheader";
      children?: React.ReactNode;
    } & CommonProps);

export default function PageToolbar({
  mode,
  children,
  tab,
  subHeaderTitle,
  subHeaderButton,
}: Props) {
  const navigate = useNavigate();
  const { classes } = useStyles(!!subHeaderTitle, mode)();

  return (
    <AppBar
      position="fixed"
      color="inherit"
      variant="outlined"
      className={classes.appBar}
      elevation={0}
    >
      <Toolbar className={classes.toolBar}>
        <img alt="financroo logo" src={financrooLogo} />
        <div style={{ flex: 1 }} />

        {mode === "dialog" && children}
        {mode === "main" && (
          <>
            <Hidden mdUp>
              <IconButton
                edge="start"
                color="inherit"
                aria-label="menu"
                size="large"
              >
                <MenuIcon />
              </IconButton>
            </Hidden>
            <Hidden mdDown>
              <Tabs
                value={tab || "accounts"}
                indicatorColor="primary"
                aria-label="menu tabs"
                style={{ height: 64 }}
              >
                <Tab
                  label="Accounts"
                  value="accounts"
                  id="accounts-tab"
                  style={{ height: 64 }}
                  onClick={() => navigate("/")}
                />
                {window.featureFlags?.Investments && (
                <Tab
                  label="Investments"
                  value="investments"
                  id="investments-tab"
                  style={{ height: 64 }}
                  onClick={() => navigate("/investments")}
                />
                )}
                <Tab label="Spending" value="spending" style={{ height: 64 }} />
                <Tab label="Settings" value="settings" style={{ height: 64 }} />
              </Tabs>
            </Hidden>
            <Button variant="outlined" onClick={logout}>
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
              id={subHeaderButton.id}
            >
              {subHeaderButton.title}
            </Button>
          )}
        </div>
      )}
    </AppBar>
  );
}
