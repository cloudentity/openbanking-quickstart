import React from "react";
import { useNavigate } from "react-router-dom";
import { makeStyles } from "tss-react/mui";

import PageContainer from "../common/PageContainer";
import PageToolbar from "../common/PageToolbar";
import dashboardImgGBP from "../../assets/investments-dashboard-GBP.svg";
import dashboardImgBRL from "../../assets/investments-dashboard-BRL.svg";
import dashboardImgUSD from "../../assets/investments-dashboard-USD.svg";
import dashboardImgEUR from "../../assets/investments-dashboard-EUR.svg";

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

function getDashboardImage(currency: string) {
  switch (currency) {
    case "USD":
      return dashboardImgUSD;
    case "AUD":
      return dashboardImgUSD;
    case "GBP":
      return dashboardImgGBP;
    case "EUR":
      return dashboardImgEUR;
    case "BRL":
      return dashboardImgBRL;
    default:
      return dashboardImgGBP;
  }
}

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
            if (window.spec !== "cdr" && window.spec !== "fdx") {
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
          id={`dashboard-${window.currency}`}
          src={getDashboardImage(window.currency || "GBP")}
          className={classes.dashboardImage}
        />
      </PageContainer>
    </div>
  );
}
