import React from "react";
import IconButton from "@mui/material/IconButton";
import CloseIcon from "@mui/icons-material/Close";
import { makeStyles } from "tss-react/mui";
import Dialog from "@mui/material/Dialog";
import financrooLogo from "../assets/financroo-logo.svg";
import icon from "../assets/icon-check.svg";
import Button from "@mui/material/Button";
import { useNavigate } from "react-router-dom";

const useStyles = makeStyles()(theme => ({
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

interface Props {
  open: boolean;
  setOpen: (open: boolean) => void;
}

export default function AcccountsAddedDialog({ open, setOpen }: Props) {
  const { classes } = useStyles();
  const navigate = useNavigate();

  return (
    <Dialog onClose={() => setOpen(false)} open={open}>
      <div className={classes.dialogHeader}>
        <div className={classes.closeButton}>
          <IconButton
            id="close-icon"
            onClick={() => setOpen(false)}
            size="large"
          >
            <CloseIcon />
          </IconButton>
        </div>
        <img src={financrooLogo} alt="financroo logo" />
      </div>
      <div id="modal-content" className={classes.dialogContent}>
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
          id="cancel-button"
          style={{ marginRight: 16 }}
          onClick={() => setOpen(false)}
        >
          Cancel
        </Button>
        <Button
          variant="outlined"
          id="start-investing-button"
          onClick={() => {
            navigate("/investments");
          }}
        >
          Start investing
        </Button>
      </div>
    </Dialog>
  );
}
