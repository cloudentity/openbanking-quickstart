import React, { useState } from "react";
import { makeStyles } from "@material-ui/core/styles";
import { Theme } from "@material-ui/core";
import Tabs from "@material-ui/core/Tabs";
import Tab from "@material-ui/core/Tab";
import ConsentTable from "./ConsentTable";
import { ClientType } from "../utils";

const useStyles = makeStyles((theme: Theme) => ({
  container: {
    background: "#FFFFFF",
    boxShadow:
      "0px 1px 1px rgba(0, 0, 0, 0.08), 0px 0px 1px rgba(0, 0, 0, 0.31)",
    borderRadius: 4,
    maxWidth: 850,
    margin: "0 auto 24px auto",
    boxSizing: "border-box",
    paddingTop: 10,
    marginTop: 24,
  },
  header: {
    borderBottom: "1px solid #ECECEC",
    padding: "0 32px",
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
    minWidth: "unset",
    paddingLeft: 0,
    paddingRight: 0,
    marginRight: 48,
  },
}));

type Props = {
  consents: ClientType["consents"];
  id?: string;
};

function ConsentTabs({ consents, id }: Props) {
  const classes = useStyles();
  const [tab, setTab] = useState<"account" | "payment">("account");

  return (
    <div id={id} className={classes.container}>
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
          <ConsentTable
            data={consents.filter((v) => {
              return v.consent_type === "account_access" || v.consent_type === "cdr_arrangement" || v.consent_type === "consents"
            })}
            type="account"
          />
        )}
        {tab === "payment" && (
          <ConsentTable
            data={consents.filter((v) => v.consent_type === "domestic_payment")}
            type="payment"
          />
        )}
      </div>
    </div>
  );
}

export default ConsentTabs;
