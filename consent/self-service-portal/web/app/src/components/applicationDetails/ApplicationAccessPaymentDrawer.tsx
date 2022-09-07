import React from "react";
import { makeStyles } from "tss-react/mui";

import Chip from "../Chip";
import logo from "../../assets/welcome-image.png";
import ApplicationAccessDrawer from "./ApplicationAccessDrawer";
import { getDate } from "../ApplicationSimpleCard";
import { getCurrency, drawerStyles } from "./utils";
import { Consent } from "../types";

const useStyles = makeStyles()(() => ({
  ...drawerStyles,
  cardsWrapper: {
    display: "flex",
    flexWrap: "wrap",
  },
}));

interface Props {
  drawerData: Consent;
  setDrawerData: (data: Consent | undefined) => void;
  status: string;
}

function ApplicationAccessPaymentDrawer({
  drawerData,
  setDrawerData,
  status,
}: Props) {
  const { classes } = useStyles();

  const transactionDetails = {
    Amount: `${getCurrency(drawerData?.Currency)} ${drawerData?.Amount}`,
    Status: drawerData?.Status,
    "Consent id": drawerData?.ConsentID,
    "Debtor Account": {
      id: drawerData.DebtorAccountIdentification,
      name: drawerData.DebtorAccountName,
    },
    "Creditor Account": {
      id: drawerData.CreditorAccountIdentification,
      name: drawerData.CreditorAccountName,
    },
  };

  const permissionDates = {
    Authorised: getDate(drawerData?.CompletionDateTime),
    "Last updated": getDate(drawerData?.StatusUpdateDateTime),
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
        <div className={`${classes.purposeHeader} purpose-header`}>
          Purpose for sharing data:
        </div>
        <div>To enable payments to Financroo investments</div>
      </div>

      <div>
        <div className={classes.subHeader}>TRANSACTION Details</div>
        <div className={classes.cardsWrapper} id="transactions-details">
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
