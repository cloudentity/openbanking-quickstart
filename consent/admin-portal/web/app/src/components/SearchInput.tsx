import React, { useState } from "react";
import { makeStyles } from "tss-react/mui";
import IconButton from "@mui/material/IconButton";
import OutlinedInput from "@mui/material/OutlinedInput";
import InputAdornment from "@mui/material/InputAdornment";
import CancelOutlined from "@mui/icons-material/CancelOutlined";
import Search from "@mui/icons-material/Search";

const useStyles = makeStyles()(theme => ({
  root: {
    padding: "16px 0 24px 0",
    width: "100%",
    boxSizing: "border-box",
  },
  margin: {
    margin: theme.spacing(1),
  },
  withoutLabel: {
    marginTop: theme.spacing(3),
  },
  textField: {
    width: "25ch",
  },
  searchIconContainer: {
    backgroundColor: "#DC1B37",
    height: 48,
    borderRadius: "0 4px 4px 0",
    display: "flex",
    marginLeft: 2,
  },
  input: {
    height: 48,
    boxSizing: "border-box",
  },
}));

interface Props {
  placeholder: string;
  onSearch: (text: string) => void;
  inputValue?: string;
}

export default function SearchInput({
  placeholder,
  onSearch,
  inputValue,
}: Props) {
  const { classes } = useStyles();
  const [value, setValue] = useState(inputValue || "");

  return (
    <div className={classes.root}>
      <div>
        <OutlinedInput
          id="outlined-adornment-password"
          type="text"
          value={value}
          onKeyDown={e => {
            if (e.key === "Enter" && value !== "") {
              onSearch(value);
            }
          }}
          onChange={e => {
            setValue(e.target.value);
          }}
          placeholder={placeholder}
          classes={{
            input: classes.input,
          }}
          style={{
            paddingRight: 0,
            height: 48,
            backgroundColor: "#F7FAFF",
            width: "100%",
          }}
          endAdornment={
            <InputAdornment position="end">
              {value !== "" ? (
                <IconButton
                  style={{ padding: 4 }}
                  onClick={() => setValue("")}
                  onMouseDown={() => setValue("")}
                  size="large"
                >
                  <CancelOutlined
                    fontSize="small"
                    style={{ color: "#BDBDBD" }}
                  />
                </IconButton>
              ) : (
                <div style={{ width: 25.5 }} />
              )}
              <div className={classes.searchIconContainer}>
                <IconButton
                  id="search-account"
                  onClick={() => value !== "" && onSearch(value)}
                  size="large"
                >
                  <Search style={{ color: "white", fontSize: 24 }} />
                </IconButton>
              </div>
            </InputAdornment>
          }
        />
      </div>
    </div>
  );
}
