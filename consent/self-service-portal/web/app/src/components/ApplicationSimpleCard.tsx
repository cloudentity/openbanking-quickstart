import React from "react";
import { useHistory } from "react-router";
import { makeStyles } from "@material-ui/core/styles";
import { Theme } from "@material-ui/core";
import Avatar from "@material-ui/core/Avatar";
import { uniq } from "ramda";

import Chip from "./Chip";

const useStyles = (clickable: boolean) =>
  makeStyles((theme: Theme) => ({
    container: {
      background: "#FFFFFF",
      boxShadow:
        "0px 1px 1px rgba(0, 0, 0, 0.08), 0px 0px 1px rgba(0, 0, 0, 0.31)",
      borderRadius: 4,
      maxWidth: 850,
      margin: "0 auto 24px auto",
      padding: "15px 32px",
      boxSizing: "border-box",
      cursor: clickable ? "pointer" : "default",
    },
    header: {
      display: "flex",
      alignItems: "center",
      borderBottom: "1px solid #ECECEC",
      paddingBottom: 16,
    },
    name: {
      fontWeight: "normal",
      fontSize: 20,
      lineHeight: "32px",
      margin: "0 16px",
    },
    logo: {
      border: "1.5px solid #F4F4F4",
      borderRadius: 4,
      width: 48,
      height: 48,
      objectFit: "contain",
    },
    content: {
      display: "flex",
      paddingTop: 16,
      "& > div": {
        flex: 1,
      },
    },
    label: {
      fontWeight: "bold",
      fontSize: 12,
      lineHeight: "22px",
    },
    caption: {
      ...theme.custom.caption,
    },
  }));

const monthNames = [
  "January",
  "February",
  "March",
  "April",
  "May",
  "June",
  "July",
  "August",
  "September",
  "October",
  "November",
  "December",
];

export function getDate(date) {
  const d = new Date(date);
  if (d.getFullYear() === 1) return "N/A";
  return `${d.getDate()} ${monthNames[d.getMonth()]} ${d.getFullYear()}`;
}

function ApplicationSimpleCard({ client, clickable = true }) {
  const classes = useStyles(clickable)();
  const history = useHistory();

  const nonZeroStatusDateContents = client?.consents?.filter((v) => {
    const d = new Date(v?.StatusUpdateDateTime);
    return d.getFullYear() !== 1;
  });

  const newestConsent = nonZeroStatusDateContents?.reduce((prev, curr) =>
    prev?.StatusUpdateDateTime <
    curr?.StatusUpdateDateTime
      ? curr
      : prev
  );

  const oldestConsent = nonZeroStatusDateContents?.reduce((prev, curr) =>
    prev?.StatusUpdateDateTime >
    curr?.StatusUpdateDateTime
      ? curr
      : prev
  );

  const permissions = uniq(
    client?.consents
      ?.map(
        (v) =>
          (v.type === "account_access" && "Accounts") ||
          (v.type === "domestic_payment" && "Payments") ||
          null
      )
      .filter((v) => v)
  ).join(", ");

  const expirationDateTime = new Date(
    newestConsent?.ExpirationDateTime
  );

  const status =
    (expirationDateTime.getFullYear() !== 1 &&
      (expirationDateTime < new Date() ? "Expired" : "Active")) ||
    "Active";

  return (
    <div
      id={client.id}
      className={`${classes.container} application-card`}
      onClick={() => {
        if (clickable) {
          history.push(`/app/${client.id}`);
        }
      }}
    >
      <div className={classes.header}>
        {client.logo_uri ? (
          <img className={classes.logo} src={client.logo_uri} alt="app logo" />
        ) : (
          <Avatar
            variant="square"
            className={classes.logo}
            style={{ backgroundColor: "white", color: "#626576" }}
          >
            {client.name[0]?.toUpperCase()}
          </Avatar>
        )}
        <h3 className={classes.name}>{client.name}</h3>
        <div style={{ flex: 1 }} />
        <Chip type="active">{status}</Chip>
      </div>
      <div className={classes.content}>
        <div>
          <div className={classes.label}>Authorised</div>
          <div className={classes.caption}>
            {getDate(oldestConsent.CreationDateTime)}
          </div>
        </div>
        <div>
          <div className={classes.label}>Last updated</div>
          <div className={classes.caption}>
            {getDate(newestConsent.StatusUpdateDateTime)}
          </div>
        </div>
        <div>
          <div className={classes.label}>Active until</div>
          <div className={classes.caption}>
            {getDate(newestConsent.ExpirationDateTime)}
          </div>
        </div>
        <div>
          <div className={classes.label}>Permissions</div>
          <div className={classes.caption}>{permissions}</div>
        </div>
      </div>
    </div>
  );
}

export default ApplicationSimpleCard;
