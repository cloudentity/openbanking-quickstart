import React, { useState } from "react";
import { makeStyles } from "@material-ui/core/styles";
import { Theme } from "@material-ui/core";
import Tabs from "@material-ui/core/Tabs";
import Tab from "@material-ui/core/Tab";
import ApplicationAccessTable from "./ApplicationAccessTable";

const useStyles = makeStyles((theme: Theme) => ({
  container: {
    background: "#FFFFFF",
    boxShadow:
      "0px 1px 1px rgba(0, 0, 0, 0.08), 0px 0px 1px rgba(0, 0, 0, 0.31)",
    borderRadius: 4,
    maxWidth: 850,
    margin: "0 auto 24px auto",
    boxSizing: "border-box",
    paddingTop: 18,
  },
  header: {
    borderBottom: "1px solid #ECECEC",
  },
  tabsRoot: {
    textTransform: "none",
    color: "#626576",
  },
  tabsIndicator: {
    backgroundColor: "#DC1B37",
  },
  tab: {
    ...theme.custom.body2,
    textTransform: "none",
  },
}));

type Props = {
  data: any; // FIXME
  accounts: any; // FIXME
  handleRevoke: (id: string, consent_type: string) => void;
  status: string;
};

function ApplicationAccessTabs({
  data,
  handleRevoke,
  accounts,
  status,
}: Props) {
  const classes = useStyles();
  const [tab, setTab] = useState<"account" | "payment">("account");

  return (
    <div className={classes.container}>
      <div className={classes.header}>
        <Tabs
          value={tab}
          onChange={(_, newValue) => setTab(newValue)}
          classes={{
            root: classes.tabsRoot,
            indicator: classes.tabsIndicator,
          }}
        >
          <Tab
            value="account"
            label="Account access"
            className={`${classes.tab} accounts-tab`}
            style={tab === "account" ? { color: "#DC1B37" } : {}}
          />
          <Tab
            value="payment"
            label="Payment access"
            className={`${classes.tab} payments-tab`}
            style={tab === "payment" ? { color: "#DC1B37" } : {}}
          />
        </Tabs>
      </div>
      <div>
        {tab === "account" && (
          <ApplicationAccessTable
            data={data.consents.filter((v) => v.type === "account_access" || v.type == "cdr_arrangement" || v.type == "consents")}
            type="account"
            handleRevoke={handleRevoke}
            accounts={accounts}
            status={status}
          />
        )}
        {tab === "payment" && (
          <ApplicationAccessTable
            data={data.consents.filter((v) => v.type === "domestic_payment")}
            type="payment"
            handleRevoke={handleRevoke}
            accounts={accounts}
            status={status}
          />
        )}
      </div>
    </div>
  );
}

export default ApplicationAccessTabs;
