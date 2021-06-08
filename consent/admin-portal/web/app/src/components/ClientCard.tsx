import React from "react";
import { makeStyles, Theme } from "@material-ui/core/styles";
import Avatar from "@material-ui/core/Avatar";
import { getChipForStatus } from "./utils";

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
    cursor: "pointer",
    "&:hover": {
      boxShadow:
        "0px 1px 16px -4px rgba(0, 0, 0, 0.25), 0px 0px 1px rgba(0, 0, 0, 0.31)",
    },
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
}));

export default function ClientCard({ client, onClick }) {
  const classes = useStyles();
  return (
    <div className={classes.card} onClick={onClick}>
      <div className={classes.cardName}>
        <Avatar variant="square" className={classes.avatar}>
          {client?.client_name[0]?.toUpperCase()}
        </Avatar>
        <div className={classes.name}>{client?.client_name}</div>
      </div>
      {getChipForStatus(client?.mainStatus)}
    </div>
  );
}
