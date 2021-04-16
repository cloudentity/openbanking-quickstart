import React, {useContext, useEffect, useState} from "react";
import ArrowBack from "@material-ui/icons/ArrowBack";
import IconButton from "@material-ui/core/IconButton";
import { makeStyles } from "@material-ui/core";
import { useHistory, useParams } from "react-router";

import PageToolbar from "../PageToolbar";
import Subheader from "../Subheader";
import ApplicationSimpleCard from "../ApplicationSimpleCard";
import ApplicationAccessTabs from "./ApplicationAccessTabs";
import { api } from "../../api/api";
import {CommonCtx} from "../../services/common";
import {Snacks} from "../Snacks";

const useStyles = makeStyles(() => ({
  backButton: {
    color: "white",
    marginRight: 10,
    position: "relative",
    bottom: 1,
  },
  content: {
    padding: "48px 0",
  },
}));

function ApplicationDetailsController({
  authorizationServerURL,
  authorizationServerId,
  tenantId,
}) {
  const { id } = useParams<Record<string, string | undefined>>();
  const classes = useStyles();
  const history = useHistory();
  const [isProgress, setProgress] = useState(true);
  const [clientConsent, setClientConsent] = useState<any>([]);
  const [accounts, setAccounts] = useState([]);
  const commons = useContext(CommonCtx);
  const setError = commons!.setError

  useEffect(() => {
    setProgress(true);
    api
      .getConsents()
      .then((res) => {
        const client = res.client_consents.find((v) => v.id === id);
        if (client) {
          setClientConsent(client);
        } else {
          setError && setError("Wrong client id");
        }
        setAccounts(res?.accounts?.accounts ?? []);
      })
      .catch((err) => {
        console.log(err);
        setError && setError(err.error)
      })
      .finally(() => setProgress(false));
  }, [id]);

  const handleRevoke = (id) => {
    setProgress(true);
    api
      .deleteConsent({ id })
      .then(api.getConsents)
      .then((res) => setClientConsent(res.client_consents))
      .then(() => history.push("/"))
      .catch((err) => {
        console.log(err);
        setError && setError(err.error)
      })
      .finally(() => setProgress(false));
  };

  const nonZeroStatusDateContents = clientConsent?.consents?.filter((v) => {
    const d = new Date(v?.account_access_consent?.StatusUpdateDateTime);
    return d.getFullYear() !== 1;
  });

  const newestConsent = nonZeroStatusDateContents?.reduce((prev, curr) =>
    prev?.account_access_consent?.StatusUpdateDateTime <
    curr?.account_access_consent?.StatusUpdateDateTime
      ? curr
      : prev
  );

  const expirationDateTime = new Date(
    newestConsent?.account_access_consent?.ExpirationDateTime
  );

  const status =
    (expirationDateTime.getFullYear() !== 1 &&
      (expirationDateTime < new Date() ? "Expired" : "Active")) ||
    "Active";

  return !isProgress
    ? (
        <div
          style={{
            background: "#F7FAFF",
            marginTop: 64,
            position: "relative",
            minHeight: "100vh",
          }}
        >
          <PageToolbar
            authorizationServerURL={authorizationServerURL}
            authorizationServerId={authorizationServerId}
            tenantId={tenantId}
          />
          <Subheader
            title={
              <div>
                <IconButton
                  onClick={() => {
                    history.push("/");
                  }}
                  className={classes.backButton}
                >
                  <ArrowBack />
                </IconButton>
                Connected apps
              </div>
            }
          />
          <div className={classes.content}>
            <ApplicationSimpleCard
              key={clientConsent.id}
              client={clientConsent}
              clickable={false}
            />
            <ApplicationAccessTabs
              data={clientConsent}
              accounts={accounts}
              handleRevoke={handleRevoke}
              status={status}
            />
          </div>
        </div>
      )
    : null;
}

export default ApplicationDetailsController;
