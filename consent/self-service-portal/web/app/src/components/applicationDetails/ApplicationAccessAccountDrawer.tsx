import React, { useState } from "react";
import { makeStyles } from "tss-react/mui";
import Button from "@mui/material/Button";
import Checkbox from "@mui/material/Checkbox";
import Avatar from "@mui/material/Avatar";
import Alert from "@mui/material/Alert";

import Chip from "../Chip";
import ApplicationAccessDrawer from "./ApplicationAccessDrawer";
import { getDate } from "../ApplicationSimpleCard";
import { drawerStyles, permissionsDict } from "./utils";
import uniq from "lodash/uniq";
import { Consent, ConsentAccount } from "../types";

const useStyles = makeStyles()(() => ({
  ...drawerStyles,
  cardsWrapperGrid: {
    display: "grid",
    gridTemplateColumns: "1fr 1fr 1fr",
    gridColumnGap: 16,
    "& > div": {
      marginRight: 0,
    },
  },
  button: {
    width: "100%",
    "&:first-of-type": {
      marginRight: 16,
    },
    textTransform: "none",
    //  ...theme.custom.button,
    color: "#626576",
    "&:disabled": {
      backgroundColor: "#626576 !important",
    },
  },
  alertRoot: {
    backgroundColor: "#FFE3E6",
    border: "1px solid rgba(189, 39, 30, 0.3)",
    //   ...theme.custom.body2,
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

function getRevokeHeader() {
  return (
    <div
      style={{
        height: 72,
        backgroundColor: "#F7FAFF",
        width: "100%",
        display: "flex",
        alignItems: "center",
        paddingLeft: 32,
        fontWeight: 600,
        fontSize: 16,
        lineHeight: "24px",
        color: "#BD271E",
      }}
    >
      Revoke access
    </div>
  );
}

function getAccounts(accountIds: string[], accounts: ConsentAccount[]) {
  return accountIds.map(id => accounts.find(v => v.id === id)).filter(v => v);
}

interface Props {
  drawerData: Consent;
  accounts: ConsentAccount[];
  setDrawerData: (data: Consent | undefined) => void;
  handleRevoke: (id: string, consent_type: string) => void;
  status: string;
}

function ApplicationAccessPaymentDrawer({
  drawerData,
  setDrawerData,
  handleRevoke,
  accounts,
  status,
}: Props) {
  const { classes } = useStyles();
  const [revokeAccess, setRevokeAccess] = useState(false);
  const [revokeAccessAgree, setRevokeAccessAgree] = useState(false);

  const accountsDetails = getAccounts(
    drawerData?.AccountIDs ?? [],
    accounts
  ).reduce((acc, curr) => ({ ...acc, [curr?.name ?? ""]: "*" + curr?.id }), {});

  const permissionDates = {
    Authorised: getDate(drawerData?.CreationDateTime),
    "Last updated": getDate(drawerData?.StatusUpdateDateTime),
    "Active until": getDate(drawerData?.ExpirationDateTime),
  };

  const clusters = uniq(
    drawerData.Permissions?.map(v => permissionsDict[v].Cluster)
  );

  const permissionItems = clusters.map(cluster => ({
    title: cluster,
    items: Object.values(permissionsDict)
      .filter(p => p.Cluster === cluster)
      .map(v => v.Language),
  }));

  return (
    <ApplicationAccessDrawer
      header={
        revokeAccess ? (
          getRevokeHeader()
        ) : (
          <div className={classes.headerContent}>
            <Avatar
              variant="square"
              className={classes.logo}
              style={{ backgroundColor: "white", color: "#626576" }}
            >
              {drawerData.CreditorAccountName?.toUpperCase()}
            </Avatar>

            <h3 className={classes.name}>{drawerData.CreditorAccountName}</h3>
            <div style={{ flex: 1 }} />
            <Chip type="active">{status}</Chip>
          </div>
        )
      }
      setDrawerData={setDrawerData}
      bottomBar={
        <>
          <Button
            id="cancel-revoke-access-button"
            variant="outlined"
            className={classes.button}
            onClick={() => {
              if (revokeAccess) {
                setRevokeAccess(false);
              } else {
                setDrawerData(undefined);
              }
            }}
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
            disabled={
              (revokeAccess && !revokeAccessAgree) ||
              drawerData?.Status !== "Authorised"
            }
            onClick={() => {
              if (revokeAccess) {
                handleRevoke(drawerData?.ConsentID, drawerData?.type);
                setRevokeAccess(false);
                setDrawerData(undefined);
              } else {
                setRevokeAccess(true);
              }
            }}
          >
            Revoke access
          </Button>
        </>
      }
    >
      {revokeAccess ? (
        <div id="account-revoke-info">
          <Alert
            variant="outlined"
            severity="warning"
            classes={{ root: classes.alertRoot, icon: classes.alertIcon }}
            color="error"
          >
            Warning: Deleting this consent will remove access to all accounts
            to which you have previously granted access.
          </Alert>
          <div className={classes.revokeInfo}>
            Are you sure you want to revoke access for all accounts connected
            with this application?
          </div>
          <div className={classes.revokeInfoCheckbox}>
            <Checkbox
              id="account-revoke-checkbox"
              checked={revokeAccessAgree}
              onChange={e => setRevokeAccessAgree(e.target.checked)}
              color="primary"
            />{" "}
            I agree
          </div>
        </div>
      ) : (
        <>
          <div id="account-permission-dates">
            <div className={classes.subHeader}>Permission dates</div>
            <div className={classes.cardsWrapperGrid}>
              {Object.entries(permissionDates).map(([key, value]: any) => (
                <div className={classes.card} key={key}>
                  <div className={classes.cardTitle}>{key}</div>
                  <div className={classes.cardContent}>{value}</div>
                </div>
              ))}
            </div>
          </div>

          <div id="accounts-info">
            <div className={classes.subHeader}>Accounts</div>
            <div className={classes.cardsWrapperGrid}>
              {Object.entries(accountsDetails).map(([key, value]: any) => (
                <div
                  className={classes.card}
                  style={{
                    display: "flex",
                    alignItems: "baseline",
                    padding: 10,
                    justifyContent: "center",
                  }}
                  key={key}
                >
                  <div className={classes.cardTitle} style={{ marginRight: 8 }}>
                    {key}
                  </div>
                  <div className={classes.cardContent}>{value}</div>
                </div>
              ))}
            </div>
          </div>

          <div id="account-details">
            <div className={classes.subHeader}>Details being shared</div>
            <div>
              {permissionItems.map(v => (
                <div key={v.title}>
                  <div className={classes.detailsTitle}>{v.title}</div>
                  <ul className={classes.ulList}>
                    {v.items.map(item => (
                      <li key={item}>{item}</li>
                    ))}
                  </ul>
                </div>
              ))}
            </div>
          </div>
        </>
      )}
    </ApplicationAccessDrawer>
  );
}

export default ApplicationAccessPaymentDrawer;
