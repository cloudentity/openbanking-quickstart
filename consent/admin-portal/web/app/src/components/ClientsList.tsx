import React, { useEffect, useState } from "react";
import { makeStyles, Theme } from "@material-ui/core";
import Chip from "@material-ui/core/Chip";

import ClientCard from "./ClientCard";
import CustomDrawer from "./CustomDrawer";
import { ConsentStatus, getRawConsents } from "./utils";

const useStyles = makeStyles((theme: Theme) => ({
  container: {
    maxWidth: 850,
    margin: "32px auto",
  },
  header: {
    ...theme.custom.heading3,
  },
  subheader: {
    ...theme.custom.body2,
    paddingBottom: 16,
    borderBottom: "1px solid #ECECEC",
    marginBottom: 24,
  },
  filterTitle: {
    ...theme.custom.label,
    marginBottom: 12,
  },
  filterChips: {
    marginBottom: 24,
    "& > div": {
      marginRight: 8,
      ...theme.custom.label,
    },
  },
}));

const activeChipStyle = {
  backgroundColor: "#002D4C",
  color: "white",
};

export default function ClientsList({
  clients,
  onRevokeClient,
  onRevokeConsent,
}) {
  const classes = useStyles();
  const [drawerData, setDrawerData] = useState<any>(null);
  const [filter, setFilter] =
    useState<"authorised" | "pending" | "revoked" | "all">("all");
  const [clientsWithStatus, setClientsWithStatus] = useState(clients);

  useEffect(() => {
    const withStatus = clients.map((client) => {
      const rawConsents = getRawConsents(client?.consents ?? []);

      const status =
        (rawConsents.length === 0 && ConsentStatus.Revoked) ||
        (rawConsents.some((v) => v?.status === "AwaitingAuthorisation") &&
          ConsentStatus.Pending) ||
        rawConsents.some(
          (v) => v?.status === "Rejected" && ConsentStatus.Revoked
        ) ||
        ConsentStatus.Authorised;

      return {
        ...client,
        mainStatus: status,
      };
    });

    setClientsWithStatus(withStatus);
  }, [clients]);

  const filteredClients =
    (filter === "authorised" &&
      clientsWithStatus.filter(
        (v) => v?.mainStatus === ConsentStatus.Authorised
      )) ||
    (filter === "pending" &&
      clientsWithStatus.filter(
        (v) => v?.mainStatus === ConsentStatus.Pending
      )) ||
    (filter === "revoked" &&
      clientsWithStatus.filter(
        (v) => v?.mainStatus === ConsentStatus.Revoked
      )) ||
    clientsWithStatus;

  return (
    <div className={classes.container}>
      <div className={classes.header}>All Connected Applications</div>
      <div className={classes.subheader}>
        Manage and revoke access to connected third-party applications
      </div>
      <div>
        <div className={classes.filterTitle}>Filter by permissions:</div>
        <div className={classes.filterChips}>
          <Chip
            label="All types"
            onClick={() => setFilter("all")}
            style={filter === "all" ? activeChipStyle : {}}
          />
          <Chip
            label="Authorised"
            onClick={() => setFilter("authorised")}
            style={filter === "authorised" ? activeChipStyle : {}}
          />
          <Chip
            label="Pending"
            onClick={() => setFilter("pending")}
            style={filter === "pending" ? activeChipStyle : {}}
          />
          <Chip
            label="Revoked"
            onClick={() => setFilter("revoked")}
            style={filter === "revoked" ? activeChipStyle : {}}
          />
        </div>
      </div>
      {filteredClients
        .sort((a, b) =>
          String(a?.client_name ?? "").localeCompare(b?.client_name ?? "")
        )
        .map((client) => (
          <ClientCard
            key={client?.client_id}
            client={client}
            onClick={() => setDrawerData(client)}
          />
        ))}

      {drawerData && (
        <CustomDrawer
          data={drawerData}
          setData={setDrawerData}
          onRevokeClient={onRevokeClient}
          onRevokeConsent={onRevokeConsent}
        />
      )}
    </div>
  );
}
