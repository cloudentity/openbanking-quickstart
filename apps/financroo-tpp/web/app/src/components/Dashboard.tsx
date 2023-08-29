import React, { useEffect, useState } from "react";
import PageToolbar from "./common/PageToolbar";
import Connected from "./Connected";
import Welcome from "./Welcome";
import ConnectAccount from "./ConnectAccount";
import { useQuery } from "react-query";
import { api } from "../api/api";
import Progress from "./Progress";
import PageContent from "./common/PageContent";
import PageContainer from "./common/PageContainer";
import { useLocation, useNavigate } from "react-router-dom";
import Snackbar from "@mui/material/Snackbar";
import IconButton from "@mui/material/IconButton";
import CloseIcon from "@mui/icons-material/Close";
import Alert from "@mui/material/Alert";
import { makeStyles } from "tss-react/mui";
import AccountsAddedDialog from "./AccountsAddedDialog";
import { BanksResponse } from "./types";

const useStyles = makeStyles()(() => ({
  alert: {
    width: "100%",
    "& > div:last-of-type": {
      position: "relative",
      top: 2,
      fontWeight: 600,
    },
  },
}));

export default function Dashboard() {
  const { classes } = useStyles();

  const [connectAccountOpen, setConnectAccountOpen] = useState(false);
  const [isProgress, setProgress] = useState(false);
  const [snackbar, setSnackbar] = useState("");
  const [accountAddedDialog, setAccountAddedDialog] = useState(false);

  const navigate = useNavigate();
  const location = useLocation();
  const queryParams = new URLSearchParams(useLocation().search);
  const state = location.state as { bankNeedsReconnect: boolean } | undefined;

  useEffect(() => {
    if (state?.bankNeedsReconnect) {
      setSnackbar("Error: unauthorized. Bank needs reconnect");
      navigate(location, {
        state: { bankNeedsReconnect: false },
        replace: true,
      });
    }
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [state]);

  const {
    isLoading: fetchBanksProgress,
    data: banksRes,
    refetch: refetchBanks,
  } = useQuery<BanksResponse>("fetchBanks", api.fetchBanks, {
    refetchOnWindowFocus: false,
    retry: false,
  });

  const connectedBanks = banksRes?.connected_banks ?? [];

  useEffect(() => {
    if (queryParams.get("connected") === "yes" && window.featureFlags?.Investments) {
      setAccountAddedDialog(true);
    }
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

  useEffect(() => {
    if (accountAddedDialog === false) {
      queryParams.delete("connected");
      navigate(
        { pathname: location.pathname, search: queryParams.toString() },
        { replace: true }
      );
    }
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [accountAddedDialog]);

  const handleAllowAccess = (
    bankId: string,
    permissions: string[] | undefined
  ) => {
    setProgress(true);
    api
      .connectBank(bankId, { permissions })
      .then(res => {
        window.location.href = res.login_url;
      })
      .catch(() => setProgress(false));
  };

  const handleDisconnectBank = (bankId: string) => () => {
    setProgress(true);
    api
      .disconnectBank(bankId)
      .then(refetchBanks)
      .finally(() => setProgress(false));
  };

  const handleReconnectBank =
    (bankId: string, permissions: string[] | undefined) => () => {
      setProgress(true);
      api
        .connectBank(bankId, { permissions })
        .then(res => {
          window.location.href = res.login_url;
        })
        .catch(() => setProgress(false));
    };

  const showProgress = isProgress || fetchBanksProgress;

  return (
    <div style={{ position: "relative" }}>
      <PageToolbar mode="main" />

      {showProgress && <Progress />}

      {!showProgress && (
        <>
          {connectedBanks.length === 0 ? (
            <PageContainer withBackground>
              <Welcome onConnectClick={() => setConnectAccountOpen(true)} />
            </PageContainer>
          ) : (
            <PageContent>
              <Connected
                banks={connectedBanks}
                onConnectClick={() => setConnectAccountOpen(true)}
                onDisconnect={handleDisconnectBank}
                onReconnect={handleReconnectBank}
              />
            </PageContent>
          )}
        </>
      )}

      {connectAccountOpen && (
        <ConnectAccount
          banks={banksRes}
          onAllowAccess={handleAllowAccess}
          onClose={() => setConnectAccountOpen(false)}
        />
      )}

      <Snackbar
        anchorOrigin={{
          vertical: "top",
          horizontal: "center",
        }}
        open={!!snackbar}
        autoHideDuration={6000}
        onClose={() => setSnackbar("")}
        action={
          <>
            <IconButton
              size="small"
              aria-label="close"
              color="inherit"
              onClick={() => setSnackbar("")}
            >
              <CloseIcon fontSize="small" />
            </IconButton>
          </>
        }
      >
        <Alert severity="error" className={classes.alert}>
          {snackbar}
        </Alert>
      </Snackbar>

      <AccountsAddedDialog
        open={accountAddedDialog === null ? false : accountAddedDialog}
        setOpen={setAccountAddedDialog}
      />
    </div>
  );
}
