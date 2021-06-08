import React, { useState } from "react";
import { makeStyles } from "@material-ui/core/styles";
import { Theme } from "@material-ui/core";
import Drawer from "@material-ui/core/Drawer";
import Avatar from "@material-ui/core/Avatar";
import Button from "@material-ui/core/Button";
import Alert from "@material-ui/lab/Alert";
import Checkbox from "@material-ui/core/Checkbox";
import { uniq } from "ramda";

import {
  drawerStyles,
  getChipForStatus,
  getDate,
  getRawConsents,
  permissionsDict,
} from "./utils";

const useStyles = makeStyles((theme: Theme) => ({
  ...drawerStyles,
  header: {
    display: "flex",
    alignItems: "center",
    borderBottom: "1px solid #ECECEC",
    padding: "12px 24px",
  },
  container: {
    width: 500,
    marginBottom: 84,
    overflow: "auto",
    height: "100%",
  },
  content: {
    padding: 32,
    height: "calc(100% - 64px)",
  },
  name: {
    ...theme.custom.heading3,
  },
  id: {
    ...theme.custom.caption,
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
  empty: {
    height: "100%",
    display: "flex",
    alignItems: "center",
    justifyContent: "center",
    color: "darkgray",
  },
}));

function getRevokeHeader(
  revokeAccess: "client" | "consent" | null,
  consentToRevoke: string | null,
  clientName: string
) {
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
      {revokeAccess === "client" &&
        clientName &&
        `Revoke access for ${clientName}`}
      {revokeAccess === "consent" &&
        consentToRevoke &&
        `Revoke consent ${consentToRevoke}`}
    </div>
  );
}

interface PropTypes {
  data: any;
  setData: (data: string | null) => void;
  onRevokeClient: (id: string) => void;
  onRevokeConsent: (id: string) => void;
}

function CustomDrawer({
  data,
  setData,
  onRevokeClient,
  onRevokeConsent,
}: PropTypes) {
  const classes = useStyles();
  const [revokeAccess, setRevokeAccess] =
    useState<"client" | "consent" | null>(null);
  const [consentToRevoke, setConsentToRevoke] = useState<string | null>(null);
  const [revokeAccessAgree, setRevokeAccessAgree] = useState(false);

  const rawConsents = getRawConsents(data?.consents ?? []);

  return (
    <Drawer anchor="right" open={true} onClose={() => setData(null)}>
      {revokeAccess ? (
        getRevokeHeader(revokeAccess, consentToRevoke, data?.client_name)
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
              {revokeAccess === "consent"
                ? "Are you sure you want to revoke this consent?"
                : "Are you sure you want to revoke access for all accounts connected with this application?"}
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
            {rawConsents.length === 0 && (
              <div className={classes.empty}>No consents</div>
            )}
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
                <div
                  className={classes.consentContainer}
                  key={consent?.ConsentId}
                >
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
                      {getChipForStatus(consent?.Status)}
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
                          setRevokeAccess("consent");
                          setConsentToRevoke(consent?.ConsentId);
                        }}
                      >
                        Revoke consent
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
              setRevokeAccess(null);
            } else {
              setData(null);
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
          disabled={revokeAccess !== null && !revokeAccessAgree}
          onClick={() => {
            if (revokeAccess) {
              if (revokeAccess === "client") {
                onRevokeClient(data?.client_id);
              } else if (revokeAccess === "consent" && consentToRevoke) {
                onRevokeConsent(consentToRevoke);
              }
              setRevokeAccess(null);
              setData(null);
            } else {
              setRevokeAccess("client");
            }
          }}
        >
          {revokeAccess === "consent" ? "Revoke consent" : "Revoke access"}
        </Button>
      </div>
    </Drawer>
  );
}

export default CustomDrawer;
