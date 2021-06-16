import React from "react";
import DialogContent from "@material-ui/core/DialogTitle";
import DialogActions from "@material-ui/core/DialogActions";
import Dialog from "@material-ui/core/Dialog";
import Button from "@material-ui/core/Button";
import { makeStyles, Theme } from "@material-ui/core";
import clsx from "clsx";

const useStyles = makeStyles((theme: Theme) => ({
  button: {
    ...theme.custom.button,
    padding: "6px 18px",
    color: "#212533",
    fontSize: 14,
  },
  revokeButton: {
    backgroundColor: "#DC1B37",
    color: "white",
    "&:hover": {
      backgroundColor: "#DC1B37",
      opacity: 0.9,
    },
  },
  content: {
    ...theme.custom.body2,
    padding: "18px 0",
  },
  actions: {
    borderTop: "solid 1px #ECECEC",
  },
}));

interface PropTypes {
  onConfirm: () => void;
  handleClose: () => void;
  clientName?: string;
}

export default function RevokeDialog({
  onConfirm,
  handleClose,
  clientName,
}: PropTypes) {
  const classes = useStyles();

  return (
    <Dialog onClose={handleClose} open={true} id="revoke-confirm-dialog">
      <DialogContent>
        <div className={classes.content}>
          Are your sure you want to revoke access for <b>{clientName}</b>?
        </div>
      </DialogContent>
      <DialogActions classes={{ root: classes.actions }}>
        <Button autoFocus onClick={handleClose} className={classes.button}>
          Cancel
        </Button>
        <Button
          onClick={onConfirm}
          className={clsx(classes.button, classes.revokeButton)}
          id="revoke-confirm-button"
        >
          Revoke access
        </Button>
      </DialogActions>
    </Dialog>
  );
}
