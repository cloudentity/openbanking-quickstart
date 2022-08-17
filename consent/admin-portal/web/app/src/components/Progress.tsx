import React from "react";
import { makeStyles } from "@material-ui/core/styles";
import CircularProgress from "@material-ui/core/CircularProgress";

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

const Progress = ({ size = 100, top = 40 }) => {
  const classes = useStyles();

  return (
    <div className={classes.progress} style={{ top }}>
      <CircularProgress size={size} className={classes.circle} thickness={3} />
    </div>
  );
};

export default Progress;
