import React, { useEffect } from "react";
import Chip from "@mui/material/Chip";
import Select from "@mui/material/Select";
import MenuItem from "@mui/material/MenuItem";
import Radio from "@mui/material/Radio";
import Alert from "@mui/material/Alert";
import { makeStyles } from "tss-react/mui";

import ContributionCard from "./ContributionCard";
import Field from "./Field";
import bankIcon from "../../assets/icon-bank.svg";
import cardIcon from "../../assets/icon-credit-card.svg";
import paypalIcon from "../../assets/icon-paypal.svg";
import walletIcon from "../../assets/icon-wallet.svg";
import { Bank } from "../banks";
import { Account, Balance } from "../types";
import { getCurrency } from "../utils";

const useStyles = makeStyles()(theme => ({
  titleContainer: {
    display: "flex",
    alignItems: "center",
    justifyContent: "space-between",
    paddingLeft: "32px",
    paddingRight: "32px",
  },
  title: {
    fontWeight: "bold",
    fontSize: 12,
    lineHeight: "16px",
    textTransform: "uppercase",
    color: "#626576",
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
      marginBottom: 8,
    },
    "& > span": {
      ...theme.custom.caption,
    },
  },
  disabled: {
    border: "1px solid #E8EAED",
    opacity: 0.6,
    boxShadow: "none",
  },
  radioGroup: {
    display: "flex",
    "& > div": {
      fontWeight: 600,
      fontSize: 12,
      lineHeight: "24px",
      color: "#212533",
    },
  },
  accountSelect: {
    "& :first-child": {
      borderTopLeftRadius: 4,
      borderTopRightRadius: 4,
    },
    "& :last-child": {
      borderBottomLeftRadius: 4,
      borderBottomRightRadius: 4,
    },
  },
  accountSelectItem: {
    display: "flex",
    alignItems: "center",
    border: "solid 1px #ECECEC",
    padding: "10px 20px 10px 11px",
    "& > img": {
      paddingLeft: 3,
      paddingRight: 12,
    },
  },
  accountSelectItemLabel: {
    "& :first-child": {
      fontWeight: 600,
      fontSize: 12,
      lineHeight: "24px",
      color: "#212533",
    },
    "& :last-child": {
      fontSize: 12,
      lineHeight: "22px",
      color: "#A0A3B5",
    },
  },
  active: {
    border: `solid 1px ${theme.palette.primary.main}`,
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
  alert: {
    width: "100%",
    "& > div:last-of-type": {
      position: "relative",
      top: 2,
    },
  },
  heading: {
    ...theme.custom.heading6,
  },
}));

interface Props {
  amount: string;
  currency: string | undefined;
  handleBack: () => void;
  handleNext: () => void;
  selectedBankId: string;
  setSelectedBankId: (bankId: string) => void;
  selectedAccountId: string;
  setSelectedAccountId: (accountId: string) => void;
  banks: Bank[];
  alert: string;
  setAlert: (message: string) => void;
  balances: Balance[];
  accounts: Account[];
}

export default function InvestmentsContributeMethod({
  amount,
  currency,
  handleBack,
  handleNext,
  selectedBankId,
  setSelectedBankId,
  selectedAccountId,
  setSelectedAccountId,
  alert,
  setAlert,
  balances,
  accounts,
  banks,
}: Props) {
  const { cx, classes } = useStyles();

  const selectedAccount = balances.find(
    balance => balance.AccountId === selectedAccountId
  );

  const selectedAccountInfo = accounts.find(
    account => account.AccountId === selectedAccountId
  );

  useEffect(() => {
    if (selectedAccount) {
      if (Number(amount) <= Number(selectedAccount.Amount)) {
        setAlert("");
      } else {
        setAlert("Payment amount exceeds account balance");
      }
    }
  }, [amount, selectedAccount, setAlert]);

  return (
    <ContributionCard
      title={
        <div className={classes.titleContainer}>
          <div className={classes.heading}>PAYMENT TOTAL</div>
          <Chip
            id="total-amount"
            label={`${getCurrency(currency)} ${parseFloat(amount).toFixed(2)}`}
            className={classes.chip}
          />
        </div>
      }
      backButton={{ title: "Back", onClick: handleBack }}
      nextButton={{
        title: "Next",
        onClick: () => {
          if (!alert) {
            handleNext();
          }
        },
        disabled: !!alert,
      }}
    >
      <Field label="Select payment method">
        <div className={classes.grid}>
          <div className={classes.card}>
            <img src={bankIcon} alt="bank icon" />
            <span>Bank Transfer</span>
          </div>
          <div className={cx(classes.card, classes.disabled)}>
            <img src={cardIcon} alt="card icon" />
            <span>Credit / Debit card</span>
          </div>
          <div className={cx(classes.card, classes.disabled)}>
            <img src={paypalIcon} alt="bank icon" />
            <span>Paypal Transfer</span>
          </div>
        </div>
      </Field>
      <Field
        label="Select Bank"
        helperText="Paying with your bank is completely safe and secure with Open Banking"
      >
        <Select
          value={selectedBankId}
          onChange={v => setSelectedBankId(v.target.value)}
          style={{ width: "100%" }}
          variant="outlined"
        >
          {banks.map(({ name, value }) => (
            <MenuItem value={value} key={value}>
              {name || value}
            </MenuItem>
          ))}
        </Select>
      </Field>
      <Field>
        <div className={classes.radioGroup}>
          <div style={{ marginRight: 16 }}>
            <Radio checked color="primary" />
            Select Available Account
          </div>
          <div className={classes.disabled} style={{ border: "none" }}>
            <Radio disabled />
            Add Bank details
          </div>
        </div>
      </Field>

      <Field>
        <div id="accounts-list" className={classes.accountSelect}>
          {balances
            .filter(b => b.BankId === selectedBankId)
            .map(({ AccountId, Amount }) => (
              <div
                key={AccountId}
                id={`account-id-${AccountId}`}
                className={cx(
                  classes.accountSelectItem,
                  selectedAccountId === AccountId && classes.active
                )}
              >
                <Radio
                  checked={selectedAccountId === AccountId}
                  color="primary"
                  onChange={e => {
                    if (e.target.checked) {
                      setSelectedAccountId(AccountId);
                    }
                  }}
                />
                <img src={walletIcon} alt="wallet icon" />
                <div className={classes.accountSelectItemLabel}>
                  <div>Checking account</div>
                  <div>**** ***** **** {AccountId} </div>
                </div>
                <div style={{ flex: 1, textAlign: "right" }}>
                  {getCurrency(currency)} <>{parseFloat(Amount).toFixed(2)}</>
                </div>
              </div>
            ))}
        </div>
      </Field>
      <Field label="Payee Information" style={alert ? {} : { marginBottom: 0 }}>
        <div className={classes.information}>
          <div>Payee Account Name</div>
          <div>{selectedAccountInfo?.Account[0].Name}</div>
          <div>Sort code</div>
          <div>{selectedAccountInfo?.Account[0].Identification}</div>
          <div>Account number</div>
          <div>**** ***** **** {selectedAccount?.AccountId}</div>
          <div>Payment reference</div>
          <div>Financoo investments Ltd</div>
        </div>
      </Field>
      {alert && (
        <Alert severity="error" className={classes.alert}>
          {alert}
        </Alert>
      )}
    </ContributionCard>
  );
}
