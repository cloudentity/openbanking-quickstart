import React from "react";
import Typography from "@mui/material/Typography";
import Button from "@mui/material/Button";
import BankCard from "./BankCard";
import mainClasses from "./main.module.css";
import Card from "@mui/material/Card";

export default function Accounts({
  banks,
  reconnectBank,
  accounts,
  balances,
  filtering,
  onChangeFiltering,
  onConnectClick,
  onDisconnect,
  onReconnect,
}) {
  const totalBalance = balances
    .reduce((total, b) => total + parseFloat(b.Amount), 0)
    .toFixed(2);
  const currencyType = balances[0]?.Currency || "N/A";

  return (
    <div style={{ height: "100%", display: "flex", flexDirection: "column" }}>
      <Card style={{ padding: "32px 20px" }}>
        <div
          style={{
            display: "flex",
            alignItems: "center",
            justifyContent: "space-between",
          }}
        >
          <Typography className={mainClasses.sectionTitle}>
            All accounts
          </Typography>
          <Typography>
            <strong>
              {currencyType} {totalBalance}
            </strong>
          </Typography>
        </div>
        <Typography
          style={{
            display: "inline-block",
            background: "rgba(54, 198, 175, 0.08)",
            color: "#36C6AF",
            fontSize: 14,
            padding: 2,
            marginTop: 4,
          }}
        >
          {accounts.length} accounts synced
        </Typography>
      </Card>

      {banks.map(bankId => (
        <BankCard
          key={bankId}
          bankId={bankId}
          reconnect={reconnectBank}
          accounts={accounts.filter(a => a.BankId === bankId)}
          balances={balances}
          filtering={filtering}
          onChangeFiltering={onChangeFiltering}
          onDisconnect={onDisconnect}
          onReconnect={onReconnect}
          style={{ marginTop: 32 }}
        />
      ))}

      <div style={{ flex: 1 }} />

      <Button
        color="secondary"
        variant="contained"
        size="large"
        style={{ width: "100%" }}
        onClick={onConnectClick}
      >
        Connect your bank
      </Button>
    </div>
  );
}
