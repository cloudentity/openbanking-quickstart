import IconButton from '@material-ui/core/IconButton';
import InputAdornment from '@material-ui/core/InputAdornment';
import {makeStyles} from '@material-ui/core/styles';
import TextField from '@material-ui/core/TextField';
import {Visibility, VisibilityOff} from '@material-ui/icons';
import copy from 'clipboard-copy';
import React, {useState} from 'react';
import {Copy, ExternalLink} from 'react-feather';

const useStyles = makeStyles(theme => ({
  root: {
    '&:hover $adornment': {
      opacity: 1
    }
  },
  adornment: {
    opacity: 0.1
  },
  primaryMainColor: {
    color: theme.palette.primary.main
  },
  disabled: {
    color: 'rgba(0, 0, 0, 0.38)',
    cursor: 'default'
  }
}));

export default function AdornmentTextField ({label, getValue, disabled, toggleVisibility, defaultVisibility = true, withLink, withCopy, ...props}) {
  const classes = useStyles();

  const [visibility, setVisibility] = useState(defaultVisibility);

  return (
    <TextField
      {...props}
      type={visibility ? 'text' : 'password'}
      autoComplete={'new-password'}
      className={classes.root}
      InputProps={{
        readOnly: disabled,
        classes: {input: disabled && classes.disabled},
        endAdornment:
          <InputAdornment position="end" classes={{root: classes.adornment}}>
            {withCopy && <IconButton
              aria-label="copy to clipboard"
              tabIndex={-1}
              onClick={() => copy(getValue())}
            >
              <Copy className={classes.primaryMainColor}/>
            </IconButton>
            }
            {toggleVisibility && (
              <IconButton
                aria-label="toggle password visibility"
                tabIndex={-1}
                onClick={() => setVisibility(!visibility)}
              >
                {visibility ? <Visibility className={classes.primaryMainColor}/> : <VisibilityOff className={classes.primaryMainColor}/>}
              </IconButton>
            )}
            {withLink && (
              <IconButton
                aria-label="link"
                tabIndex={-1}
                onClick={() => window.open('string' === typeof withLink ? withLink : getValue(), '_blank')}
              >
                <ExternalLink className={classes.primaryMainColor}/>
              </IconButton>
            )}
          </InputAdornment>

      }}
    />
  );
};
