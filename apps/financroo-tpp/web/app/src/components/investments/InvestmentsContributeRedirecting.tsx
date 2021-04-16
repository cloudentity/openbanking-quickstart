import React, { useEffect } from "react";
import { makeStyles } from "@material-ui/core/styles";
import LinearProgress from "@material-ui/core/LinearProgress";

import iconBank from "../../assets/icon-bank2.svg";

const useStyles = makeStyles(() => ({
  container: {
    position: "fixed",
    top: 0,
    left: 0,
    width: "100vw",
    height: "100vh",
    overflow: "hidden",
    backgroundColor: "#F7FAFF",
    display: "flex",
    alignItems: "center",
    justifyContent: "center",
    flexDirection: "column",
    textAlign: "center",
    zIndex: 10000,
  },
  progress: {
    width: 260,
    marginTop: 16,
  },
  text: {
    marginTop: 34,
    fontWeight: 600,
    fontSize: 16,
    lineHeight: "34px",
  },
}));

type Props = {
  handleNext: () => void;
};
export default function InvestmentsContributeRedirecting({
  handleNext,
}: Props) {
  const classes = useStyles();

  useEffect(() => {
    setTimeout(() => {
      handleNext();
    }, 2000);
  }, [handleNext]);

  return (
    <div className={classes.container}>
      <img src={iconBank} alt="icon" />
      <div className={classes.progress}>
        <LinearProgress />
      </div>
      <div className={classes.text}>
        We are redirecting you to Go Bank
        <br />
        <br />
        You will be redirected back here after you <br />
        confirm your transaction
      </div>
    </div>
  );
}
