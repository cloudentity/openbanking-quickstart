import React, { useEffect } from "react";
import FormControl from "@mui/material/FormControl";
import { pathOr } from "ramda";
import Autocomplete from "@mui/material/Autocomplete";
import TextField from "@mui/material/TextField";
import FormHelperText from "@mui/material/FormHelperText";
import FormInputLabel from "./FormInputLabel";

export default function AutocompleteField({
  id,
  form,
  name,
  label,
  helperText,
  ...props
}) {
  useEffect(() => {
    form.register({ name }, props.validate);
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [name]);

  return (
    <FormControl
      style={{ marginBottom: 32, width: "100%" }}
      error={!!pathOr("", ["errors", ...name.split(".")], form)}
    >
      <FormInputLabel
        id={`${id}-${name}-label`}
        label={label}
        icon={props.labelIcon}
        caption={props.labelCaption}
      />
      <Autocomplete
        options={props.options || []}
        id={`${id}-${name}-checkbox`}
        value={form.watch(name)}
        onChange={(e, option) => {
          form.setValue(name, option, { shouldValidate: true });
        }}
        renderInput={params => (
          <TextField
            {...params}
            error={!!pathOr("", ["errors", ...name.split(".")], form)}
            variant="outlined"
            fullWidth
          />
        )}
        autoHighlight
        {...props}
      />
      <FormHelperText id={`${id}-${name}-helper-text`} style={{ marginTop: 3 }}>
        {pathOr(helperText, ["errors", ...name.split("."), "message"], form)}
      </FormHelperText>
    </FormControl>
  );
}
