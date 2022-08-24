import React from "react";
import { makeStyles } from "tss-react/mui";
import Avatar from "@mui/material/Avatar";

import CustomDrawer from "./CustomDrawer";
import { currencyDict, drawerStyles, getDate } from "../../utils";
import Chip from "../../Chip";

const useStyles = makeStyles()(() => ({
  ...drawerStyles,
  cardsWrapper: {
    display: "flex",
    flexWrap: "wrap",
  },
}));

interface Props {
  drawerData: any;
  setDrawerData: (data: string | null) => void;
}

function PaymentDrawer({ drawerData, setDrawerData }: Props) {
  const { classes } = useStyles();

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
  };

  const status = drawerData?.status;

  return (
    <CustomDrawer
      header={
        <div className={classes.headerContent}>
          <Avatar
            variant="square"
            className={classes.logo}
            style={{ backgroundColor: "white", color: "#626576" }}
          >
            {drawerData?.Client?.name[0]?.toUpperCase()}
          </Avatar>
          <h3 className={classes.name}>Financroo</h3>
          <div style={{ flex: 1 }} />
          <Chip type={status && status.toLowerCase()}>{status}</Chip>
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
    </CustomDrawer>
  );
}

export default PaymentDrawer;
