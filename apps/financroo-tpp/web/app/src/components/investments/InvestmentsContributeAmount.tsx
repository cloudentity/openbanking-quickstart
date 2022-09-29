import React from "react";
import TextField from "@mui/material/TextField";
import InputAdornment from "@mui/material/InputAdornment";
import { makeStyles } from "tss-react/mui";

import CardIcon from "../../assets/icon-card.svg";
import ContributionCard from "./ContributionCard";
import { getCurrency } from "../utils";

const useStyles = makeStyles()(theme => ({
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

interface Props {
  amount: string;
  setAmount: (amount: string) => void;
  handleBack: () => void;
  handleNext: () => void;
  accountDetails:
    | { amount: string | undefined; currency: string | undefined }
    | undefined;
  setAlert: (message: string) => void;
}

export default function InvestmentsContributeAmount({
  amount,
  setAmount,
  handleBack,
  handleNext,
  accountDetails,
  setAlert,
}: Props) {
  const { classes } = useStyles();

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
        id="amount-to-contribute"
        variant="outlined"
        value={amount}
        onChange={v => {
          setAmount(v.target.value);
          if (
            accountDetails &&
            accountDetails.amount &&
            Number(v.target.value) <= Number(accountDetails.amount)
          ) {
            setAlert("");
          }
        }}
        type="number"
        InputProps={{
          startAdornment: (
            <InputAdornment
              position="start"
              id="contribution-currency"
              style={{ position: "relative", top: 1 }}
            >
              {getCurrency(accountDetails?.currency)}
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
