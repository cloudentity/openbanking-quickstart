import React, { useState } from "react";
import { makeStyles } from "@material-ui/core/styles";
import { Theme } from "@material-ui/core";
import Button from "@material-ui/core/Button";
import Checkbox from "@material-ui/core/Checkbox";
import Avatar from "@material-ui/core/Avatar";
import Alert from "@material-ui/lab/Alert";

import Chip from "../Chip";
import ApplicationAccessDrawer from "./ApplicationAccessDrawer";
import { getDate } from "../ApplicationSimpleCard";
import { drawerStyles, permissionsDict } from "./utils";
import { uniq } from "ramda";

const useStyles = makeStyles((theme: Theme) => ({
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

function getAccounts(accountIds, accounts) {
  return accountIds
    .map((id) => accounts.find((v) => v.id === id))
    .filter((v) => v);
}

type Props = {
  drawerData: any;
  accounts: any;
  setDrawerData: (data: string | null) => void;
  handleRevoke: (id: string) => void;
  status: string;
};

function ApplicationAccessPaymentDrawer({
  drawerData,
  setDrawerData,
  handleRevoke,
  accounts,
  status,
}: Props) {
  const classes = useStyles();
  const [revokeAccess, setRevokeAccess] = useState(false);
  const [revokeAccessAgree, setRevokeAccessAgree] = useState(false);

  const accountsDetails = getAccounts(
    drawerData?.account_access_consent?.AccountIDs ?? [],
    accounts
  ).reduce((acc, curr) => ({ ...acc, [curr.name]: "*" + curr.id }), {});

  const permissionDates = {
    Authorised: getDate(drawerData?.account_access_consent?.CreationDateTime),
    "Last updated": getDate(
      drawerData?.account_access_consent?.StatusUpdateDateTime
    ),
    "Active until": getDate(
      drawerData?.account_access_consent?.ExpirationDateTime
    ),
  };

  const clusters = uniq(
    drawerData?.account_access_consent?.Permissions.map(
      (v) => permissionsDict[v].Cluster
    )
  ) as any;

  const permissionItems = clusters.map((cluster) => ({
    title: cluster,
    items: Object.values(permissionsDict)
      .filter((p) => p.Cluster === cluster)
      .map((v) => v.Language),
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
              {drawerData?.Client?.name[0]?.toUpperCase()}
            </Avatar>

            <h3 className={classes.name}>{drawerData?.Client?.name}</h3>
            <div style={{ flex: 1 }} />
            <Chip type="active">{status}</Chip>
          </div>
        )
      }
      setDrawerData={setDrawerData}
      bottomBar={
        <>
          <Button
            variant="outlined"
            className={classes.button}
            onClick={() => {
              if (revokeAccess) {
                setRevokeAccess(false);
              } else {
                setDrawerData(null);
              }
            }}
          >
            Cancel
          </Button>
          <Button
            id={'revoke-access-button'}
            variant="outlined"
            className={classes.button}
            style={{
              color: "white",
              backgroundColor: "#BD271E",
              border: "none",
            }}
            disabled={revokeAccess && !revokeAccessAgree}
            onClick={() => {
              if (revokeAccess) {
                handleRevoke(drawerData?.consent_id);
                setRevokeAccess(false);
                setDrawerData(null);
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
        <div>
          <Alert
            variant="outlined"
            severity="warning"
            classes={{ root: classes.alertRoot, icon: classes.alertIcon }}
            color="error"
          >
            Warning: Deleteing this consent will remove access to all accounts
            to which you have previously granted access.
          </Alert>
          <div className={classes.revokeInfo}>
            Are you sure you want to revoke access for all accounts connected
            with this application?
          </div>
          <div className={classes.revokeInfoCheckbox}>
            <Checkbox
              checked={revokeAccessAgree}
              onChange={(e) => setRevokeAccessAgree(e.target.checked)}
              color="primary"
            />{" "}
            I agree
          </div>
        </div>
      ) : (
        <>
          <div>
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

          <div>
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

          <div>
            <div className={classes.subHeader}>Details being shared</div>
            <div>
              {permissionItems.map((v) => (
                <div key={v.title}>
                  <div className={classes.detailsTitle}>{v.title}</div>
                  <ul className={classes.ulList}>
                    {v.items.map((item) => (
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
