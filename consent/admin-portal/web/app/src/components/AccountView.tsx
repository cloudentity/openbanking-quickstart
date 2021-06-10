import React, { useEffect, useState } from "react";
import { makeStyles, Theme } from "@material-ui/core";
import { useHistory, useParams } from "react-router";
import ArrowBack from "@material-ui/icons/ArrowBack";
import IconButton from "@material-ui/core/IconButton";

import CustomTabs from "./CustomTabs";
import { searchTabs } from "./Dashboard";
import PageToolbar from "./PageToolbar";
import Progress from "./Progress";
import Subheader from "./Subheader";

const useStyles = makeStyles((theme: Theme) => ({
  subtitle: {
    ...theme.custom.body1,
  },
  subHeaderToolbar: {
    backgroundColor: "#002D4C",
    height: 80,
    maxWidth: 850,
    margin: "0 auto",
    display: "flex",
  },
  header: {
    ...theme.custom.heading3,
    borderBottom: "solid 1px #ECECEC",
    paddingBottom: 24,
  },
  container: {
    maxWidth: 850,
    margin: "32px auto",
  },
  accountInfo: {
    backgroundColor: "white",
    boxShadow: "0px 0px 0px 1px #ECECEC",
    padding: 24,
    borderRadius: 4,
    marginBottom: 24,
  },
  subheader: {
    ...theme.custom.heading6,
  },
}));

interface PropTypes {
  authorizationServerURL?: string;
  authorizationServerId?: string;
  tenantId?: string;
}

export default function AccountView({
  authorizationServerURL,
  authorizationServerId,
  tenantId,
}: PropTypes) {
  const { id } = useParams<Record<string, string | undefined>>();
  const history = useHistory();
  const [isProgress, setProgress] = useState(true);
  const classes = useStyles();

  useEffect(() => {
    setProgress(false);
  }, []);

  return (
    <div>
      <PageToolbar
        authorizationServerURL={authorizationServerURL}
        authorizationServerId={authorizationServerId}
        tenantId={tenantId}
      />
      <div style={{ position: "relative" }}>
        {isProgress && <Progress />}
        {!isProgress && (
          <>
            <Subheader
              title=""
              containerStyle={{
                backgroundColor: "transparent",
                padding: 0,
                textAlign: "left",
              }}
              contentStyle={{ maxWidth: "100%" }}
            >
              <div style={{ backgroundColor: "#002D4C" }}>
                <div className={classes.subHeaderToolbar}>
                  <IconButton
                    style={{ padding: 4 }}
                    onClick={() => history.push("/")}
                    onMouseDown={() => history.push("/")}
                  >
                    <ArrowBack fontSize="small" style={{ color: "white" }} />
                  </IconButton>
                </div>
              </div>
              <div
                style={{
                  maxWidth: 590,
                  margin: "0 auto",
                  position: "relative",
                  top: -32,
                }}
              >
                <CustomTabs tabs={searchTabs(history)} />
              </div>
            </Subheader>
            <div className={classes.container}>
              <div className={classes.accountInfo}>
                <div className={classes.subheader}>Account</div>
                <div>{id}</div>
              </div>
              <div className={classes.header}>All Connected Applications</div>
            </div>
          </>
        )}
      </div>
    </div>
  );
}
