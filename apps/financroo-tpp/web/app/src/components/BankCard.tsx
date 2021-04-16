import React from "react";
import {Button, Theme} from "@material-ui/core";
import {makeStyles} from "@material-ui/core/styles";
import Typography from "@material-ui/core/Typography";
import IconButton from "@material-ui/core/IconButton";
import {Plus} from "react-feather";
import Card from "@material-ui/core/Card";
import Checkbox from "@material-ui/core/Checkbox";
import {banks} from "./banks";
import {filter, pathOr} from "ramda";
import requestAccessPermissions from "./request-access-permissions.json";

const useStyles = makeStyles((theme: Theme) => ({
  accountRoot: {
    borderBottom: '1px solid #ECECEC',
    '&:hover': {
      cursor: 'pointer'
    }
  }
}));

export default function BankCard({bankId, reconnect, accounts, balances, filtering, style = {}, onChangeFiltering, onDisconnect, onReconnect}) {
  const classes = useStyles();

  const getAccountBalance = (accountId, balances) => balances.find(b => b.AccountId === accountId);
  const getAccountAmountAsString = (accountId, balances) => {
    const accountBalance = getAccountBalance(accountId, balances);
    return accountBalance
      ? `GBP ${pathOr(0, ['Amount', 'Amount'], accountBalance)}`
      : 'N/A'

  }
  const isAccountChecked = id => filtering?.accounts?.includes(id);

  return (
    <Card style={style} id={bankId}>
      <div style={{padding: 20, display: 'flex', alignItems: 'center', borderBottom: '1px solid #ECECEC'}}>
        <div style={{
          background: '#FCFCFF',
          borderRadius: '50%',
          width: 52,
          height: 52,
          display: 'flex',
          alignItems: 'center',
          justifyContent: 'center',
          boxShadow: '0px 0.574468px 0.574468px rgba(0, 0, 0, 0.08), 0px 0px 0.574468px rgba(0, 0, 0, 0.31)'
        }}>
          <img src={banks.find(b => b.value === bankId)?.icon} style={{width: 24, height: 24}} alt={'bank icon'}/>
        </div>
        <div style={{marginLeft: 24}}>
          <Typography>{banks.find(b => b.value === bankId)?.name}</Typography>
          <Typography
            style={{background: 'rgba(54, 198, 175, 0.08)', color: '#36C6AF', fontSize: 14, padding: 2, marginTop: 4}}
          >{accounts.length} accounts synced</Typography>
        </div>
        <div style={{flex: 1}}/>
        <div>
          {reconnect && (
            <Button size={"small"}
                    className={`reconnect-button`}
                    variant={'contained'}
                    color={'primary'}
                    style={{color: '#fff'}}
                    onClick={onReconnect(bankId, requestAccessPermissions.permissions.map(p => p.value).filter(p => p))}>reconnect</Button>
          )}
          {!reconnect && (
            <Button size={"small"} className={'disconnect-button'} variant={'outlined'} onClick={onDisconnect(bankId)}>disconnect</Button>
          )}
        </div>
      </div>
      {accounts.map(account => (
        <div
          key={account.AccountId}
          onClick={() => onChangeFiltering(
            {
              accounts:
                isAccountChecked(account.AccountId)
                  ? filter(a => a !== account.AccountId, filtering?.accounts)
                  : [...filtering?.accounts, account.AccountId],
              months: [],
              categories: []
            })}
          className={classes.accountRoot}
          style={{
            height: 62,
            background: isAccountChecked(account.AccountId) ? '#36C6AF' : 'initial',
            color: isAccountChecked(account.AccountId) ? '#fff' : 'initial',
            display: 'flex',
            alignItems: 'center',
            justifyContent: 'space-between',
            padding: '0 20px'
          }}>

          <div style={{display: 'flex', alignItems: 'center'}}>
            <Checkbox
              checked={isAccountChecked(account.AccountId)}
              onChange={() => onChangeFiltering(
                {
                  accounts:
                    isAccountChecked(account.AccountId)
                      ? filter(a => a !== account.AccountId, filtering?.accounts)
                      : [...filtering?.accounts, account.AccountId],
                  months: [],
                  categories: []
                })}
              color={'primary'}
              style={{color: isAccountChecked(account.AccountId) ? '#fff' : 'initial'}}
              inputProps={{'aria-label': 'primary checkbox'}}
            />
            <div style={{marginLeft: 12}}>
              <Typography className={`account-name`}>{account.Nickname}</Typography>
              <Typography>**** ***** **** {account.AccountId}</Typography>
            </div>
          </div>
          <div>
            <Typography>{getAccountAmountAsString(account.AccountId, balances)}</Typography>
          </div>
        </div>
      ))}
      <div style={{
        height: 52,
        padding: '0 21px',
        background: 'rgba(54, 198, 175, 0.08)',
        color: '#36C6AF',
        display: 'flex',
        alignItems: 'center',
        justifyContent: 'space-between'
      }}>
        <Typography>Add new account</Typography>
        <IconButton>
          <Plus style={{color: '#36C6AF'}}/>
        </IconButton>
      </div>
    </Card>
  )
}
;
