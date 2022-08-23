import React, { useState } from "react";
import { makeStyles } from "tss-react/mui";
import Table from "@mui/material/Table";
import TableBody from "@mui/material/TableBody";
import TableCell from "@mui/material/TableCell";
import TableHead from "@mui/material/TableHead";
import TableRow from "@mui/material/TableRow";

import Chip from "../Chip";
import { getDate } from "../utils";
import PaymentDrawer from "./Drawers/PaymentDrawer";
import AccountAccessDrawer from "./Drawers/AccountAccessDrawer";

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
      "&:first-of-type": {
        paddingLeft: 32,
      },
      "&:last-of-type": {
        paddingRight: 32,
      },
    },
    "& tbody > tr:last-of-type": {
      borderBottom: "none",
    },
    "& td": {
      "&:first-of-type": {
        paddingLeft: 32,
      },
      "&:last-of-type": {
        paddingRight: 32,
      },
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

const useTableStyles = makeStyles()(() => ({
  tableRowRoot: {
    "& th": {
      fontWeight: "bold",
    },
  },
}));

function TableHeadComponent({ type }: { type: "account" | "payment" }) {
  const { classes } = useTableStyles();

  if (type === "account") {
    return (
      <TableHead>
        <TableRow classes={{ root: classes.tableRowRoot }}>
          <TableCell>Account</TableCell>
          <TableCell>Consent ID</TableCell>
          <TableCell>Authorised</TableCell>
          <TableCell>Active until</TableCell>
          <TableCell align="right">Status</TableCell>
        </TableRow>
      </TableHead>
    );
  }
  if (type === "payment") {
    return (
      <TableHead>
        <TableRow classes={{ root: classes.tableRowRoot }}>
          <TableCell>Authorised</TableCell>
          <TableCell>Account</TableCell>
          <TableCell>Consent ID</TableCell>
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
        {rows.map(row => (
          <TableRow
            key={row.id}
            className="consent-row"
            onClick={() => {
              setDrawerData(data.find(v => row.id === v.consent_id));
            }}
          >
            <TableCell>{row.account}</TableCell>
            <TableCell>{row.id}</TableCell>
            <TableCell>{row.authorised}</TableCell>
            <TableCell>{row.activeUntil}</TableCell>

            <TableCell align="right">
              <Chip type={row.status && row.status.toLowerCase()}>
                {row.status}
              </Chip>
            </TableCell>
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
              setDrawerData(data.find(v => row.id === v.consent_id));
            }}
          >
            <TableCell>{row.authorised}</TableCell>
            <TableCell>{row.account}</TableCell>
            <TableCell>{row.id}</TableCell>
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

interface Props {
  data: any;
  type: "account" | "payment";
}

function ConsentTable({ data, type }: Props) {
  const { classes } = useStyles();
  const [drawerPaymentData, setDrawerPaymentData] = useState<any>(null);
  const [drawerAccountData, setDrawerAccountData] = useState<any>(null);

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
      ? data.map(
          ({ created_at, account_ids, status, expires_at, consent_id }) =>
            createDataAccount(
              getDate(created_at),
              account_ids.join(", "),
              status,
              getDate(expires_at),
              consent_id
            )
        )
      : [];

  const rowsPayment =
    type === "payment"
      ? data.map(({ domestic_payment, account_ids }) =>
          createDataPayment(
            getDate(domestic_payment?.created_at),
            account_ids.join(", "),
            domestic_payment?.Initiation?.CreditorAccount?.Name,
            domestic_payment?.status,
            domestic_payment?.Initiation?.InstructedAmount?.Amount,
            domestic_payment?.consent_id
          )
        )
      : [];

  return (
    <>
      <Table className={classes.table} aria-label="simple table">
        <TableHeadComponent type={type} />
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
        <PaymentDrawer
          drawerData={drawerPaymentData}
          setDrawerData={setDrawerPaymentData}
        />
      )}
      {drawerAccountData && (
        <AccountAccessDrawer
          drawerData={drawerAccountData}
          setDrawerData={setDrawerAccountData}
        />
      )}
    </>
  );
}

export default ConsentTable;
