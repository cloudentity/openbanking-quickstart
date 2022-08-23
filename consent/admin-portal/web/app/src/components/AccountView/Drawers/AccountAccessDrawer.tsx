import React from "react";
import { makeStyles } from "@material-ui/core/styles";
import Avatar from "@material-ui/core/Avatar";
import { uniq } from "ramda";

import CustomDrawer from "./CustomDrawer";
import {
  drawerStyles,
  permissionsDict,
  getDate,
  ClientType,
} from "../../utils";
import Chip from "../../Chip";

const useStyles = makeStyles(theme => ({
  ...drawerStyles,
  cardsWrapperGrid: {
    display: "grid",
    gridTemplateColumns: "1fr 1fr 1fr",
    gridColumnGap: 16,
    "& > div": {
      marginRight: 0,
    },
  },
  button: {
    width: "100%",
    "&:first-of-type": {
      marginRight: 16,
    },
    ...theme.custom.button,
    color: "#626576",
    "&:disabled": {
      backgroundColor: "#626576 !important",
    },
  },
  alertRoot: {
    backgroundColor: "#FFE3E6",
    border: "1px solid rgba(189, 39, 30, 0.3)",
    ...theme.custom.body2,
  },
  alertIcon: {
    position: "relative",
    top: 2,
  },
  revokeInfo: {
    fontSize: 16,
    lineHeight: "24px",
    margin: "32px 0",
  },
  revokeInfoCheckbox: {
    display: "flex",
    alignItems: "center",
    "& > span": {
      marginRight: 3,
    },
  },
}));

interface Props {
  drawerData: ClientType["consents"][0];
  setDrawerData: (data: string | null) => void;
}

function AccountAccessDrawer({ drawerData, setDrawerData }: Props) {
  const classes = useStyles();

  const permissionDates = {
    Authorised: getDate(drawerData?.created_at),
    "Last updated": getDate(drawerData?.updated_at),
    "Active until": getDate(drawerData?.expires_at),
  };

  const clusters = uniq(
    drawerData?.permissions?.map(v => permissionsDict[v].Cluster) ?? []
  );

  const permissionItems = clusters.map(cluster => ({
    title: cluster,
    items: Object.values(permissionsDict)
      .filter(p => p.Cluster === cluster)
      .map(v => v.Language),
  }));

  const status = drawerData?.status as any;

  return (
    <CustomDrawer
      header={
        <div className={classes.headerContent}>
          <Avatar
            variant="square"
            className={classes.logo}
            style={{ backgroundColor: "white", color: "#626576" }}
          >
            {drawerData?.Client?.name[0]?.toUpperCase()}
          </Avatar>
          <h3 className={classes.name}>{drawerData?.Client?.name}</h3>
          <div style={{ flex: 1 }} />
          <Chip type={status && status.toLowerCase()}>{status}</Chip>
        </div>
      }
      setDrawerData={setDrawerData}
    >
      <div>
        <div className={classes.subHeader}>Permission dates</div>
        <div className={classes.cardsWrapperGrid}>
          {Object.entries(permissionDates).map(([key, value]: any) => (
            <div className={classes.card} key={key}>
              <div className={classes.cardTitle}>{key}</div>
              <div className={classes.cardContent}>{value}</div>
            </div>
          ))}
        </div>
      </div>

      <div>
        <div className={classes.subHeader}>Accounts</div>
        <div className={classes.cardsWrapperGrid}>
          {drawerData?.account_ids?.map(id => (
            <div className={classes.card} key={id}>
              <div className={classes.cardContent}>{id}</div>
            </div>
          ))}
        </div>
      </div>

      <div>
        <div className={classes.subHeader}>Details being shared</div>
        <div>
          {permissionItems.map(v => (
            <div key={v.title}>
              <div className={classes.detailsTitle}>{v.title}</div>
              <ul className={classes.ulList}>
                {v.items.map(item => (
                  <li key={item}>{item}</li>
                ))}
              </ul>
            </div>
          ))}
        </div>
      </div>
    </CustomDrawer>
  );
}

export default AccountAccessDrawer;
