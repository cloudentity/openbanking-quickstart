import React, { ReactNode } from "react";
import { makeStyles } from "@material-ui/core/styles";

const useStyles = makeStyles(() => ({
  container: {
    border: "1px solid #4CAF50",
    boxSizing: "border-box",
    borderRadius: 3,
    padding: "0 4px",
    lineHeight: "22px",
    textTransform: "capitalize",
    display: "inline-block",
    fontWeight: 400,
  },
}));

const colorMapper = {
  active: "#4CAF50",
  inactive: "#BD271E",
};

type Props = {
  children: ReactNode;
  type: keyof typeof colorMapper;
};

function Chip({ children, type }: Props) {
  const classes = useStyles();

  return (
    <div
      className={classes.container}
      style={{ color: colorMapper[type], borderColor: colorMapper[type] }}
    >
      {children}
    </div>
  );
}

export default Chip;
