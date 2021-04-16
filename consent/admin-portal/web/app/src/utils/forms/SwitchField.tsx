import React, {useEffect} from "react";
import FormInputLabel from "./FormInputLabel";
import {Switch} from "@material-ui/core";

export default function SwitchField({form, id, name, label, ...props}) {

  useEffect(() => {
    form.register({name});
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [name]);

  return (
    <div style={{marginBottom: 32, width: '100%'}}>
      <FormInputLabel id={`${id}-${name}-switch-label`} label={label}/>
      <Switch
        id={`${id}-${name}-switch`}
        checked={form.watch(name)}
        onChange={e => {
          props.onChange && props.onChange(e);
          form.setValue(name, e.target.checked);
        }}
        color="primary"
      />
    </div>
  )
};
