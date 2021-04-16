import React, {useState} from "react";
import {Theme} from "@material-ui/core";
import {makeStyles} from "@material-ui/core/styles";
import Grid from "@material-ui/core/Grid";
import Accounts from "./Accounts";
import Analytics from "./Analytics";
import {useQuery} from "react-query";
import {api} from "../api/api";
import Progress from "./Progress";
import {applyFiltering} from "./analytics.utils";
import {path, pick} from "ramda";

const useStyles = makeStyles((theme: Theme) => ({
  root: {
    height: '100%;'
  }
}));

export default function Connected({banks, onConnectClick, onDisconnect, onReconnect}) {
  const classes = useStyles();
  const [filtering, setFiltering] = useState({
    accounts: [],
    months: [],
    categories: []
  });

  const {
    isLoading: fetchAccountsProgress,
    error: fetchAccountsError,
    data: accountsRes
  } = useQuery('fetchAccounts', api.fetchAccounts, {
    refetchOnWindowFocus: false,
    retry: false,
    onSuccess: data => {
      setFiltering(m => ({...m, accounts: (data.accounts || []).map(a => a.AccountId)}));
    }
  });

  const {isLoading: fetchBalancesProgress, data: balancesRes} = useQuery('fetchBalances', api.fetchBalances, {
    refetchOnWindowFocus: false,
    retry: false
  });

  const {isLoading: fetchTransactionsProgress, data: transactionsRes} = useQuery('fetchTransactions', api.fetchTransactions, {
    refetchOnWindowFocus: false,
    retry: false
  });

  const accounts = accountsRes?.accounts || [];
  const balances = balancesRes?.balances || [];
  const transactions = applyFiltering(pick(['accounts'], filtering), transactionsRes?.transactions || []);

  const isLoading = fetchAccountsProgress || fetchBalancesProgress || fetchTransactionsProgress;

  const bankNeedsReconnect = path(['response', 'error', 'status'], fetchAccountsError) === 401;

  if (isLoading) {
    return <Progress/>;
  }

  return (
    <Grid container className={classes.root}>
      <Grid item xs={4} style={{background: '#F7FAFF', padding: '16px 32px', borderRight: '1px solid #EAECF1'}}>
        <Accounts
          banks={banks}
          reconnectBank={bankNeedsReconnect}
          accounts={accounts}
          balances={balances}
          filtering={filtering}
          onChangeFiltering={f => setFiltering({...filtering, ...f})}
          onConnectClick={onConnectClick}
          onDisconnect={onDisconnect}
          onReconnect={onReconnect}
        />
      </Grid>
      <Grid item xs={8} style={{background: '#FCFCFF', padding: '32px 32px 16px 32px'}}>
        <Analytics transactions={transactions} filtering={filtering} onChangeFiltering={f => setFiltering({...filtering, ...f})}/>
      </Grid>
    </Grid>
  )
};
