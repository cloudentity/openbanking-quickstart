import React, {useEffect} from "react";
import FormControlLabel from "@material-ui/core/FormControlLabel";
import Checkbox from "@material-ui/core/Checkbox";
import FormHelperText from "@material-ui/core/FormHelperText";

export default function CheckBox ({form, id, name, label, style = {}, helperText, ...props}) {

  useEffect(() => {
    form.register({name});
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [name]);

  return (
    <div style={{marginBottom: 32, width: '100%', ...style}}>
      <FormControlLabel
        id={`${id}-${name}-checkbox`}
        control={
          <Checkbox
            checked={form.watch(name)}
            onChange={e => {
              props.onChange && props.onChange(e);
              form.setValue(name, e.target.checked);
            }}
            color="primary"
          />
        }
        label={label}
        {...props}
      />
      {helperText && <FormHelperText style={{marginLeft: 28}}>{helperText}</FormHelperText>
      }
    </div>
  )
};
