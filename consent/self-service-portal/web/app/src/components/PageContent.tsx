import React from "react";
import {Theme} from "@material-ui/core";
import {makeStyles} from "@material-ui/core/styles";
import brandingMask from "../assets/branding-mask.png";

const useStyles = withBackground => makeStyles((theme: Theme) => ({
  root: {
    marginTop: 56,
    height: 'calc(100vh - 56px)',
    [theme.breakpoints.up('sm')]: {
      marginTop: 64,
      height: 'calc(100vh - 64px)',
    },
    overflowY: 'auto',
    backgroundColor: '#F7FAFF',
    position: 'fixed',
    width: '100vw',
    ...withBackground
      ? {
        [theme.breakpoints.up('lg')]: {
          backgroundImage: `url(${brandingMask})`,
          backgroundPosition: 'left top',
          backgroundRepeat: 'no-repeat',
        },
      }
      : {}
  }
}));

export default function PageContent ({children, withBackground = false, style = {}}) {
  const classes = useStyles(withBackground)();

  return (
    <div className={classes.root} style={style}>
      {children}
    </div>
  )
};
