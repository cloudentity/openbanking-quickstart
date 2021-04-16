import React from "react";
import Grid from "@material-ui/core/Grid";
import welcomeImage from "../assets/welcome-image.png";
import Typography from "@material-ui/core/Typography";
import Button from "@material-ui/core/Button";

export default function Welcome ({onConnectClick}) {

  return (
    <Grid container justify={'center'}>
      <Grid item xs={12} sm={8} md={6} style={{textAlign: 'center'}}>
        <img alt="welcome" src={welcomeImage} style={{marginTop: 128}}/>
        <Typography color={'primary'} variant={'h2'} style={{marginTop: 24, fontSize: 28}}>Welcome to Financroo Smart Banking </Typography>
        <Typography variant={'body1'} style={{marginTop: 16}}>Connect your bank, bills and credit cards to uncover insights that can improve
          your
          financial well being</Typography>
        <Button style={{marginTop: 24}} size={'large'} variant={'contained'} className={'connect-button'} color={'secondary'} onClick={onConnectClick}>Connect
          your bank</Button>
      </Grid>
    </Grid>
  )
};
