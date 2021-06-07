import React, { useState } from "react";
import { makeStyles } from "@material-ui/core/styles";
import { Theme } from "@material-ui/core";
import Drawer from "@material-ui/core/Drawer";
import Avatar from "@material-ui/core/Avatar";
import Button from "@material-ui/core/Button";
import Alert from "@material-ui/lab/Alert";
import Checkbox from "@material-ui/core/Checkbox";
import { uniq } from "ramda";

import { drawerStyles, getDate, permissionsDict } from "./utils";
import Chip from "./Chip";

const useStyles = makeStyles((theme: Theme) => ({
  ...drawerStyles,
  container: {
    width: 500,
    marginBottom: 84,
    overflow: "auto",
  },
  header: {
    display: "flex",
    alignItems: "center",
    borderBottom: "1px solid #ECECEC",
    padding: "12px 24px",
  },
  name: {
    ...theme.custom.heading3,
  },
  id: {
    ...theme.custom.caption,
  },
  content: {
    padding: 32,
  },
  bottomBar: {
    position: "absolute",
    bottom: 0,
    boxShadow:
      "0px 6px 10px rgba(0, 0, 0, 0.14), 0px 1px 18px rgba(0, 0, 0, 0.12), 0px 3px 5px rgba(0, 0, 0, 0.2)",
    width: "100%",
    padding: 24,
    display: "flex",
    justifyContent: "space-between",
    boxSizing: "border-box",
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
  cardsWrapperGrid: {
    display: "grid",
    gridTemplateColumns: "1fr 1fr 1fr",
    gridColumnGap: 16,
    "& > div": {
      marginRight: 0,
    },
  },
  consentContainer: {
    borderBottom: "1px solid darkgray",
    paddingBottom: 24,
    marginBottom: 48,
    "&:last-of-type": {
      borderBottom: "none",
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

interface PropTypes {
  data: any;
  setData: (data: string | null) => void;
  onRevoke: (id: string) => void;
}

const availableConstentTypes = [
  "account_access_consent",
  "domestic_payment_consent",
  "domestic_scheduled_payment_consent",
  "domestic_standing_order_consent",
  "file_payment_consent",
  "international_payment_consent",
  "international_scheduled_payment_consent",
  "international_standing_order_consent",
];

function CustomDrawer({ data, setData, onRevoke }: PropTypes) {
  const classes = useStyles();
  const [revokeAccess, setRevokeAccess] = useState(false);
  const [revokeAccessAgree, setRevokeAccessAgree] = useState(false);

  const rawConsents = data?.consents?.reduce((acc, consent) => {
    const consents = Object.entries(consent)
      .map(([key, value]: [key: string, value: any]) =>
        availableConstentTypes.includes(key) && value?.ConsentId
          ? { type: key, consent: value, accounts: consent?.account_ids ?? [] }
          : null
      )
      .filter((v) => v);
    return [...acc, ...consents, ...consents];
  }, []);

  return (
    <Drawer anchor="right" open={true} onClose={() => setData(null)}>
      {revokeAccess ? (
        getRevokeHeader()
      ) : (
        <div className={classes.header}>
          <Avatar variant="square" className={classes.avatar}>
            {data?.client_name[0]?.toUpperCase()}
          </Avatar>
          <div>
            <div className={classes.name}>{data?.client_name}</div>
            <div className={classes.id}>Client ID: {data?.client_id}</div>
          </div>
        </div>
      )}
      <div className={classes.container}>
        {revokeAccess ? (
          <div className={classes.content}>
            <Alert
              variant="outlined"
              severity="warning"
              classes={{ root: classes.alertRoot, icon: classes.alertIcon }}
              color="error"
            >
              Warning: Deleteing this consent will remove access to all
              accounts.
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
          <div className={classes.content}>
            {rawConsents.map(({ type, consent, accounts }) => {
              const permissionDates = {
                Authorised: getDate(consent?.CreationDateTime),
                "Last updated": getDate(consent?.StatusUpdateDateTime),
                "Active until": getDate(consent?.ExpirationDateTime),
              };

              const clusters = uniq(
                consent.Permissions.map((v) => permissionsDict[v].Cluster)
              ) as any;

              const permissionItems = clusters.map((cluster) => ({
                title: cluster,
                items: Object.values(permissionsDict)
                  .filter((p) => p.Cluster === cluster)
                  .map((v) => v.Language),
              }));

              return (
                <div className={classes.consentContainer}>
                  <div>
                    <div
                      className={classes.subHeader}
                      style={{
                        display: "flex",
                        justifyContent: "space-between",
                        alignItems: "center",
                      }}
                    >
                      <span>Consent ID</span>
                      <Chip type="active">{consent?.Status}</Chip>
                    </div>
                    <div
                      style={{
                        paddingBottom: 16,
                        display: "flex",
                        justifyContent: "space-between",
                        alignItems: "center",
                      }}
                    >
                      <span>{consent?.ConsentId}</span>
                      <Button
                        style={{
                          backgroundColor: "#bd271e",
                          marginLeft: 4,
                          color: "white",
                          textTransform: "none",
                          fontWeight: 400,
                        }}
                        onClick={() => {
                          console.log("revoke access for", consent?.ConsentId);
                        }}
                      >
                        Revoke access
                      </Button>
                    </div>
                  </div>
                  <div>
                    <div className={classes.subHeader}>Consent Type</div>
                    <div
                      style={{ textTransform: "capitalize", paddingBottom: 16 }}
                    >
                      {type?.replaceAll("_", " ")}
                    </div>
                  </div>
                  <div>
                    <div className={classes.subHeader}>Permission dates</div>
                    <div className={classes.cardsWrapperGrid}>
                      {Object.entries(permissionDates).map(
                        ([key, value]: any) => (
                          <div className={classes.card} key={key}>
                            <div className={classes.cardTitle}>{key}</div>
                            <div className={classes.cardContent}>{value}</div>
                          </div>
                        )
                      )}
                    </div>
                  </div>

                  <div>
                    <div className={classes.subHeader}>Accounts</div>
                    <div className={classes.cardsWrapperGrid}>
                      {accounts.map((id: string) => (
                        <div
                          className={classes.card}
                          style={{
                            display: "flex",
                            alignItems: "baseline",
                            padding: 10,
                            justifyContent: "center",
                          }}
                          key={id}
                        >
                          <div className={classes.cardContent}>*{id}</div>
                        </div>
                      ))}
                    </div>
                  </div>

                  <div>
                    <div className={classes.subHeader}>
                      Details being shared
                    </div>
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
                </div>
              );
            })}
          </div>
        )}
      </div>
      <div className={classes.bottomBar}>
        <Button
          variant="outlined"
          className={classes.button}
          onClick={() => {
            if (revokeAccess) {
              setRevokeAccess(false);
            } else {
              setData(null);
            }
          }}
        >
          Cancel
        </Button>
        <Button
          id={"revoke-access-button"}
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
              onRevoke(data?.client_id);
              setRevokeAccess(false);
              setData(null);
            } else {
              setRevokeAccess(true);
            }
          }}
        >
          Revoke all
        </Button>
      </div>
    </Drawer>
  );
}

export default CustomDrawer;
