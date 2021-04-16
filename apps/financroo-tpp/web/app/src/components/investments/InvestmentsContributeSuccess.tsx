import React from "react";
import {useHistory, useParams} from "react-router";
import Button from "@material-ui/core/Button";
import Chip from "@material-ui/core/Chip";
import { makeStyles } from "@material-ui/core/styles";

import PageContainer from "../common/PageContainer";
import PageToolbar from "../common/PageToolbar";
import ContributionCard from "./ContributionCard";
import Field from "./Field";
import Confetti from "./Confetti";
import { theme } from "../../theme";
import bankIcon from "../../assets/banks/gobank-icon.svg";
import checkIcon from "../../assets/icon-check.svg";
import qs from 'query-string';

const useStyles = makeStyles((theme) => ({
  title: {
    marginBottom: 68,
    display: "flex",
    flexDirection: "column",
    alignItems: "center",
    "& > h2": {
      fontWeight: "normal",
      fontSize: 28,
      lineHeight: "40px",
      color: "#626576",
      marginBottom: 16,
    },
    "& > div": {
      fontSize: 16,
      lineHeight: "24px",
      color: "#626576",
    },
  },
  chip: {
    backgroundColor: theme.palette.primary.main,
    color: "white",
  },
  grid: {
    display: "grid",
    gridTemplateColumns: "1fr 1fr 1fr",
    gridColumnGap: 14,
  },
  card: {
    position: "relative",
    display: "flex",
    flexDirection: "column",
    alignItems: "center",
    justifyContent: "center",
    padding: "16px 14px",
    backgroundColor: "#FBFCFD",
    border: "1px solid #36C6AF",
    boxSizing: "border-box",
    boxShadow:
      "0px 1px 1px rgba(0, 0, 0, 0.04), 0px 3px 2px rgba(0, 0, 0, 0.04)",
    borderRadius: 4,
    "& > img": {
      width: 29,
    },
    "& > span": {
      fontSize: 12,
      lineHeight: "22px",
      color: "#626576",
    },
  },
  information: {
    display: "grid",
    gridTemplateColumns: "1fr 1fr",
    gridColumnGap: 8,
    justifyContent: "space-between",
    padding: 20,
    paddingBottom: 8,
    color: "#626576",
    fontSize: 12,
    lineHeight: "22px",
    background: "#FBFCFD",
    border: "1px solid #F4F4F4",
    borderRadius: 4,
    "& :nth-child(2n)": {
      textAlign: "right",
    },
    "& > div": {
      paddingBottom: 12,
    },
  },
  bankLogoImage: {
    width: 15,
    marginLeft: 20,
  },
  header: {
    textAlign: "center",
    ...theme.custom.label,
  },
}));

export default function InvestmentsContributeSuccess({
  authorizationServerURL,
  authorizationServerId,
  tenantId,
}) {
  const history = useHistory();
  const classes = useStyles();
  const {id} = useParams<{id: string}>();
  const search = history.location.search
  const searchParsed = qs.parse(search)

  return (
    <div style={{ position: "relative" }}>
      <PageToolbar
        mode="main"
        authorizationServerURL={authorizationServerURL}
        authorizationServerId={authorizationServerId}
        tenantId={tenantId}
        tab="investments"
      />
      <PageContainer
        style={{ paddingTop: 48, paddingBottom: 48 }}
        containerStyle={{
          display: "flex",
          flexDirection: "column",
          alignItems: "center",
        }}
      >
        <div className={classes.title}>
          <img src={checkIcon} alt="check icon" />
          <h2>Transaction completed</h2>
          <div>You have reached another level with your retirement goal</div>
        </div>
        <ContributionCard>
          <Confetti />
          <Field
            label={
              <div style={{ textAlign: "center" }}>Account Information</div>
            }
          >
            <div
              className={classes.information}
              style={{ alignItems: "center", paddingBottom: 20 }}
            >
              <div style={{ ...theme.custom.caption }}>Transaction ID</div>
              <div>{id}</div>
              <div style={{ ...theme.custom.label, paddingBottom: 0 }}>
                Total contribution paid
              </div>
              <div style={{ paddingBottom: 0 }}>
                <Chip label={`${searchParsed.amount} ${searchParsed.currency}`} className={classes.chip} />
              </div>
            </div>
          </Field>

          <Field
            label={<div style={{ textAlign: "center" }}>Payment details</div>}
            style={{ marginBottom: 0 }}
          >
            <div
              className={classes.information}
              style={{ alignItems: "center", paddingBottom: 20 }}
            >
              <div style={{ ...theme.custom.caption }}>Order reference</div>
              <div>Financoo investments Ltd</div>
              <div style={{ ...theme.custom.label, paddingBottom: 0 }}>
                Bank name
              </div>
              <div>
                Go Bank{" "}
                <img
                  src={bankIcon}
                  alt="bank logo"
                  className={classes.bankLogoImage}
                />
              </div>
            </div>
          </Field>
        </ContributionCard>
        <Button
          onClick={() => history.push("/investments")}
          variant="contained"
          color="primary"
          style={{ ...theme.custom.button, marginTop: 44 }}
          disableElevation
        >
          Back to portfolio
        </Button>
      </PageContainer>
    </div>
  );
}
