import React, { useContext, useState } from "react";
import { CommonCtx } from "../services/common";
import Alert from '@mui/material/Alert';
import Typography from "@mui/material/Typography";
import Snackbar from "@mui/material/Snackbar";

export const Snacks = () => {
  const commons = useContext(CommonCtx);
  const { error, clearError } = commons!;
  const onClose = (e, reason) => {
    if (reason !== "clickaway") {
      setOpen(false);
      // workaround to prevent empty snackbar from rendering
      setTimeout(clearError, 100);
    }
  };
  const [open, setOpen] = useState<boolean>(!!error);

  return (
    <Snackbar
      open={open}
      autoHideDuration={6000}
      onClose={onClose}
      anchorOrigin={{ horizontal: "right", vertical: "bottom" }}
    >
      <Alert
        variant="filled"
        severity="error"
        onClose={() => onClose(null, "close-alert")}
      >
        <Typography>{error}</Typography>
      </Alert>
    </Snackbar>
  );
};
