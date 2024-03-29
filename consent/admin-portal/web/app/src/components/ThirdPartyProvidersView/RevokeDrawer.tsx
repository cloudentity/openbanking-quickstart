import React, { useState } from "react";
import Checkbox from "@mui/material/Checkbox";
import Alert from "@mui/material/Alert";
import Button from "@mui/material/Button";
import { makeStyles } from "tss-react/mui";
import CustomDrawer from "../AccountView/Drawers/CustomDrawer";
import { ClientType, drawerStyles } from "../utils";

const useStyles = makeStyles()(theme => ({
  ...drawerStyles,
  button: {
    width: "100%",
    "&:first-of-type": {
      marginRight: 16,
    },
    ...theme.custom.button,
    color: "#626576",
    "&:disabled": {
      backgroundColor: "#626576 !important",
    },
  },
  alertRoot: {
    backgroundColor: "#FFE3E6",
    border: "1px solid rgba(189, 39, 30, 0.3)",
    ...theme.custom.body2,
  },
  alertIcon: {
    position: "relative",
    top: 2,
  },
  revokeInfo: {
    fontSize: 16,
    lineHeight: "24px",
    margin: "32px 0",
  },
  revokeInfoCheckbox: {
    display: "flex",
    alignItems: "center",
    "& > span": {
      marginRight: 3,
    },
  },
}));

interface Props {
  onConfirm: () => void;
  handleClose: () => void;
  client?: ClientType;
}

export default function RevokeDrawer({
  onConfirm,
  handleClose,
  client,
}: Props) {
  const { classes } = useStyles();
  const [revokeAccessAgree, setRevokeAccessAgree] = useState(false);

  return (
    <CustomDrawer
      header={
        <div className={classes.drawerHeader}>Revoke third party access</div>
      }
      handleClose={() => handleClose()}
      bottomBar={
        <>
          <Button
            variant="outlined"
            className={classes.button}
            onClick={() => handleClose()}
          >
            Cancel
          </Button>
          <Button
            id="revoke-access-button"
            variant="outlined"
            className={classes.button}
            style={{
              color: "white",
              backgroundColor: "#BD271E",
              border: "none",
            }}
            disabled={!revokeAccessAgree}
            onClick={onConfirm}
          >
            Revoke access
          </Button>
        </>
      }
    >
      <div>
        <Alert
          variant="outlined"
          severity="warning"
          classes={{ root: classes.alertRoot, icon: classes.alertIcon }}
          color="error"
        >
          Warning: Warning: Go Bank members will no longer be able to use{" "}
          <b>{client?.client_name}</b> with all Go Bank accounts
        </Alert>
        <div className={classes.revokeInfo}>
          Are you sure you want to revoke all API access for{" "}
          <b>{client?.client_name}</b>?
        </div>
        <div className={classes.revokeInfoCheckbox}>
          <Checkbox
            checked={revokeAccessAgree}
            onChange={e => setRevokeAccessAgree(e.target.checked)}
            color="primary"
            id="revoke-access-accept-checkbox"
          />{" "}
          I agree
        </div>
      </div>
    </CustomDrawer>
  );
}
