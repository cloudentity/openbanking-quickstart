import React, { ReactNode, CSSProperties } from "react";
import { makeStyles } from "tss-react/mui";

const useStyles = makeStyles()(theme => ({
  container: {
    backgroundColor: "#002D4C",
    padding: "48px 0",
    boxSizing: "border-box",
    textAlign: "center",
  },
  content: {
    maxWidth: 588,
    margin: "auto",
    color: "white",
  },
  title: {
    ...theme.custom.heading2,
    marginBottom: 16,
  },
  info: {
    marginTop: 12,
    ...theme.custom.body2,
  },
  icon: {
    marginBottom: 15,
  },
}));

interface Props {
  title: string | ReactNode;
  children?: ReactNode;
  containerStyle?: CSSProperties;
  contentStyle?: CSSProperties;
  icon?: string;
}

function Subheader({
  title,
  children,
  containerStyle,
  contentStyle,
  icon,
}: Props) {
  const { classes } = useStyles();

  return (
    <div className={classes.container} style={containerStyle}>
      <div className={classes.content} style={contentStyle}>
        {icon && <img src={icon} alt="icon" className={classes.icon} />}
        <div className={classes.title}>{title}</div>
        {children && <div className={classes.info}>{children}</div>}
      </div>
    </div>
  );
}

export default Subheader;
