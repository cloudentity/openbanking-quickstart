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
  authorised: "#4CAF50",
  inactive: "#BD271E",
  revoked: "#BD271E",
  pending: "#BD271E",
  other: "#BD271E",
};

type Props = {
  children: ReactNode;
  type?: keyof typeof colorMapper;
  id?: string;
};

function Chip({ children, type, id }: Props) {
  const classes = useStyles();

  return type ? (
    <div
      className={classes.container}
      id={id}
      style={{
        color: colorMapper[type] || colorMapper.other,
        borderColor: colorMapper[type] || colorMapper.other,
      }}
    >
      {children}
    </div>
  ) : null;
}

export default Chip;
