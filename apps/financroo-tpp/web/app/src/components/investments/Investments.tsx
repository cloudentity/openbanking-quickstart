import React from "react";
import { useHistory } from "react-router";
import { makeStyles } from "@material-ui/core/styles";
import { Theme } from "@material-ui/core";

import PageContainer from "../common/PageContainer";
import PageToolbar from "../common/PageToolbar";
import dashboardImg from "../../assets/investments-dashboard.svg";

const useStyles = makeStyles((theme: Theme) => ({
  dashboardImage: {
    width: "100%",
    [theme.breakpoints.down("md")]: {
      position: "absolute",
      left: 24,
      right: 24,
      width: "calc(100% - 48px)",
    },
  },
}));

export default function Investments({
  authorizationServerURL,
  authorizationServerId,
  tenantId,
}) {
  const history = useHistory();
  const classes = useStyles();

  return (
    <div style={{ position: "relative" }}>
      <PageToolbar
        mode="main"
        authorizationServerURL={authorizationServerURL}
        authorizationServerId={authorizationServerId}
        tenantId={tenantId}
        tab="investments"
        subHeaderTitle="Investments"
        subHeaderButton={{
          title: "Contribute now",
          onClick: () => {
            history.push("/investments/contribute");
          },
        }}
      />
      <PageContainer
        withSubheader
        style={{ paddingTop: 48, paddingBottom: 48 }}
      >
        <img
          alt="financroo logo"
          src={dashboardImg}
          className={classes.dashboardImage}
        />
      </PageContainer>
    </div>
  );
}
