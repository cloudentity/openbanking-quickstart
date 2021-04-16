import FormInputLabel from "./FormInputLabel";
import {pathOr} from "ramda";
import React from "react";
import AdornmentTextField from "./AdornmentTextField";

export default function CommonTextField({
                                          validate, form, id, name, label, defaultValue = '', helperText = '', disabled, externalErrors = null, style, onChange = () => {
  }, ...props
                                        }) {
  return (
    <div style={{width: '100%', marginBottom: 32, ...style}}>
      {!props.hideLabel && <FormInputLabel id={`${id}-${name}-label`} label={label} icon={props.labelIcon} caption={props.labelCaption}/>}
      <AdornmentTextField
        id={`${id}-${name}-input`}
        name={name}
        label={label}
        onChange={onChange}
        getValue={() => form.getValues(name)}
        toggleVisibility={props.toggleVisibility}
        defaultVisibility={props.defaultVisibility}
        disabled={!!disabled}
        defaultValue={defaultValue}
        error={!!pathOr(externalErrors, ['errors', ...name.split('.')], form)}
        helperText={pathOr(externalErrors || helperText, ['errors', ...name.split('.'), 'message'], form)}
        inputRef={form.register({validate})}
        fullWidth
        variant="outlined"
        withCopy={disabled}
        withLink={props.withLink}
        {...props}
      />
    </div>
  )
};
