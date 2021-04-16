import React from "react";
import {Theme} from "@material-ui/core";
import {makeStyles} from "@material-ui/core/styles";
import backgroundLogin from "../assets/background-login.png";
import financrooLogo from "../assets/financroo-logo.svg";
import Grid from "@material-ui/core/Grid";
import Button from "@material-ui/core/Button";

const useStyles = makeStyles((theme: Theme) => ({
  root: {
    height: '100vh'
  },
  image: {
    width: '100%',
    height: '100%',
    backgroundImage: `url(${backgroundLogin})`,
    backgroundRepeat: 'no-repeat',
    backgroundPosition: 'center',
    backgroundSize: 'cover'
  },
  formContainerRoot: {
    height: '100%',
    display: 'flex',
    flexDirection: 'column',
    alignItems: 'flex-start',
    justifyContent: 'center',

    [theme.breakpoints.down('sm')]: {
      padding: 16
    },
    [theme.breakpoints.up('sm')]: {
      padding: 48
    },
    [theme.breakpoints.up('lg')]: {
      padding: 82
    }
  }
}));

export default function Register({onLogin}) {
  const classes = useStyles();

  return (
    <div className={classes.root}>
      <Grid container style={{height: '100%'}}>
        <Grid item sm={6} lg={7}>
          <div className={classes.image}/>
        </Grid>
        <Grid item xs={12} sm={6} lg={5}>
          <div className={classes.formContainerRoot}>
            <img alt="financroo logo" src={financrooLogo} style={{marginBottom: 44}}/>
            <Button className={'login-button'} onClick={onLogin} color={'secondary'} style={{width: '100%', minHeight: 50}} variant={'contained'}>Login</Button>
          </div>
        </Grid>
      </Grid>
    </div>
  )
};
