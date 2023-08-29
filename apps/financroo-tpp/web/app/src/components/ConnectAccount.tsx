import React, { useState } from "react";
import { makeStyles } from "tss-react/mui";
import Dialog from "@mui/material/Dialog";
import PageContainer from "./common/PageContainer";
import PageToolbar from "./common/PageToolbar";
import CloseIcon from "@mui/icons-material/Close";
import IconButton from "@mui/material/IconButton";
import Grid from "@mui/material/Grid";
import Typography from "@mui/material/Typography";
import Card from "@mui/material/Card";
import { banks as banksConfig, Bank, getUnknownBankConfig } from "./banks";
import Slide from "@mui/material/Slide";
import Button from "@mui/material/Button";
import connectArrows from "../assets/connect-arrows.svg";
import Paper from "@mui/material/Paper";
import Accordion from "@mui/material/Accordion";
import AccordionSummary from "@mui/material/AccordionSummary";
import AccordionDetails from "@mui/material/AccordionDetails";
import ExpandMoreIcon from "@mui/icons-material/ExpandMore";
import financrooIcon from "../assets/banks/financroo-icon.svg";
import { BanksResponse } from "./types";

const useStyles = makeStyles()(() => ({
  cardRoot: {
    height: 116,
    padding: "0 16px",
    display: "flex",
    "&:hover": {
      cursor: "pointer",
    },
  },
  disabled: {
    opacity: 0.6,
    "&:hover": {
      cursor: "initial",
    },
  },
  footer: {
    display: "flex",
    position: "fixed",
    justifyContent: "center",
    alignItems: "center",
    bottom: 0,
    height: 96,
    width: "100%",
    background: "#fff",
    borderTop: "1px solid #ECECEC",
  },
}));

interface Props {
  banks: BanksResponse | undefined;
  onAllowAccess: (bankId: string, permissions: string[] | undefined) => void;
  onClose: () => void;
}

