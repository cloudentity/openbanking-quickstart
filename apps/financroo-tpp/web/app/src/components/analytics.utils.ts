import {
  lensIndex,
  lensProp,
  map,
  over,
  pipe,
  reduce,
  toPairs,
  zip,
} from "ramda";
import { Transaction } from "./types";

const months = [
  "JAN",
  "FEB",
  "MAR",
  "APR",
  "MAY",
  "JUN",
  "JUL",
  "AUG",
  "SEP",
  "OCT",
  "NOV",
  "DEV",
];

export const applyFiltering = (
  filtering: { accounts?: string[]; months?: string[]; categories?: string[] },
  data: Transaction[]
) => {
  let transactions = data;

  if (filtering.accounts && filtering.accounts.length > 0) {
    transactions = transactions.filter(t =>
      filtering.accounts?.includes(t.AccountId)
    );
  } else {
    transactions = [];
  }

  if (filtering.months && filtering.months.length > 0) {
    transactions = transactions.filter(t => {
      const arr = t.BookingDateTime.split("-");
      const monthIndex = parseInt(arr[1], 10) - 1;
      const monthName = months[monthIndex];

      return filtering.months?.includes(monthName);
    });
  }

  if (filtering?.categories && filtering.categories.length > 0) {
    transactions = transactions.filter(t =>
      filtering.categories?.includes(t.BankTransactionCode.Code)
    );
  }

  return transactions;
};

export const mapTransactionsToBarChartData = pipe(
  reduce(
    (acc, t: any) => {
      const arr = t.BookingDateTime.split("-");
      const monthIndex = parseInt(arr[1], 10) - 1;

      return over(
        lensIndex(monthIndex),
        v => v + parseFloat(t.Amount.Amount),
        acc
      );
    },
    months.map(() => 0)
  ),
  zip(months),
  map(([name, value]) => ({ name, value }))
);

export const mapTransactionsToPieChartData = pipe(
  reduce(
    (acc: any, t: any) =>
      over(
        lensProp(t.BankTransactionCode.Code),
        v => (v || 0) + parseFloat(t.Amount.Amount),
        acc
      ),
    {}
  ),
  toPairs,
  map(([name, value]) => ({ name, value }))
);

export const stringToHex = str => {
  var hash = 0;
  if (str.length === 0) return hash;
  for (var i = 0; i < str.length; i++) {
    hash = str.charCodeAt(i) + ((hash << 5) - hash);
    hash = hash & hash;
  }
  var color = "#";
  for (var j = 0; j < 3; j++) {
    var value = (hash >> (j * 8)) & 255;
    color += ("00" + value.toString(16)).substr(-2);
  }
  return color;
};
