import React, { useEffect, useMemo, useState } from "react";
import IconButton from "@mui/material/IconButton";
import { makeStyles } from "tss-react/mui";
import { ArrowLeft, Lock } from "react-feather";
import { useNavigate } from "react-router-dom";
import { useQuery } from "react-query";

import PageContainer from "../common/PageContainer";
import PageToolbar, { subHeaderHeight } from "../common/PageToolbar";
import InvestmentsContributeAmount from "./InvestmentsContributeAmount";
import InvestmentsContributeMethod from "./InvestmentsContributeMethod";
import InvestmentsContributeSummary from "./InvestmentsContributeSummary";
import { api } from "../../api/api";
import Progress from "../Progress";
import { banks as banksArray } from "../banks";
import InvestmentsContributeRedirecting from "./InvestmentsContributeRedirecting";
import { AccountsResponse, BalancesResponse, BanksResponse } from "../types";

const useStyles = makeStyles()(theme => ({
  toolbarButton: {
    color: theme.palette.primary.main,
  },
  title: {
    display: "flex",
    alignItems: "center",
  },
  footer: {
    padding: 34,
    display: "flex",
    alignItems: "center",
    ...theme.custom.caption,
  },
  spacer: {
    flex: 1,
  },
}));

const stepsTitle = {
  0: "Contributions",
  1: "Select the payment method",
  2: "Investment summary",
};

export default function InvestmentsContribute() {
  const { classes } = useStyles();
  const navigate = useNavigate();

  const [step, setStep] = useState(0);
  const [amount, setAmount] = useState("");
  const [alert, setAlert] = useState("");
  const [isProgress, setProgress] = useState(false);
  const [selectedBankId, setSelectedBankId] = useState("");
  const [selectedAccountId, setSelectedAccountId] = useState("");

  const { isLoading: fetchBanksProgress, data: banksRes } =
    useQuery<BanksResponse>("fetchBanks", api.fetchBanks, {
      refetchOnWindowFocus: false,
      retry: false,
    });

  const { isLoading: fetchBalancesProgress, data: balancesRes } =
    useQuery<BalancesResponse>("fetchBalances", api.fetchBalances, {
      refetchOnWindowFocus: false,
      retry: true,
    });

  const { isLoading: fetchAccountsProgress, data: accountsRes } =
    useQuery<AccountsResponse>("fetchAccounts", api.fetchAccounts, {
      refetchOnWindowFocus: false,
      retry: true,
    });

  const balances = balancesRes?.balances ?? [];
  const accounts = useMemo(() => accountsRes?.accounts ?? [], [accountsRes]);
  const connectedBanks = useMemo(
    () => banksRes?.connected_banks ?? [],
    [banksRes]
  );
  const banks = useMemo(
    () => banksArray.filter(({ value }) => connectedBanks.includes(value)),
    [connectedBanks]
  );

  useEffect(() => {
    if (banks.length) {
      setSelectedBankId(banks[0].value);
    }
  }, [banks]);

  useEffect(() => {
    if (accounts.length) {
      setSelectedAccountId(accounts[0].AccountId);
    }
  }, [accounts]);

  const accountId = balances.length ? balances[0].AccountId : undefined;
  const selectedBalance = balances.find(
    balance => balance.AccountId === accountId
  );
  const accountDetails = selectedBalance
    ? { amount: selectedBalance.Amount, currency: selectedBalance.Currency }
    : undefined;

  useEffect(() => {
    if (step === -1) {
      navigate("/investments");
      setStep(0);
      setAmount("");
    }
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [step, amount]);

  function handleBack() {
    setStep(step => step - 1);
  }

  function handleNext() {
    if (step === 3) {
      setProgress(true);

      api
        .domesticPaymentConsent({
          amount: amount,
          bank_id: selectedBankId,
          account_id: accountId,
          payee_account_name: "financroo-investment",
          payee_account_number: "12345678",
          payee_account_sort_code: "123456",
          payment_reference: "financroo-investment-123",
        })
        .then(res => {
          window.location.href = res.login_url;
        })
        .finally(() => {
          setProgress(false);
        });
    } else {
      setStep(step => step + 1);
    }
  }

  const showProgress =
    isProgress ||
    fetchBanksProgress ||
    fetchBalancesProgress ||
    fetchAccountsProgress;

  return (
    <div style={{ position: "relative" }}>
      <PageToolbar
        mode="onlySubheader"
        subHeaderTitle={
          <div className={classes.title}>
            <IconButton
              className={classes.toolbarButton}
              onClick={() => {
                setStep(step => step - 1);
              }}
              size="large"
            >
              <ArrowLeft style={{ color: "#36C6AF" }} />
            </IconButton>
            <span>{stepsTitle[step]}</span>
          </div>
        }
      />
      {showProgress ? (
        <Progress top={150} />
      ) : (
        <PageContainer
          withOnlySubheader
          style={{
            paddingTop: 48,
            marginTop: subHeaderHeight,
            position: "relative",
          }}
          containerStyle={{
            display: "flex",
            flexDirection: "column",
            alignItems: "center",
            justifyContent: "center",
            minHeight: "100%",
          }}
        >
          {step === 0 && (
            <InvestmentsContributeAmount
              amount={amount}
              setAmount={setAmount}
              handleBack={handleBack}
              handleNext={handleNext}
              accountDetails={accountDetails}
              setAlert={setAlert}
            />
          )}
          {step === 1 && (
            <InvestmentsContributeMethod
              amount={amount}
              currency={accountDetails?.currency}
              handleBack={handleBack}
              handleNext={handleNext}
              selectedBankId={selectedBankId}
              setSelectedBankId={setSelectedBankId}
              selectedAccountId={selectedAccountId}
              setSelectedAccountId={setSelectedAccountId}
              balances={balances}
              accounts={accounts}
              banks={banks}
              alert={alert}
              setAlert={setAlert}
            />
          )}
          {step === 2 && (
            <InvestmentsContributeSummary
              amount={amount}
              currency={accountDetails?.currency}
              handleBack={handleBack}
              handleNext={handleNext}
              selectedBankId={selectedBankId}
              selectedAccountId={selectedAccountId}
              balances={balances}
              accounts={accounts}
            />
          )}
          {step === 3 && (
            <InvestmentsContributeRedirecting handleNext={handleNext} />
          )}
          <div className={classes.spacer} />
          <div className={classes.footer}>
            <Lock style={{ marginRight: 12 }} />
            We use multi-level ecryption measures to protect your data
          </div>
        </PageContainer>
      )}
    </div>
  );
}
