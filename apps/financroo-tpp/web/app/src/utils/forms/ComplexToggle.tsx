import React, {useEffect} from "react";
import {Avatar, Switch, Theme, Typography} from "@material-ui/core";
import {makeStyles} from "@material-ui/core/styles";
import FormInputLabel from "./FormInputLabel";

const useStyles = makeStyles((theme: Theme) => ({
  root: {
    padding: '24px 16px',
    border: '1px solid #ECECEC',
    borderRadius: 4,
    display: 'flex',
    alignItems: 'center',
    justifyContent: 'space-between',
    marginBottom: 32,
  },
  left: {
    display: 'flex',
    alignItems: 'center',
    justifyContent: 'center'
  },
  leftContent: {
    display: 'flex',
    flexDirection: 'column',
    alignItems: 'center',
    justifyContent: 'center'
  },
  title: {
    fontWeight: 'bold',
    fontSize: 12,
    color: '#626576',
    textTransform: 'uppercase'
  },
  subtitle: {
    marginTop: 8,
    fontSize: 12,
    color: '#626576',
  }
}));

export default function ComplexToggle ({form, name, label, title, subtitle, icon: Icon, ...props}) {
  const classes = useStyles();

  useEffect(() => {
    form.register({name});
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [name]);

  return (
    <>
      <FormInputLabel id={`${props.id}-label`} label={label}/>
      <div className={classes.root}>
        <div className={classes.left}>
          <Avatar style={{background: '#008EFF'}}>
            <Icon/>
          </Avatar>
          <div style={{marginLeft: 12}}>
            <Typography className={classes.title}>{title}</Typography>
            <Typography className={classes.subtitle}>{subtitle}</Typography>
          </div>
        </div>
        <Switch
          color={'primary'}
          checked={form.watch(name)}
          value={props.defaultValue}
          onChange={e => {
            props.onChange && props.onChange(e);
            form.setValue(name, e.target.checked);
          }}
          disabled={props.disabled}
        />
      </div>
    </>
  )
};