export default function ConnectAccount({
  banks,
  onAllowAccess,
  onClose,
}: Props) {
  const { cx, classes } = useStyles();
  const [selected, setSelected] = useState<Bank | null>(null);

  const connectedBanks = banks?.connected_banks ?? [];
  const availableBanks = banks?.available_banks ?? [];

  return (
    <Dialog open={true} fullScreen>
      <PageToolbar mode="dialog">
        {!selected && (
          <IconButton
            edge="start"
            color="inherit"
            aria-label="close"
            onClick={onClose}
            size="large"
          >
            <CloseIcon />
          </IconButton>
        )}
        {selected && (
          <img alt="icon" src={selected.logo} style={{ maxHeight: 40 }} />
        )}
      </PageToolbar>
      <PageContainer style={{ paddingBottom: 112 }} withBackground>
        {!selected && (
          <Grid container justifyContent="center" style={{ marginTop: 64 }}>
            <Grid item xs={12} sm={8} md={6} style={{ textAlign: "center" }}>
              <Typography
                color="primary"
                variant="h2"
                style={{ marginTop: 24, fontSize: 28 }}
              >
                Connect your accounts
              </Typography>
              <Typography variant="body1" style={{ marginTop: 16 }}>
                By connecting your bank, bills and credit cards, you allow us to
                help you uncover insights that can improve your financial well
                being
              </Typography>
              <Grid container style={{ marginTop: 48 }} spacing={3}>
                {availableBanks.map(bank => {
                  const isConnected = connectedBanks.includes(bank.id);
                  const configBank =
                    banksConfig.find(b => b.value === bank.id) ||
                    getUnknownBankConfig(bank);

                  return (
                    <Grid item xs={6} sm={4} key={bank.id} id={bank.id}>
                      <Card
                        className={cx(
                          classes.cardRoot,
                          (isConnected) &&
                            classes.disabled
                        )}
                        onClick={() => {
                          if (!isConnected) {
                            setSelected(configBank);
                          }
                        }}
                      >
                        <img
                          alt="icon"
                          src={configBank?.logo}
                          style={{ width: "100%" }}
                        />
                      </Card>
                    </Grid>
                  );
                })}

                {banksConfig
                  .filter(
                    bank => !availableBanks.find(b => b.id === bank.value)
                  )
                  .map(bank => {
                    const isConnected = connectedBanks.includes(bank.value);

                    return (
                      <Grid item xs={6} sm={4} key={bank.value} id={bank.value}>
                        <Card
                          className={cx(
                            classes.cardRoot,
                            (isConnected || bank.disabled) && classes.disabled
                          )}
                          onClick={() => {
                            if (!(isConnected || bank.disabled)) {
                              setSelected(bank);
                            }
                          }}
                        >
                          <img
                            alt="icon"
                            src={bank.logo}
                            style={{ width: "100%" }}
                          />
                        </Card>
                      </Grid>
                    );
                  })}
              </Grid>
            </Grid>
          </Grid>
        )}
        <Slide
          direction="left"
          in={!!selected}
          mountOnEnter
          unmountOnExit
          exit={false}
        >
          <div>
            <Grid container justifyContent="center" style={{ marginTop: 64 }}>
              <Grid item xs={12} sm={8} md={6} style={{ textAlign: "center" }}>
                <Typography
                  color="primary"
                  variant="h2"
                  style={{ marginTop: 24, fontSize: 28 }}
                >
                  Requested access
                </Typography>
                <Typography variant="body1" style={{ marginTop: 16 }}>
                  In order to use this service, Financroo needs to access the
                  following information from your account service provider.
                </Typography>
                <div
                  style={{
                    display: "flex",
                    alignItems: "center",
                    justifyContent: "center",
                    marginTop: 32,
                  }}
                >
                  <div
                    style={{
                      display: "flex",
                      alignItems: "center",
                      justifyContent: "center",
                      background: "#FCFCFF",
                      width: 80,
                      height: 80,
                      borderRadius: "50%",
                      border: "1px solid rgb(236 236 236)",
                      marginRight: 16,
                    }}
                  >
                    <img
                      alt="icon"
                      src={financrooIcon}
                      style={{ width: "60%" }}
                    />
                  </div>
                  <img
                    alt="icon"
                    src={connectArrows}
                    style={{ marginRight: 16 }}
                  />
                  <div
                    style={{
                      display: "flex",
                      alignItems: "center",
                      justifyContent: "center",
                      background: "#FCFCFF",
                      width: 80,
                      height: 80,
                      borderRadius: "50%",
                      border: "1px solid rgb(236 236 236)",
                    }}
                  >
                    <img
                      alt="icon"
                      src={selected?.icon || selected?.logo}
                      style={{ width: "60%" }}
                    />
                  </div>
                </div>
                <Paper
                  style={{ marginTop: 32, padding: 16, textAlign: "left" }}
                >
                  <Typography
                    variant="h4"
                    style={{ fontSize: 16, marginBottom: 24 }}
                  >
                    What we need you to share
                  </Typography>
                  {selected?.permissions?.map((permission, index) => (
                    <Accordion elevation={0} key={permission.value + index}>
                      <AccordionSummary
                        expandIcon={<ExpandMoreIcon />}
                        aria-controls="panel1a-content"
                        id="panel1a-header"
                      >
                        <Typography>{permission.title}</Typography>
                      </AccordionSummary>
                      <AccordionDetails>
                        <Typography>{permission.description}</Typography>
                      </AccordionDetails>
                    </Accordion>
                  ))}
                </Paper>
                <Typography
                  style={{ marginTop: 32, display: "block" }}
                  variant="caption"
                >
                  Adding your accounts provides <strong>Financroo</strong> with
                  read-only access for 90 days. You can manage access at any
                  time. Authorizing will redirect to{" "}
                  <a href={`https://${selected?.value}.com`}>
                    https://{selected?.value}.com
                  </a>
                </Typography>
              </Grid>
            </Grid>
          </div>
        </Slide>
      </PageContainer>
      {selected && (
        <div className={classes.footer}>
          <div>
            <Button
              size="large"
              variant="outlined"
              id="cancel-button"
              onClick={() => setSelected(null)}
            >
              Cancel
            </Button>
            <Button
              size="large"
              variant="contained"
              id="allow-button"
              color="secondary"
              style={{ marginLeft: 16 }}
              onClick={() =>
                onAllowAccess(
                  selected.value,
                  selected?.permissions?.map(p => p.value)?.filter(p => p)
                )
              }
            >
              Allow access
            </Button>
          </div>
        </div>
      )}
    </Dialog>
  );
}
