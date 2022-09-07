import React from "react";
import { makeStyles } from "tss-react/mui";
import Avatar from "@mui/material/Avatar";

import CustomDrawer from "./CustomDrawer";
import { getCurrency, drawerStyles, getDate, Consent } from "../../utils";
import Chip from "../../Chip";

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
}

function PaymentDrawer({ drawerData, setDrawerData }: Props) {
  const { classes } = useStyles();

  const transactionDetails = {
    Amount: `${getCurrency(drawerData.currency)} ${drawerData.Amount}`,
    Status: drawerData.status,
    "Consent id": drawerData.consent_id,
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
    Authorised: getDate(drawerData.completed_at),
    "Last updated": getDate(drawerData.updated_at),
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
            {drawerData.CreditorAccountName?.length
              ? drawerData.CreditorAccountName[0].toUpperCase()
              : ""}
          </Avatar>
          <h3 className={classes.name}>Financroo</h3>
          <div style={{ flex: 1 }} />
          <Chip type={status && (status.toLowerCase() as any)}>{status}</Chip>
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
          {Object.entries(transactionDetails).map(([key, value]) => (
            <div className={classes.card} key={key}>
              <div className={classes.cardTitle}>{key}</div>
              <div className={classes.cardContent}>
                {typeof value === "object" && value?.id ? (
                  <>
                    <div>{value.id}</div>
                    <div>{value.name}</div>
                  </>
                ) : (
                  <>{value}</>
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
