import React, { ReactNode } from "react";
import { makeStyles } from "@material-ui/core/styles";

const useStyles = makeStyles(theme => ({
  container: {
    backgroundColor: "#002D4C",
    padding: "30px 0",
    boxSizing: "border-box",
  },
  content: {
    maxWidth: 850,
    margin: "auto",
    color: "white",
  },
  title: {
    //  ...theme.custom.heading2,
  },
  info: {
    marginTop: 12,
    ...theme.custom.body2,
  },
}));

type Props = {
  title: string | ReactNode;
  children?: ReactNode;
};

function Subheader({ title, children }: Props) {
  const classes = useStyles();

  return (
    <div className={classes.container}>
      <div className={classes.content}>
        <div className={classes.title}>{title}</div>
        {children && <div className={classes.info}>{children}</div>}
      </div>
    </div>
  );
}

export default Subheader;
