import React from "react";
import { Cell, Pie, PieChart, ResponsiveContainer } from "recharts";
import Typography from "@mui/material/Typography";
import classes from "./chartsStyles.module.css";
import { mapTransactionsToPieChartData, stringToHex } from "./analytics.utils";
import { Filter, Transaction } from "./types";

interface Props {
  currencyType: string;
  data: Transaction[];
  filtering: Filter;
  onChangeFiltering: (filter: Filter) => void;
}

export default function AnalyticsPieChart({
  currencyType,
  data,
  filtering,
  onChangeFiltering,
}: Props) {
  const mappedAsNameValue = mapTransactionsToPieChartData(data);

  const filteredByCategories =
    filtering.categories && filtering.categories.length > 0
      ? data.filter(t =>
          filtering.categories?.includes(t.BankTransactionCode.Code)
        )
      : data;

  const filteredSumAsString = filteredByCategories
    .reduce((sum, t) => sum + parseInt(t.Amount.Amount), 0)
    .toFixed(2);

  return (
    <div style={{ position: "relative" }}>
      <div className={classes.pieChartContent}>
        <Typography
          style={{
            fontSize: 14,
            fontWeight: 600,
            color: "#626576",
            minHeight: 22,
          }}
        >
          {filtering?.months?.join(" ")}
        </Typography>
        <Typography style={{ fontSize: 16, fontWeight: 600, marginTop: 6 }}>
          {currencyType} {filteredSumAsString}
        </Typography>
        <Typography style={{ fontSize: 12, marginTop: 2, color: "#626576" }}>
          {filtering?.categories?.join(" ")}
        </Typography>
      </div>
      <ResponsiveContainer width="100%" height={300}>
        <PieChart>
          <Pie
            data={mappedAsNameValue}
            // cx={120}
            // cy={200}
            innerRadius={90}
            outerRadius={120}
            fill="#8884d8"
            // paddingAngle={5}
            dataKey="value"
            onClick={e =>
              !filtering?.categories?.includes(e.name)
                ? onChangeFiltering({ categories: [e.name] })
                : onChangeFiltering({ categories: [] })
            }
          >
            {mappedAsNameValue.map((entry, index) => (
              <Cell
                cursor="pointer"
                key={`cell-${index}`}
                fill={
                  filtering?.categories?.includes(entry.name)
                    ? "#36C6AF"
                    : (stringToHex(entry.name) as string)
                }
              />
            ))}
          </Pie>
        </PieChart>
      </ResponsiveContainer>
    </div>
  );
}
