import React from "react";
import Card from "@mui/material/Card";
import AnalyticsTable, { mapTransactionToData } from "./AnalyticsTable";
import AnalyticsBarChart from "./AnalyticsBarChart";
import AnalyticsPieChart from "./AnalyticsPieChart";
import {
  applyFiltering,
  mapTransactionsToBarChartData,
} from "./analytics.utils";
import { Filter, Transaction } from "./types";

interface Props {
  currencyType: string;
  transactions: Transaction[];
  filtering: Filter;
  onChangeFiltering: (filter: Filter) => void;
}

export default function Analytics({
  currencyType,
  transactions,
  filtering,
  onChangeFiltering,
}: Props) {
  const barChartData = mapTransactionsToBarChartData(
    applyFiltering({ accounts: filtering.accounts }, transactions)
  );
  const pieChartData = applyFiltering(
    { accounts: filtering.accounts, months: filtering.months },
    transactions
  );
  const tableData = applyFiltering(filtering, transactions).map(
    mapTransactionToData
  );

  return (
    <>
      <Card style={{ padding: 16, display: "flex", alignItems: "center" }}>
        <div style={{ flex: 3 }}>
          <AnalyticsBarChart
            data={barChartData}
            filtering={filtering}
            onChangeFiltering={onChangeFiltering}
          />
        </div>
        <div style={{ flex: 1 }}>
          <AnalyticsPieChart
            currencyType={currencyType}
            data={pieChartData}
            filtering={filtering}
            onChangeFiltering={onChangeFiltering}
          />
        </div>
      </Card>
      <AnalyticsTable
        data={tableData}
        style={{ marginTop: 24, height: "calc(100% - 332px - 24px" }}
      />
    </>
  );
}
