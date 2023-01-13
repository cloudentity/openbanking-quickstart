import React, { useState } from "react";
import { makeStyles } from "tss-react/mui";
import Grid from "@mui/material/Grid";
import Accounts from "./Accounts";
import { useQuery } from "react-query";
import { api } from "../api/api";
import Progress from "./Progress";
import { applyFiltering } from "./analytics.utils";
import Analytics from "./Analytics";
import { getCurrency } from "./utils";
import {
  AccountsResponse,
  BalancesResponse,
  Filter,
  TransactionsResponse,
} from "./types";

const useStyles = makeStyles()(() => ({
  root: {
    height: "100%;",
  },
}));

interface Props {
  banks: string[];
  onConnectClick: () => void;
  onDisconnect: (bankId: string) => () => void;
  onReconnect: (bankId: string, permissions: string[]) => () => void;
}

export default function Connected({
  banks,
  onConnectClick,
  onDisconnect,
  onReconnect,
}: Props) {
  const { classes } = useStyles();
  const [filtering, setFiltering] = useState<Filter>({
    accounts: [],
    months: [],
    categories: [],
  });

  const {
    isLoading: fetchAccountsProgress,
    error: fetchAccountsError,
    data: accountsRes,
  } = useQuery<AccountsResponse, any>("fetchAccounts", api.fetchAccounts, {
    refetchOnWindowFocus: false,
    retry: false,
    onSuccess: data => {
      setFiltering(m => ({
        ...m,
        accounts: (data.accounts || []).map(a => a.AccountId),
      }));
    },
  });

  const { isLoading: fetchBalancesProgress, data: balancesRes } =
    useQuery<BalancesResponse>("fetchBalances", api.fetchBalances, {
      refetchOnWindowFocus: false,
      retry: false,
    });

  const { isLoading: fetchTransactionsProgress, data: transactionsRes } =
    useQuery<TransactionsResponse>("fetchTransactions", api.fetchTransactions, {
      refetchOnWindowFocus: false,
      retry: false,
    });

  const accounts = accountsRes?.accounts || [];
  const balances = balancesRes?.balances || [];

  const transactions = applyFiltering(
    filtering,
    transactionsRes?.transactions ?? []
  );

  const isLoading =
    fetchAccountsProgress || fetchBalancesProgress || fetchTransactionsProgress;

  const bankNeedsReconnect =
    fetchAccountsError?.response?.error?.status === 401;

  const currencyType = getCurrency(balances[0]?.Currency);

  if (isLoading) {
    return <Progress />;
  }

  return (
    <Grid container className={classes.root}>
      <Grid
        item
        xs={4}
        style={{
          background: "#F7FAFF",
          padding: "16px 32px",
          borderRight: "1px solid #EAECF1",
        }}
      >
        <Accounts
          banks={banks}
          reconnectBank={bankNeedsReconnect}
          accounts={accounts}
          balances={balances}
          filtering={filtering}
          onChangeFiltering={(f: Filter) =>
            setFiltering({ ...filtering, ...f })
          }
          onConnectClick={onConnectClick}
          onDisconnect={onDisconnect}
          onReconnect={onReconnect}
        />
      </Grid>
      <Grid
        item
        xs={8}
        style={{ background: "#FCFCFF", padding: "32px 32px 16px 32px" }}
      >
        <Analytics
          currencyType={currencyType}
          transactions={transactions}
          filtering={filtering}
          onChangeFiltering={(f: Filter) =>
            setFiltering({ ...filtering, ...f })
          }
        />
      </Grid>
    </Grid>
  );
}
