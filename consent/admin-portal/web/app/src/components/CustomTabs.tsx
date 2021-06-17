import React, { ReactNode } from "react";
import { makeStyles } from "@material-ui/core/styles";
import { Theme } from "@material-ui/core";

const useStyles = makeStyles((theme: Theme) => ({
  container: {
    background: "#FFFFFF",
    boxShadow:
      "0px 1px 16px -4px rgba(0, 0, 0, 0.25), 0px 0px 1px rgba(0, 0, 0, 0.31)",
    borderRadius: 4,
    maxWidth: 850,
    margin: "0 auto 24px auto",
    boxSizing: "border-box",
    padding: "10px 24px 0 24px",
  },
  tabsRoot: {
    textTransform: "none",
    color: "#626576",
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

type Props = {
  tabs: {
    key: string;
    label: string;
    content: ReactNode;
  }[];
};

function CustomTabs({ tabs }: Props) {
  const classes = useStyles();
  const curentTab = tabs[0];

  return (
    <div className={classes.container}>
      <div className={classes.content}>{curentTab?.content}</div>
    </div>
  );
}

export default CustomTabs;
