import React, { ReactNode, useState } from "react";
import { makeStyles } from "@material-ui/core/styles";
import { Theme } from "@material-ui/core";
import Tabs from "@material-ui/core/Tabs";
import Tab from "@material-ui/core/Tab";

const useStyles = makeStyles((theme: Theme) => ({
  container: {
    background: "#FFFFFF",
    boxShadow:
      "0px 1px 1px rgba(0, 0, 0, 0.08), 0px 0px 1px rgba(0, 0, 0, 0.31)",
    borderRadius: 4,
    maxWidth: 850,
    margin: "0 auto 24px auto",
    boxSizing: "border-box",
    padding: "10px 24px 0 24px",
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
  content: {
    color: "black",
  },
  tabRoot: {
    minWidth: "unset",
  },
}));

type Props = {
  tabs: {
    key: string;
    label: string;
    content: ReactNode;
  }[];
};

function CustomTabs({ tabs }: Props) {
  const classes = useStyles();
  const [tab, setTab] = useState(tabs[0].key);

  const curentTab = tabs.find(({ key }) => tab === key);

  return (
    <div className={classes.container}>
      <div>
        <Tabs
          value={tab}
          onChange={(_, newValue) => setTab(newValue)}
          classes={{
            root: classes.tabsRoot,
            indicator: classes.tabsIndicator,
          }}
        >
          {tabs.map(({ key, label }) => (
            <Tab
              key={key}
              value={key}
              label={label}
              className={`${classes.tab} accounts-tab`}
              style={tab === key ? { color: "#DC1B37" } : {}}
              classes={{
                root: classes.tabRoot,
              }}
            />
          ))}
        </Tabs>
      </div>
      <div className={classes.content}>{curentTab?.content}</div>
    </div>
  );
}

export default CustomTabs;
