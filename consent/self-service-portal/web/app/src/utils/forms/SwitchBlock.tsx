import {FormControl, Switch} from "@material-ui/core";
import React, {useEffect} from "react";
import FormControlLabel from "@material-ui/core/FormControlLabel";

export default function SwitchBlock({form, id, name, label, style = {}, onChange}) {
  useEffect(() => {
    form.register({name});
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [name]);

  return (
    <FormControl fullWidth style={style}>
      <FormControlLabel
        label={label}
        id={`${id}-${name}-switch-label`}
        control={<Switch
          id={`${id}-${name}-switch`}
          checked={form.watch(name)}
          onChange={e => {
            onChange && onChange(e);
            form.setValue(name, e.target.checked);
          }}
          color="primary"
        />
        }
      />
    </FormControl>
  )
}
