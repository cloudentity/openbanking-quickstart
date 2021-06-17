import React, { useState } from "react";
import { makeStyles, Theme } from "@material-ui/core/styles";
import Avatar from "@material-ui/core/Avatar";
import Button from "@material-ui/core/Button";
import { uniq } from "ramda";
import clsx from "clsx";
import { useHistory } from "react-router";

import {
  ClientType,
  enrichClientWithStatus,
  getChipForStatus,
  getDate,
  getRawConsents,
} from "../utils";
import RevokeDrawer from "../ThirdPartyProvidersView/RevokeDrawer";

const useStyles = makeStyles((theme: Theme) => ({
  card: {
    background: "#FFFFFF",
    padding: "12px 24px",
    boxShadow:
      "0px 1px 1px rgba(0, 0, 0, 0.08), 0px 0px 1px rgba(0, 0, 0, 0.31)",
    borderRadius: 4,
  },
  clickable: {
    "&:hover": {
      boxShadow:
        "0px 1px 16px -4px rgba(0, 0, 0, 0.25), 0px 0px 1px rgba(0, 0, 0, 0.31)",
      cursor: "pointer",
      "& button": {
        backgroundColor: "#002D4C",
        color: "white",
      },
    },
  },
  header: {
    paddingBottom: 16,
    marginBottom: 16,
    borderBottom: "solid 1px #ECECEC",
    display: "flex",
    alignItems: "center",
    justifyContent: "space-between",
  },
  cardName: {
    display: "flex",
    alignItems: "center",
    justifyContent: "space-between",
    width: "100%",
    "& > div": {
      display: "flex",
      alignItems: "center",
    },
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
  details: {
    display: "flex",
    alignItems: "center",
    justifyContent: "space-between",
  },
  label: {
    ...theme.custom.label,
  },
  caption: {
    ...theme.custom.caption,
  },
  manageButton: {
    textTransform: "capitalize",
    ...theme.custom.body2,
    color: "#002D4C",
    padding: "8px 24px",
    backgroundColor: "#FCFCFF",
    border: "1px solid #002D4C",

    "&:hover": {
      backgroundColor: "#E0E0E0",
    },
  },
  revokeButton: {
    ...theme.custom.button,
    backgroundColor: "#DC1B37",
    padding: "8px 24px",
    "&:hover": {
      backgroundColor: "#DC1B37",
      opacity: 0.9,
    },
  },
}));

const consentTypesMapper = {
  account_access_consent: "Accounts",
  domestic_payment_consent: "Payments",
  domestic_scheduled_payment_consent: "Payments",
  domestic_standing_order_consent: null,
  file_payment_consent: "Payments",
  international_payment_consent: "Payments",
  international_scheduled_payment_consent: "Payments",
  international_standing_order_consent: null,
};

interface PropTypes {
  client?: ClientType;
  accountId?: string;
  accounts?: string[];
  onRevokeClient?: (id: string) => void;
}

export default function AccountClientCard({
  client,
  accountId,
  accounts,
  onRevokeClient,
}: PropTypes) {
  const classes = useStyles();
  const history = useHistory();
  const [openDrawer, setOpenDrawer] = useState(false);

  const rawConsents = getRawConsents(client?.consents ?? []);

  const accountAccessConsent = rawConsents.find(
    (v) => v?.type === "account_access_consent"
  );

  const permissionDates = {
    authorised: getDate(accountAccessConsent?.consent?.CreationDateTime),
    lastUpdated: getDate(accountAccessConsent?.consent?.StatusUpdateDateTime),
    activeUntil: getDate(accountAccessConsent?.consent?.ExpirationDateTime),
  };

  const types = rawConsents
    .map(({ type }) => consentTypesMapper[type] || null)
    .filter((v) => v);

  const isApplicationListView = accountId && accounts;
  const clientWithStatus = client && enrichClientWithStatus(client);

  return (
    <div
      className={clsx(classes.card, isApplicationListView && classes.clickable)}
      onClick={() => {
        if (isApplicationListView) {
          history.push({
            pathname: `/accounts/${accountId}/apps/${client?.client_id}`,
            state: { accounts, client: clientWithStatus },
          });
        }
      }}
    >
      <div className={classes.header}>
        <div className={classes.cardName}>
          <div>
            <Avatar variant="square" className={classes.avatar}>
              {clientWithStatus?.client_name[0]?.toUpperCase()}
            </Avatar>
            <div className={classes.name}>{clientWithStatus?.client_name}</div>
          </div>
          {getChipForStatus(clientWithStatus?.mainStatus)}
        </div>
      </div>
      <div className={classes.details}>
        <div>
          <div className={classes.label}>Access granted</div>
          <div className={classes.caption}>{permissionDates.lastUpdated}</div>
        </div>
        <div>
          <div className={classes.label}>Last updated</div>
          <div className={classes.caption}>{permissionDates.lastUpdated}</div>
        </div>
        <div>
          <div className={classes.label}>Active until</div>
          <div className={classes.caption}>{permissionDates.activeUntil}</div>
        </div>
        <div>
          <div className={classes.label}>Permissions</div>
          <div className={classes.caption}>{uniq(types).join(", ") || "-"}</div>
        </div>
        {isApplicationListView ? (
          <div>
            <div></div>
            <div>
              <Button className={classes.manageButton}>Manage</Button>
            </div>
          </div>
        ) : (
          <Button
            className={classes.revokeButton}
            onClick={() => setOpenDrawer(true)}
          >
            Revoke access
          </Button>
        )}
      </div>

      {openDrawer && (
        <RevokeDrawer
          handleClose={() => setOpenDrawer(false)}
          onConfirm={() =>
            client?.client_id &&
            onRevokeClient &&
            onRevokeClient(client?.client_id)
          }
          client={client}
        />
      )}
    </div>
  );
}
