import React from "react";
import TextField from "@material-ui/core/TextField";
import InputAdornment from "@material-ui/core/InputAdornment";
import { makeStyles } from "@material-ui/core/styles";

import CardIcon from "../../assets/icon-card.svg";
import ContributionCard from "./ContributionCard";

const useStyles = makeStyles((theme) => ({
  input: {
    marginTop: 30,
    "& label.Mui-focused": {
      color: "#ECECEC",
    },
    "& .MuiInput-underline:after": {
      borderBottomColor: theme.palette.primary.main,
    },
    "& .MuiOutlinedInput-root": {
      "& fieldset": {
        borderColor: "#ECECEC",
      },
      "&:hover fieldset": {
        borderColor: "#ECECEC",
      },
      "&.Mui-focused fieldset": {
        borderColor: theme.palette.primary.main,
      },
    },
  },
}));

export default function InvestmentsContributeAmount({
  amount,
  setAmount,
  handleBack,
  handleNext,
  account,
  setAlert,
}) {
  const classes = useStyles();

  return (
    <ContributionCard
      title="How much would you like to transfer?"
      backButton={{ title: "Back", onClick: handleBack }}
      nextButton={{
        title: "Next",
        onClick: handleNext,
        disabled: amount === "0" || !amount,
      }}
    >
      <img src={CardIcon} alt="icon" />
      <TextField
        classes={{
          root: classes.input,
        }}
        variant="outlined"
        value={amount}
        onChange={(v) => {
          setAmount(v.target.value);
          if (
            account &&
            account.Amount?.Amount &&
            Number(v.target.value) <= Number(account.Amount?.Amount)
          ) {
            setAlert("");
          }
        }}
        type="number"
        InputProps={{
          startAdornment: (
            <InputAdornment
              position="start"
              style={{ position: "relative", top: 1 }}
            >
              £
            </InputAdornment>
          ),
          inputProps: {
            min: 0,
          },
        }}
      />
    </ContributionCard>
  );
}
