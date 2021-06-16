import React, { useState } from "react";
import { makeStyles, Theme } from "@material-ui/core/styles";
import Avatar from "@material-ui/core/Avatar";
import Button from "@material-ui/core/Button";

import { ClientType, getDate } from "./utils";
import RevokeDialog from "./ThirdPartyProvidersView/RevokeDialog";

const useStyles = makeStyles((theme: Theme) => ({
  card: {
    background: "#FFFFFF",
    padding: "12px 24px",
    boxShadow:
      "0px 1px 1px rgba(0, 0, 0, 0.08), 0px 0px 1px rgba(0, 0, 0, 0.31)",
    borderRadius: 4,
    display: "flex",
    alignItems: "center",
    justifyContent: "space-between",
  },
  date: {
    ...theme.custom.caption,
  },
  cardName: {
    display: "flex",
    alignItems: "center",
  },
  avatar: {
    background: "#FCFCFF",
    border: "1px solid #626576",
    color: "#626576",
    width: 48,
    height: 48,
    marginRight: 12,
    borderRadius: 4,
  },
  name: {
    ...theme.custom.heading3,
  },
  buttonRoot: {
    textTransform: "capitalize",
    ...theme.custom.body2,
    color: "#626576",
  },
}));

function getAuthorisedDate(client) {
  const accountAccessConsent = client?.consents?.find(
    (v) => v.type === "account_access"
  );
  if (accountAccessConsent) {
    const date = accountAccessConsent?.account_access_consent?.CreationDateTime;
    return getDate(date);
  }
  return null;
}

interface PropTypes {
  client: ClientType;
  onRevokeClient: (id: string) => void;
}

export default function ClientCard({ client, onRevokeClient }: PropTypes) {
  const classes = useStyles();
  const date = getAuthorisedDate(client);
  const [openDialog, setOpenDialog] = useState(false);

  return (
    <div className={classes.card}>
      <div className={classes.cardName}>
        <Avatar variant="square" className={classes.avatar}>
          {client?.client_name[0]?.toUpperCase()}
        </Avatar>
        <div>
          <div className={classes.name}>{client?.client_name}</div>
          {date && <div className={classes.date}>Connected {date}</div>}
        </div>
      </div>
      <Button
        variant="outlined"
        classes={{ root: classes.buttonRoot }}
        onClick={() => setOpenDialog(true)}
      >
        Revoke
      </Button>

      {openDialog && (
        <RevokeDialog
          handleClose={() => setOpenDialog(false)}
          onConfirm={() => onRevokeClient(client?.client_id)}
          clientName={client?.client_name}
        />
      )}
    </div>
  );
}
