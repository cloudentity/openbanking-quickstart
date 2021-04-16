import React from "react";
import { makeStyles } from "@material-ui/core/styles";
import Chip from "@material-ui/core/Chip";
import clsx from "clsx";

import ContributionCard from "./ContributionCard";
import Field from "./Field";
import { theme } from "../../theme";
import { BalanceType, AccountType } from "./InvestmentsContribute";
import { banks } from "../banks";

const useStyles = makeStyles((theme) => ({
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
}));

type Props = {
  amount: string;
  bank: string;
  account: string;
  balances: BalanceType[];
  handleBack: () => void;
  handleNext: () => void;
  accounts: AccountType[];
};

export default function InvestmentsContributeSummary({
  amount,
  bank,
  account,
  balances,
  handleBack,
  handleNext,
  accounts,
}: Props) {
  const classes = useStyles();

  const selectedBalance = balances.find((a) => a.AccountId === account);
  const selectedBank = banks.find((a) => a.value === bank);
  const selectedAccountInfo = accounts.find((a) => a.AccountId === account);

  return (
    <ContributionCard
      title={<div className={classes.title}>INVESTMENT SUMMARY</div>}
      backButton={{ title: "Back", onClick: handleBack }}
      nextButton={{ title: "Confirm", onClick: handleNext }}
    >
      <Field>
        <div
          className={clsx([classes.information, classes.informationOneRow])}
          style={{ alignItems: "center", paddingBottom: 20 }}
        >
          <div style={{ ...theme.custom.heading6 }}>PAYMENT TOTAL</div>
          <div>
            <Chip label={`£ ${amount}`} className={classes.chip} />
          </div>
        </div>
      </Field>
      <Field style={{ ...theme.custom.caption }}>
        To consent to this transaction, confirm the details below
      </Field>
      <Field label="Payee Information">
        <div className={classes.information}>
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
        <div className={clsx([classes.information, classes.informationOneRow])}>
          <div style={{ ...theme.custom.heading6 }}>Bank Name</div>
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
