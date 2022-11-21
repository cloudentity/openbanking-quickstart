import React, { useState } from "react";
import { makeStyles } from "tss-react/mui";
import Table from "@mui/material/Table";
import TableBody from "@mui/material/TableBody";
import TableCell from "@mui/material/TableCell";
import TableHead from "@mui/material/TableHead";
import TableRow from "@mui/material/TableRow";
import uniq from "lodash/uniq"

import ApplicationAccessPaymentDrawer from "./ApplicationAccessPaymentDrawer";
import ApplicationAccessAccountDrawer from "./ApplicationAccessAccountDrawer";
import Chip from "../Chip";
import { getDate } from "../ApplicationSimpleCard";
import { Consent, ConsentAccount } from "../types";
import { getCurrency } from "./utils";

const useStyles = makeStyles()(() => ({
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

function getTableBody(
  type: "account" | "payment",
  rows: Row[],
  setDrawerData: (consent: Consent | undefined) => void,
  data: Consent[]
) {
  if (type === "account") {
    return (
      <TableBody>
        {rows.map(row => (
          <TableRow
            key={row.id}
            className="consent-row"
            onClick={() => {
              setDrawerData(data.find(v => row.id === v.ConsentID));
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
              <Chip type={row.status && (row.status.toLowerCase() as any)}>
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
        {rows.map(row => (
          <TableRow
            key={row.id}
            className="consent-row"
            onClick={() => {
              setDrawerData(data.find(v => row.id === v.ConsentID));
            }}
          >
            <TableCell>{row.authorised}</TableCell>
            <TableCell>{row.account.data}</TableCell>
            <TableCell>{row.creditor}</TableCell>
            <TableCell>
              <Chip type={row.status && (row.status.toLowerCase() as any)}>
                {row.status}
              </Chip>
            </TableCell>
            <TableCell align="right">
              {getCurrency(row.currency)} {row.amount}
            </TableCell>
          </TableRow>
        ))}
      </TableBody>
    );
  }
  return null;
}

function getAccountNames(accountIds: string[], accounts: ConsentAccount[]) {
  const names = accountIds
    .map(id => {
      const found = accounts.find(v => v.id === id);
      if (found) {
        return `${found.name} *${found.id}`;
      }
      return null;
    })
    .filter(v => v);

  return {
    data: names.slice(0, 2).join(", "),
    more: names.length > 2 ? names.length - 2 : null,
  };
}

type Row = {
  authorised: string;
  account: { data: string; more: number | null };
  status: string;
  activeUntil?: string;
  id: string;
  creditor?: string;
  amount?: string;
  currency?: string;
};

interface Props {
  data: Consent[];
  accounts: ConsentAccount[];
  type: "account" | "payment";
  handleRevoke: (id: string, consent_type: string) => void;
  status: string;
}

function ApplicationAccessTable({
  data,
  type,
  handleRevoke,
  accounts,
  status,
}: Props) {
  const { classes } = useStyles();
  const [drawerPaymentData, setDrawerPaymentData] = useState<Consent>();
  const [drawerAccountData, setDrawerAccountData] = useState<Consent>();

  function createDataAccount(
    authorised: string,
    account: { data: string; more: number | null },
    status: string,
    activeUntil: string,
    id: string
  ) {
    return { authorised, account, status, activeUntil, id };
  }

  function createDataPayment(
    authorised: string,
    account: { data: string; more: number | null },
    creditor: string,
    status: string,
    amount: string,
    id: string,
    currency: string
  ) {
    return { authorised, account, creditor, status, amount, id, currency };
  }

  const rowsAccount =
    type === "account"
      ? data.map(
          ({
            AccountIDs,
            CreationDateTime,
            ExpirationDateTime,
            ConsentID,
            Status,
          }) =>
            createDataAccount(
              getDate(CreationDateTime),
              getAccountNames(uniq(AccountIDs ?? []), accounts),
              Status,
              getDate(ExpirationDateTime),
              ConsentID
            )
        )
      : [];

  const rowsPayment =
    type === "payment"
      ? data.map(
          ({
            AccountIDs,
            CreationDateTime,
            Status,
            ConsentID,
            Amount,
            CreditorAccountName,
            Currency,
          }) =>
            createDataPayment(
              getDate(CreationDateTime),
              getAccountNames(AccountIDs ?? [], accounts),
              CreditorAccountName,
              Status,
              Amount,
              ConsentID,
              Currency
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
