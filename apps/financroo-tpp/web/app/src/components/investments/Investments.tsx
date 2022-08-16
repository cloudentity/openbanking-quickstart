import React from "react";
import { useNavigate } from "react-router-dom";
import { makeStyles } from "tss-react/mui";

import PageContainer from "../common/PageContainer";
import PageToolbar from "../common/PageToolbar";
import dashboardImg from "../../assets/investments-dashboard.svg";

const useStyles = makeStyles()(theme => ({
  dashboardImage: {
    width: "100%",
    [theme.breakpoints.down("lg")]: {
      position: "absolute",
      left: 24,
      right: 24,
      width: "calc(100% - 48px)",
    },
  },
}));

export default function Investments() {
  const navigate = useNavigate();
  const { classes } = useStyles();

  return (
    <div style={{ position: "relative" }}>
      <PageToolbar
        mode="main"
        tab="investments"
        subHeaderTitle="Investments"
        subHeaderButton={{
          id: "invest-button",
          title: "Contribute now",
          onClick: () => {
            if (window.spec !== "cdr") {
              navigate("/investments/contribute");
            }
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
