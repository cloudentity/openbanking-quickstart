import {includes, lensIndex, lensProp, map, over, pipe, reduce, toPairs, zip} from "ramda";

const months = ["JAN", "FEB", "MAR", "APR", "MAY", "JUN", "JUL", "AUG", "SEP", "OCT", "NOV", "DEV"];

export const applyFiltering = (filtering, data) => {
  let transactions = data;

  if (filtering?.accounts?.length > 0) {
    transactions = transactions.filter(t => includes(t.AccountId, filtering.accounts));
  } else {
    transactions = [];
  }

  if (filtering?.months?.length > 0) {
    transactions = transactions.filter(t => {
      const arr = t.BookingDateTime.split('-');
      const monthIndex = parseInt(arr[1], 10) - 1;
      const monthName = months[monthIndex];

      return includes(monthName, filtering.months);
    });
  }

  if (filtering?.categories?.length > 0) {
    transactions = transactions.filter(t => includes(t.BankTransactionCode.Code, filtering.categories));
  }

  return transactions;
}

export const mapTransactionsToBarChartData = pipe(
  reduce((acc, t: any) => {
    const arr = t.BookingDateTime.split('-');
    const monthIndex = parseInt(arr[1], 10) - 1;

    return over(lensIndex(monthIndex), v => v + parseFloat(t.Amount.Amount), acc);
  }, months.map(() => 0)),
  zip(months),
  map(([name, value]) => ({name, value}))
);

export const mapTransactionsToPieChartData = pipe(
  reduce((acc, t: any) => over(lensProp(t.BankTransactionCode.Code), v => (v || 0) + parseFloat(t.Amount.Amount), acc), {}),
  toPairs,
  map(([name, value]) => ({name, value}))
);

export const stringToHex = str => {
  var hash = 0;
  if (str.length === 0) return hash;
  for (var i = 0; i < str.length; i++) {
    hash = str.charCodeAt(i) + ((hash << 5) - hash);
    hash = hash & hash;
  }
  var color = '#';
  for (var j = 0; j < 3; j++) {
    var value = (hash >> (j * 8)) & 255;
    color += ('00' + value.toString(16)).substr(-2);
  }
  return color;
}

