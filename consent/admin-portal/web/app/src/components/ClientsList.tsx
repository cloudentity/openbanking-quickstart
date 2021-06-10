import React, { useCallback, useEffect, useState } from "react";
import { makeStyles, Theme } from "@material-ui/core";
import Chip from "@material-ui/core/Chip";
import OutlinedInput from "@material-ui/core/OutlinedInput";
import InputAdornment from "@material-ui/core/InputAdornment";
import IconButton from "@material-ui/core/IconButton";
import CancelOutlined from "@material-ui/icons/CancelOutlined";
import debounce from "lodash/debounce";

import ClientCard from "./ClientCard";
import CustomDrawer from "./ApplicationDrawer";
import { ConsentStatus } from "./utils";

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
  toolbar: {
    display: "flex",
    justifyContent: "space-between",
  },
  searchIconContainer: {
    backgroundColor: "#DC1B37",
    height: 48,
    borderRadius: "0 4px 4px 0",
    display: "flex",
    marginLeft: 2,
  },
}));

const activeChipStyle = {
  backgroundColor: "#002D4C",
  color: "white",
};

function enrichClientWithStatus(clients) {
  return clients.map((client) => {
    const found = client?.consents?.find(
      (consent) =>
        consent.account_access_consent &&
        consent.account_access_consent.Status === "Authorised"
    );
    const status = found ? ConsentStatus.Active : ConsentStatus.Inactive;
    return {
      ...client,
      mainStatus: status,
    };
  });
}

export default function ClientsList({ clients, onRevokeClient }) {
  const classes = useStyles();
  const [drawerData, setDrawerData] = useState<any>(null);
  const [filter, setFilter] = useState<"active" | "inactive" | "all">("all");
  const [searchText, setSearchText] = useState("");
  const [debouncedSearchText, setDebouncedSearchText] = useState("");
  const [clientsWithStatus] = useState(enrichClientWithStatus(clients));
  const [filteredClients, setFilteredClients] = useState(clientsWithStatus);

  // eslint-disable-next-line react-hooks/exhaustive-deps
  const handleSearch = useCallback(
    debounce((value) => {
      setDebouncedSearchText(value);
    }, 250),
    []
  );

  useEffect(() => {
    const clientsFiltered =
      (filter === "active" &&
        clientsWithStatus.filter(
          (v) => v?.mainStatus === ConsentStatus.Active
        )) ||
      (filter === "inactive" &&
        clientsWithStatus.filter(
          (v) => v?.mainStatus === ConsentStatus.Inactive
        )) ||
      clientsWithStatus;

    const clientsSearched =
      (debouncedSearchText &&
        clientsFiltered.filter((v) =>
          v?.client_name
            ?.toLowerCase()
            ?.includes(debouncedSearchText.toLowerCase())
        )) ||
      clientsFiltered;

    setFilteredClients(clientsSearched);
  }, [filter, debouncedSearchText, clientsWithStatus]);

  return (
    <div className={classes.container}>
      <div className={classes.header}>All Connected Applications</div>
      <div className={classes.subheader}>
        Manage and revoke access to connected third-party applications
      </div>
      <div className={classes.toolbar}>
        <div>
          <div className={classes.filterTitle}>Filter by permissions:</div>
          <div className={classes.filterChips}>
            <Chip
              label="All types"
              onClick={() => setFilter("all")}
              style={filter === "all" ? activeChipStyle : {}}
            />
            <Chip
              label="Active"
              onClick={() => setFilter("active")}
              style={filter === "active" ? activeChipStyle : {}}
            />
            <Chip
              label="Inactive"
              onClick={() => setFilter("inactive")}
              style={filter === "inactive" ? activeChipStyle : {}}
            />
          </div>
        </div>
        <div>
          <OutlinedInput
            type="text"
            value={searchText}
            onChange={(e) => {
              setSearchText(e.target.value);
              handleSearch(e.target.value);
            }}
            placeholder="Search applications"
            style={{
              paddingRight: 2,
              height: 32,
              fontSize: 12,
              width: 200,
              marginTop: 32,
            }}
            endAdornment={
              <InputAdornment position="end">
                {searchText !== "" ? (
                  <IconButton
                    style={{ padding: 4 }}
                    onClick={() => {
                      setSearchText("");
                      setDebouncedSearchText("");
                    }}
                    onMouseDown={() => {
                      setSearchText("");
                      setDebouncedSearchText("");
                    }}
                  >
                    <CancelOutlined
                      fontSize="small"
                      style={{ color: "#BDBDBD" }}
                    />
                  </IconButton>
                ) : (
                  <div style={{ width: 25.5 }} />
                )}
              </InputAdornment>
            }
            labelWidth={0}
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
        />
      )}
    </div>
  );
}
