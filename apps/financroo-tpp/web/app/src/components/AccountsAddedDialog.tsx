import React from "react";
import IconButton from "@material-ui/core/IconButton";
import CloseIcon from "@material-ui/icons/Close";
import { makeStyles } from "@material-ui/core/styles";
import Dialog from "@material-ui/core/Dialog";
import financrooLogo from "../assets/financroo-logo.svg";
import icon from "../assets/icon-check.svg";
import { Button } from "@material-ui/core";
import { useHistory } from "react-router";

const useStyles = makeStyles((theme) => ({
  dialog: {
    width: 454,
    borderRadius: 0,
  },
  dialogContent: {
    backgroundColor: "#F7FAFF",
    padding: "48px 80px",
    display: "flex",
    flexDirection: "column",
    alignItems: "center",
    justifyContent: "center",
    "& > div": {
      width: 280,
      textAlign: "center",
      marginTop: 16,
      fontSize: 16,
      lineHeight: "24px",
    },
    "& > img": {
      width: 68,
    },
  },
  closeButton: {
    position: "absolute",
    top: 0,
    right: 0,
  },
  dialogHeader: {
    display: "flex",
    alignItems: "center",
    justifyContent: "center",
    position: "relative",
    "& > img": {
      width: 160,
      padding: "24px 0",
    },
  },
  dialogButtons: {
    padding: "32px 65px",
    display: "flex",
    justifyContent: "center",
    "& button": {
      textTransform: "none",
      ...theme.custom.button,
      borderColor: theme.palette.primary.main,
      "&:first-of-type": {
        color: theme.palette.primary.main,
      },
      "&:last-of-type": {
        backgroundColor: theme.palette.primary.main,
      },
    },
  },
}));

export default function AcccountsAddedDialog({ open, setOpen }) {
  const classes = useStyles();
  const history = useHistory();

  return (
    <Dialog onClose={() => setOpen(false)} open={open}>
      <div className={classes.dialogHeader}>
        <div className={classes.closeButton}>
          <IconButton onClick={() => setOpen(false)}>
            <CloseIcon />
          </IconButton>
        </div>
        <img src={financrooLogo} alt="financroo logo" />
      </div>
      <div className={classes.dialogContent}>
        <img src={icon} alt="icon" />
        <div>
          Your <strong>Go Bank</strong> account(s) has been successfully
          connected to Financroo
        </div>
        <div>
          Now you can use Financroo <br />
          to make investments!
        </div>
      </div>
      <div className={classes.dialogButtons}>
        <Button
          variant="outlined"
          style={{ marginRight: 16 }}
          onClick={() => setOpen(false)}
        >
          Cancel
        </Button>
        <Button
          variant="outlined"
          onClick={() => {
            history.push("/investments");
          }}
        >
          Start investing
        </Button>
      </div>
    </Dialog>
  );
}
