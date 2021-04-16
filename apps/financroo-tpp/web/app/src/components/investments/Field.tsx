import React from "react";
import { makeStyles } from "@material-ui/core/styles";

const useStyles = makeStyles((theme) => ({
  container: {
    width: "100%",
    marginBottom: 24,
  },
  label: {
    fontWeight: 600,
    fontSize: 12,
    lineHeight: "24px",
    color: "#212533",
    marginBottom: 4,
  },
  helperText: {
    fontSize: 12,
    lineHeight: "22px",
    color: "#626576",
    marginTop: 8,
  },
}));

type Props = {
  label?: string | React.ReactNode;
  helperText?: string;
  children: React.ReactNode;
  style?: React.CSSProperties;
};

export default function Field({
  label,
  helperText,
  children,
  style = {},
}: Props) {
  const classes = useStyles();

  return (
    <div className={classes.container} style={style}>
      {label && <div className={classes.label}>{label}</div>}
      {children}
      {helperText && <div className={classes.helperText}>{helperText}</div>}
    </div>
  );
}
