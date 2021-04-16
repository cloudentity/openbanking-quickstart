import React from "react";
import { Theme } from "@material-ui/core";
import { makeStyles } from "@material-ui/core/styles";
import brandingMask from "../../assets/branding-mask.png";
import { subHeaderHeight } from "./PageToolbar";

const useStyles = (
  withBackground: boolean,
  withSubheader: boolean,
  withOnlySubheader: boolean
) => {
  const height =
    (withSubheader && 64 + subHeaderHeight) ||
    (withOnlySubheader && subHeaderHeight) ||
    64;

  return makeStyles((theme: Theme) => ({
    root: {
      marginTop: 56,
      height: `calc(100vh - 56px)`,
      [theme.breakpoints.up("sm")]: {
        marginTop: height,
        height: `calc(100vh - ${height}px)`,
      },
      overflowY: "auto",
      backgroundColor: "#F7FAFF",
      position: "fixed",
      minWidth: "100vw",
      ...(withBackground
        ? {
            [theme.breakpoints.up("lg")]: {
              backgroundImage: `url(${brandingMask})`,
              backgroundPosition: "left top",
              backgroundRepeat: "no-repeat",
            },
          }
        : {}),
    },
  }));
};

export default function PageContent({
  children,
  withBackground = false,
  withSubheader = false,
  withOnlySubheader = false,
  style = {},
}) {
  const classes = useStyles(withBackground, withSubheader, withOnlySubheader)();

  return (
    <div className={classes.root} style={style}>
      {children}
    </div>
  );
}
