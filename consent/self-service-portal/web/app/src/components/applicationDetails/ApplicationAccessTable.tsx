import React, { useState } from "react";
import { makeStyles } from "@material-ui/core/styles";
import { Theme } from "@material-ui/core";
import Table from "@material-ui/core/Table";
import TableBody from "@material-ui/core/TableBody";
import TableCell from "@material-ui/core/TableCell";
import TableHead from "@material-ui/core/TableHead";
import TableRow from "@material-ui/core/TableRow";

import ApplicationAccessPaymentDrawer from "./ApplicationAccessPaymentDrawer";
import ApplicationAccessAccountDrawer from "./ApplicationAccessAccountDrawer";
import Chip from "../Chip";
import { getDate } from "../ApplicationSimpleCard";

const useStyles = makeStyles((theme: Theme) => ({
  table: {
    "& tr": {
      borderBottom: "solid 1px #ECECEC",
      cursor: "pointer",
      "&:hover": {
        backgroundColor: "#FCFCFF",
      },
    },
    "& th": {
      padding: "10px 16px",
      cursor: "default",
    },
    "& tbody > tr:last-of-type": {
      borderBottom: "none",
    },
  },
  drawerContent: {
    width: 500,
  },
  empty: {
    display: "flex",
    alignItems: "center",
    justifyContent: "center",
    minHeight: 100,
    width: "100%",
    color: "gray",
  },
}));

function getTableHead(type: "account" | "payment") {
  if (type === "account") {
    return (
      <TableHead>
        <TableRow>
          <TableCell>Authorised</TableCell>
          <TableCell>Account</TableCell>
          <TableCell>Status</TableCell>
          <TableCell align="right">Active until</TableCell>
        </TableRow>
      </TableHead>
    );
  }
  if (type === "payment") {
    return (
      <TableHead>
        <TableRow>
          <TableCell>Authorised</TableCell>
          <TableCell>Account</TableCell>
          <TableCell>Creditor</TableCell>
          <TableCell>Status</TableCell>
          <TableCell align="right">Amount</TableCell>
        </TableRow>
      </TableHead>
    );
  }
  return null;
}

function getTableBody(type: "account" | "payment", rows, setDrawerData, data) {
  if (type === "account") {
    return (
      <TableBody>
        {rows.map((row) => (
          <TableRow
            key={row.id}
            className={`consent-row`}
            onClick={() => {
              setDrawerData(data.find((v) => row.id === v.ConsentID));
            }}
          >
            <TableCell>{row.authorised}</TableCell>
            <TableCell>
              {row.account.data}
              {row.account.more && (
                <span
                  style={{
                    color: "#212533",
                    background: "#ECECEC",
                    borderRadius: 10,
                    padding: "2px 6px",
                    fontWeight: 500,
                    fontSize: 12,
                    lineHeight: "24px",
                    marginLeft: 8,
                  }}
                >
                  + {row.account.more}
                </span>
              )}
            </TableCell>
            <TableCell>
              <Chip type={row.status && row.status.toLowerCase()}>
                {row.status}
              </Chip>
            </TableCell>
            <TableCell align="right">{row.activeUntil}</TableCell>
          </TableRow>
        ))}
      </TableBody>
    );
  }
  if (type === "payment") {
    return (
      <TableBody>
        {rows.map((row) => (
          <TableRow
            key={row.id}
            className={`consent-row`}
            onClick={() => {
              setDrawerData(data.find((v) => row.id === v.ConsentID));
            }}
          >
            <TableCell>{row.authorised}</TableCell>
            <TableCell>{row.account.data}</TableCell>
            <TableCell>{row.creditor}</TableCell>
            <TableCell>
              <Chip type={row.status && row.status.toLowerCase()}>
                {row.status}
              </Chip>
            </TableCell>
            <TableCell align="right">£ {row.amount}</TableCell>
          </TableRow>
        ))}
      </TableBody>
    );
  }
  return null;
}

function getAccountNames(accountIds, accounts) {
  const names = accountIds
    .map((id) => {
      const found = accounts.find((v) => v.id === id);
      if (found) {
        return `${found.name} *${found.id}`;
      }
      return null;
    })
    .filter((v) => v);

  return {
    data: names.slice(0, 2).join(", "),
    more: names.length > 2 ? names.length - 2 : null,
  };
}

type Props = {
  data: any;
  accounts: any;
  type: "account" | "payment";
  handleRevoke: (id: string) => void;
  status: string;
};

function ApplicationAccessTable({
  data,
  type,
  handleRevoke,
  accounts,
  status,
}: Props) {
  const classes = useStyles();
  const [drawerPaymentData, setDrawerPaymentData] = useState<any>(null); //FIXME any
  const [drawerAccountData, setDrawerAccountData] = useState<any>(null); //FIXME any

  function createDataAccount(authorised, account, status, activeUntil, id) {
    return { authorised, account, status, activeUntil, id };
  }

  function createDataPayment(
    authorised,
    account,
    creditor,
    status,
    amount,
    id
  ) {
    return { authorised, account, creditor, status, amount, id };
  }

  const rowsAccount =
    type === "account"
      ? data.map(({ AccountIDs, CreationDateTime, ExpirationDateTime, ConsentID, Status }) =>
          createDataAccount(
            getDate(CreationDateTime),
            getAccountNames(AccountIDs ?? [], accounts),
            Status,
            getDate(ExpirationDateTime),
            ConsentID
          )
        )
      : [];

  const rowsPayment =
    type === "payment"
      ? data.map(({ account_ids, domestic_payment_consent }) =>
          createDataPayment(
            getDate(domestic_payment_consent?.CreationDateTime),
            getAccountNames(
              account_ids ?? [],
              accounts
            ),
            domestic_payment_consent?.Initiation?.CreditorAccount?.Name,
            domestic_payment_consent?.Status,
            domestic_payment_consent?.Initiation?.InstructedAmount?.Amount,
            domestic_payment_consent?.ConsentID
          )
        )
      : [];

  return (
    <>
      <Table className={classes.table} aria-label="simple table">
        {getTableHead(type)}
        {getTableBody(
          type,
          type === "account" ? rowsAccount : rowsPayment,
          type === "account" ? setDrawerAccountData : setDrawerPaymentData,
          data
        )}
      </Table>

      {type === "account" && rowsAccount.length === 0 && (
        <div className={classes.empty}>No data</div>
      )}
      {type === "payment" && rowsPayment.length === 0 && (
        <div className={classes.empty}>No data</div>
      )}

      {drawerPaymentData && (
        <ApplicationAccessPaymentDrawer
          drawerData={drawerPaymentData}
          setDrawerData={setDrawerPaymentData}
          status={status}
        />
      )}
      {drawerAccountData && (
        <ApplicationAccessAccountDrawer
          drawerData={drawerAccountData}
          setDrawerData={setDrawerAccountData}
          handleRevoke={handleRevoke}
          accounts={accounts}
          status={status}
        />
      )}
    </>
  );
}

export default ApplicationAccessTable;
