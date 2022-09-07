import React from "react";
import Grid from "@mui/material/Grid";
import welcomeImage from "../assets/welcome-image.png";
import Typography from "@mui/material/Typography";
import Button from "@mui/material/Button";

interface Props {
  onConnectClick: () => void;
}

export default function Welcome({ onConnectClick }: Props) {
  return (
    <Grid container justifyContent="center">
      <Grid item xs={12} sm={8} md={6} style={{ textAlign: "center" }}>
        <img alt="welcome" src={welcomeImage} style={{ marginTop: 128 }} />
        <Typography
          color="primary"
          variant="h2"
          style={{ marginTop: 24, fontSize: 28 }}
        >
          Welcome to Financroo Smart Banking{" "}
        </Typography>
        <Typography variant="body1" style={{ marginTop: 16 }}>
          Connect your bank, bills and credit cards to uncover insights that can
          improve your financial well being
        </Typography>
        <Button
          style={{ marginTop: 24 }}
          size="large"
          variant="contained"
          className="connect-button"
          color="secondary"
          onClick={onConnectClick}
        >
          Connect your bank
        </Button>
      </Grid>
    </Grid>
  );
}
