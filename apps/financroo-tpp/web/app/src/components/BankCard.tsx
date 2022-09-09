import React, { CSSProperties } from "react";
import Button from "@mui/material/Button";
import { makeStyles } from "tss-react/mui";
import Typography from "@mui/material/Typography";
import IconButton from "@mui/material/IconButton";
import { Plus } from "react-feather";
import Card from "@mui/material/Card";
import Checkbox from "@mui/material/Checkbox";
import { banks } from "./banks";
import { getCurrency } from "./utils";
import { Account, Balance, Filter } from "./types";

const useStyles = makeStyles()(() => ({
  accountRoot: {
    borderBottom: "1px solid #ECECEC",
    "&:hover": {
      cursor: "pointer",
    },
  },
  header: {
    padding: 20,
    display: "flex",
    alignItems: "center",
    borderBottom: "1px solid #ECECEC",
  },
  iconContainer: {
    background: "#FCFCFF",
    borderRadius: "50%",
    width: 52,
    height: 52,
    display: "flex",
    alignItems: "center",
    justifyContent: "center",
    boxShadow:
      "0px 0.574468px 0.574468px rgba(0, 0, 0, 0.08), 0px 0px 0.574468px rgba(0, 0, 0, 0.31)",
  },
  footer: {
    height: 52,
    padding: "0 21px",
    background: "rgba(54, 198, 175, 0.08)",
    color: "#36C6AF",
    display: "flex",
    alignItems: "center",
    justifyContent: "space-between",
  },
}));

interface Props {
  bankId: string;
  reconnectBank: boolean;
  accounts: Account[];
  balances: Balance[];
  filtering: Filter;
  onChangeFiltering: (filter: Filter) => void;
  onDisconnect: (bankId: string) => () => void;
  onReconnect: (bankId: string, permissions: string[]) => () => void;
  style?: CSSProperties;
}

export default function BankCard({
  bankId,
  reconnectBank,
  accounts,
  balances,
  filtering,
  onChangeFiltering,
  onDisconnect,
  onReconnect,
  style = {},
}: Props) {
  const { classes } = useStyles();

  const getAccountBalance = (accountId, balances) =>
    balances.find(b => b.AccountId === accountId);
  const getAccountAmountAsString = (accountId, balances) => {
    const accountBalance = getAccountBalance(accountId, balances);

    return accountBalance
      ? `${getCurrency(accountBalance.Currency)} ${accountBalance.Amount ?? 0}`
      : "N/A";
  };
  const isAccountChecked = id => filtering?.accounts?.includes(id);

  let selectedBank = banks.find(b => b.value === bankId);

  return (
    <Card style={style} id={bankId}>
      <div className={classes.header}>
        <div className={classes.iconContainer}>
          <img
            src={selectedBank?.icon}
            style={{ width: 24, height: 24 }}
            alt="bank icon"
          />
        </div>
        <div style={{ marginLeft: 24 }}>
          <Typography>{banks.find(b => b.value === bankId)?.name}</Typography>
          <Typography
            style={{
              background: "rgba(54, 198, 175, 0.08)",
              color: "#36C6AF",
              fontSize: 14,
              padding: 2,
              marginTop: 4,
            }}
            id="accounts-synced-number"
          >
            {accounts.length} accounts synced
          </Typography>
        </div>
        <div style={{ flex: 1 }} />
        <div>
          {reconnectBank && (
            <Button
              size="small"
              className="reconnect-button"
              variant="contained"
              color="primary"
              id="access-bank-button"
              style={{ color: "#fff" }}
              onClick={onReconnect(
                bankId,
                selectedBank?.permissions?.map(p => p.value).filter(p => p) ??
                  []
              )}
            >
              reconnect
            </Button>
          )}
          {!reconnectBank && (
            <Button
              size="small"
              className="disconnect-button"
              variant="outlined"
              id="access-bank-button"
              onClick={onDisconnect(bankId)}
            >
              disconnect
            </Button>
          )}
        </div>
      </div>
      {accounts.map(account => (
        <div
          key={account.AccountId}
          onClick={() =>
            onChangeFiltering({
              accounts: isAccountChecked(account.AccountId)
                ? filtering?.accounts?.filter(a => a !== account.AccountId)
                : [...(filtering.accounts ?? []), account.AccountId],
              months: [],
              categories: [],
            })
          }
          className={classes.accountRoot}
          style={{
            height: 62,
            background: isAccountChecked(account.AccountId)
              ? "#36C6AF"
              : "initial",
            color: isAccountChecked(account.AccountId) ? "#fff" : "initial",
            display: "flex",
            alignItems: "center",
            justifyContent: "space-between",
            padding: "0 20px",
          }}
        >
          <div style={{ display: "flex", alignItems: "center" }}>
            <Checkbox
              checked={isAccountChecked(account.AccountId)}
              onChange={() =>
                onChangeFiltering({
                  accounts: isAccountChecked(account.AccountId)
                    ? filtering.accounts?.filter(a => a !== account.AccountId)
                    : [...(filtering.accounts ?? []), account.AccountId],
                  months: [],
                  categories: [],
                })
              }
              color="primary"
              style={{
                color: isAccountChecked(account.AccountId) ? "#fff" : "initial",
              }}
              inputProps={{ "aria-label": "primary checkbox" }}
            />
            <div style={{ marginLeft: 12 }} id={`account-id-${account.AccountId}`}>
              <Typography className="account-name">
                {account.Nickname}
              </Typography>
              <Typography>**** ***** **** {account.AccountId}</Typography>
            </div>
          </div>
          <div>
            <Typography>
              {getAccountAmountAsString(account.AccountId, balances)}
            </Typography>
          </div>
        </div>
      ))}
      <div className={classes.footer}>
        <Typography>Add new account</Typography>
        <IconButton size="large">
          <Plus style={{ color: "#36C6AF" }} />
        </IconButton>
      </div>
    </Card>
  );
}
