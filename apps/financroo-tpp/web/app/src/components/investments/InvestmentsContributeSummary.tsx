import React from "react";
import { makeStyles } from "tss-react/mui";
import Chip from "@mui/material/Chip";

import ContributionCard from "./ContributionCard";
import Field from "./Field";
import { theme } from "../../theme";
import { banks } from "../banks";
import { Account, Balance } from "../types";
import { getCurrency } from "../utils";

const useStyles = makeStyles()(theme => ({
  title: {
    ...theme.custom.heading6,
  },
  chip: {
    backgroundColor: theme.palette.primary.main,
    color: "white",
    fontWeight: 600,
  },
  grid: {
    display: "grid",
    gridTemplateColumns: "1fr 1fr 1fr",
    gridColumnGap: 14,
  },
  card: {
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
  informationOneRow: {
    alignItems: "center",
    paddingBottom: 20,
    "& > div": {
      paddingBottom: 0,
      alignItems: "center",
    },
  },
  bankLogoImage: {
    width: 15,
    marginLeft: 20,
  },
  bankRow: {
    display: "flex",
    justifyContent: "flex-end",
  },
  heading: {
    ...theme.custom.heading6,
  },
  caption: {
    ...theme.custom.caption,
  },
}));

interface Props {
  amount: string;
  currency: string | undefined;
  selectedBankId: string;
  selectedAccountId: string;
  balances: Balance[];
  accounts: Account[];
  handleBack: () => void;
  handleNext: () => void;
}

export default function InvestmentsContributeSummary({
  amount,
  currency,
  selectedBankId,
  selectedAccountId,
  balances,
  handleBack,
  handleNext,
  accounts,
}: Props) {
  const { cx, classes } = useStyles();

  const selectedBalance = balances.find(a => a.AccountId === selectedAccountId);
  const selectedBank = banks.find(a => a.value === selectedBankId);
  const selectedAccountInfo = accounts.find(
    a => a.AccountId === selectedAccountId
  );

  return (
    <ContributionCard
      title={<div className={classes.title}>INVESTMENT SUMMARY</div>}
      backButton={{ title: "Back", onClick: handleBack }}
      nextButton={{ title: "Confirm", onClick: handleNext }}
    >
      <Field>
        <div
          className={cx(classes.information, classes.informationOneRow)}
          style={{ alignItems: "center", paddingBottom: 20 }}
        >
          <div className={classes.heading}>PAYMENT TOTAL</div>
          <div>
            <Chip
              id="total-amount"
              label={`${getCurrency(currency)} ${parseFloat(amount).toFixed(
                2
              )}`}
              className={classes.chip}
            />
          </div>
        </div>
      </Field>
      <Field style={theme.custom.caption}>
        To consent to this transaction, confirm the details below
      </Field>
      <Field label="Payee Information">
        <div className={classes.information} id={`account-id-${selectedBalance?.AccountId}`}>
          <div>Payee Account Name</div>
          <div>{selectedAccountInfo?.Account[0].Name}</div>
          <div>Sort code</div>
          <div>{selectedAccountInfo?.Account[0].Identification}</div>
          <div>Account number</div>
          <div>**** ***** **** {selectedBalance?.AccountId}</div>
          <div>Payment reference</div>
          <div>Financoo investments Ltd</div>
        </div>
      </Field>
      <Field label="Payment Information">
        <div className={cx(classes.information, classes.informationOneRow)}>
          <div className={classes.heading}>Bank Name</div>
          <div className={classes.bankRow}>
            {selectedBank?.name}
            <img
              src={selectedBank?.icon}
              alt="bank logo"
              className={classes.bankLogoImage}
            />
          </div>
        </div>
      </Field>
      <Field style={{ ...theme.custom.caption, marginBottom: 0 }}>
        You will be securely transferred to <strong>Go Bank</strong> to
        authorize the payment
      </Field>
    </ContributionCard>
  );
}
