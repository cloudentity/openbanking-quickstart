import React, { ReactNode } from "react";
import { makeStyles } from "tss-react/mui";
import Drawer from "@mui/material/Drawer";
import { Consent } from "../types";

const useStyles = makeStyles()(() => ({
  container: {
    width: 500,
    marginBottom: 84,
    overflow: "auto",
  },
  header: {
    display: "flex",
    alignItems: "center",
    borderBottom: "1px solid #ECECEC",
  },
  content: {
    padding: 32,
  },
  bottomBar: {
    position: "absolute",
    bottom: 0,
    boxShadow:
      "0px 6px 10px rgba(0, 0, 0, 0.14), 0px 1px 18px rgba(0, 0, 0, 0.12), 0px 3px 5px rgba(0, 0, 0, 0.2)",
    width: "100%",
    padding: 24,
    display: "flex",
    justifyContent: "space-between",
    boxSizing: "border-box",
  },
}));

interface Props {
  children: ReactNode;
  header: ReactNode;
  bottomBar?: ReactNode;
  setDrawerData: (data: Consent | undefined) => void;
}

function ApplicationAccessDrawer({
  children,
  header,
  bottomBar,
  setDrawerData,
}: Props) {
  const { classes } = useStyles();

  return (
    <Drawer anchor="right" open={true} onClose={() => setDrawerData(undefined)}>
      <div className={classes.header}>{header}</div>
      <div className={classes.container}>
        <div className={classes.content}>{children}</div>
      </div>
      {bottomBar && <div className={classes.bottomBar}>{bottomBar}</div>}
    </Drawer>
  );
}

export default ApplicationAccessDrawer;
