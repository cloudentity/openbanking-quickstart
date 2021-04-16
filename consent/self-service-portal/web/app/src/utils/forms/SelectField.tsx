import React, {useEffect} from "react";
import FormControl from "@material-ui/core/FormControl";
import {pathOr} from "ramda";
import FormInputLabel from "./FormInputLabel";
import {Controller} from "react-hook-form";
import Select from "@material-ui/core/Select";
import MenuItem from "@material-ui/core/MenuItem";
import FormHelperText from "@material-ui/core/FormHelperText";

export default function SelectField({validate, id, form, name, label, options, defaultValue = '', helperText = '', ...props}) {

  useEffect(() => {
    form.register({name});
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [name]);

  return (
    <FormControl variant="outlined" fullWidth error={!!pathOr(null, ['errors', ...name.split('.')], form)} style={{marginBottom: 32}}>
      <FormInputLabel id={`${id}-${name}-label`} label={label}/>
      <Controller
        as={
          <Select
            id={id}
            fullWidth
          >
            {(options || []).map(option => <MenuItem key={option.value} value={option.value}>{option.name}</MenuItem>)}
          </Select>
        }
        rules={validate}
        control={form.control}
        name={name}
        defaultValue={defaultValue}
        disabled={props.disabled}
      />
      <FormHelperText id={`${id}-helper-text`}>{pathOr(helperText, ['errors', ...name.split('.'), 'message'], form)}</FormHelperText>
    </FormControl>
  )
}
