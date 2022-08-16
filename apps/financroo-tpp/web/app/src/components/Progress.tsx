import React from "react";
import CircularProgress from "@material-ui/core/CircularProgress";
import { makeStyles } from "@material-ui/core/styles";

const useStyles = makeStyles(() => ({
  progress: {
    width: 100,
    height: 100,
    position: "absolute",
    left: "calc(50% - 50px);",
  },
  circle: {
    position: "absolute",
    top: 0,
    left: 0,
  },
}));

interface Props {
  size?: number;
  top?: number;
}

const Progress = ({ size = 100, top = 40 }: Props) => {
  const classes = useStyles();

  return (
    <div className={classes.progress} style={{ top }}>
      <CircularProgress size={size} className={classes.circle} thickness={3} />
    </div>
  );
};

export default Progress;
