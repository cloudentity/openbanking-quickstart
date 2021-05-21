import React from "react";
import { makeStyles } from "@material-ui/core/styles";
import { Theme } from "@material-ui/core";

import Chip from "../Chip";
import logo from "../../assets/welcome-image.png";
import ApplicationAccessDrawer from "./ApplicationAccessDrawer";
import { getDate } from "../ApplicationSimpleCard";
import { currencyDict, drawerStyles } from "./utils";

const useStyles = makeStyles((theme: Theme) => ({
  ...drawerStyles,
  cardsWrapper: {
    display: "flex",
    flexWrap: "wrap",
  },
}));

type Props = {
  drawerData: any;
  setDrawerData: (data: string | null) => void;
  status: string;
};

function ApplicationAccessPaymentDrawer({
  drawerData,
  setDrawerData,
  status,
}: Props) {
  const classes = useStyles();

  const details = drawerData?.domestic_payment_consent;

  const transactionDetails = {
    Amount: `${
      currencyDict[details?.Initiation?.InstructedAmount?.Currency] ||
      currencyDict.GBP
    } ${details?.Initiation?.InstructedAmount?.Amount}`,
    Status: details?.Status,
    "Consent id": details?.ConsentId,
    "Debtor Account": {
      id: details?.Initiation?.DebtorAccount?.Identification,
      name: details?.Initiation?.DebtorAccount?.Name,
    },
    "Creditor Account": {
      id: details?.Initiation?.CreditorAccount?.Identification,
      name: details?.Initiation?.CreditorAccount?.Name,
    },
  };

  const permissionDates = {
    Authorised: getDate(details?.Authorisation?.CompletionDateTime),
    "Last updated": getDate(details?.StatusUpdateDateTime),
    // "Active until": "N/A",
  };

  return (
    <ApplicationAccessDrawer
      header={
        <div className={classes.headerContent}>
          <img className={classes.logo} src={logo} alt="app logo" />
          <h3 className={classes.name}>Financroo</h3>
          <div style={{ flex: 1 }} />
          <Chip type="active">{status}</Chip>
        </div>
      }
      setDrawerData={setDrawerData}
    >
      <div className={classes.purpose}>
        <div className={`${classes.purposeHeader} purpose-header`}>Purpose for sharing data:</div>
        <div>To enable payments to Financroo investments</div>
      </div>

      <div>
        <div className={classes.subHeader}>TRANSACTION Details</div>
        <div className={classes.cardsWrapper}>
          {Object.entries(transactionDetails).map(([key, value]: any) => (
            <div className={classes.card} key={key}>
              <div className={classes.cardTitle}>{key}</div>
              <div className={classes.cardContent}>
                {value.id ? (
                  <>
                    <div>{value.id}</div>
                    <div>{value.name}</div>
                  </>
                ) : (
                  value
                )}
              </div>
            </div>
          ))}
        </div>
      </div>

      <div>
        <div className={classes.subHeader}>Permission dates</div>
        <div className={classes.cardsWrapper}>
          {Object.entries(permissionDates).map(([key, value]: any) => (
            <div className={classes.card} key={key}>
              <div className={classes.cardTitle}>{key}</div>
              <div className={classes.cardContent}>{value}</div>
            </div>
          ))}
        </div>
      </div>

      <div>
        <div className={classes.subHeader}>Details being shared</div>
        <div>
          <div className={classes.detailsTitle}>Your Regular Payments</div>
          <ul className={classes.ulList}>
            <li>Your direct debits</li>
            <li>Your standing orders</li>
            <li>Your scheduled payments</li>
          </ul>
        </div>
      </div>
    </ApplicationAccessDrawer>
  );
}

export default ApplicationAccessPaymentDrawer;
