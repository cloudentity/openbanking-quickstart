import React, { useCallback, useEffect, useState } from "react";
import OutlinedInput from "@mui/material/OutlinedInput";
import InputAdornment from "@mui/material/InputAdornment";
import IconButton from "@mui/material/IconButton";
import CancelOutlined from "@mui/icons-material/CancelOutlined";
import debounce from "lodash/debounce";
import Search from "@mui/icons-material/Search";

import ClientCard from "./ClientCard";
import { ConsentStatus, getStatus } from "./utils";
import { makeStyles } from "tss-react/mui";

const useStyles = makeStyles()(theme => ({
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
    marginBottom: 16,
  },
  searchIconContainer: {
    backgroundColor: "#DC1B37",
    height: 48,
    borderRadius: "0 4px 4px 0",
    display: "flex",
    marginLeft: 2,
  },
}));

export default function ClientsList({ clients, onRevokeClient }) {
  const { classes } = useStyles();
  const [searchText, setSearchText] = useState("");
  const [debouncedSearchText, setDebouncedSearchText] = useState("");
  const [clientsWithStatus] = useState(
    clients.map(v => ({ ...v, mainStatus: getStatus(v) }))
  );
  const [filteredClients, setFilteredClients] = useState(clientsWithStatus);

  // eslint-disable-next-line react-hooks/exhaustive-deps
  const handleSearch = useCallback(
    debounce(value => {
      setDebouncedSearchText(value);
    }, 250),
    []
  );

  useEffect(() => {
    const clientsSearched =
      (debouncedSearchText &&
        clientsWithStatus.filter(v =>
          v?.client_name
            ?.toLowerCase()
            ?.includes(debouncedSearchText.toLowerCase())
        )) ||
      clientsWithStatus;
    setFilteredClients(clientsSearched);
  }, [debouncedSearchText, clientsWithStatus]);

  return (
    <div className={classes.container}>
      <div className={classes.toolbar}>
        <div>
          <OutlinedInput
            type="text"
            value={searchText}
            onChange={e => {
              setSearchText(e.target.value);
              handleSearch(e.target.value);
            }}
            placeholder="Search applications"
            style={{
              paddingRight: 2,
              height: 32,
              fontSize: 12,
              width: 208,
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
                    size="large"
                  >
                    <CancelOutlined
                      fontSize="small"
                      style={{ color: "#BDBDBD" }}
                    />
                  </IconButton>
                ) : (
                  <Search
                    style={{ color: "#A0A3B5", marginRight: 6, fontSize: 16 }}
                  />
                )}
              </InputAdornment>
            }
          />
        </div>
      </div>
      {filteredClients
        .slice()
        .sort((a, b) => {
          if (
            a.mainStatus === ConsentStatus.Active &&
            b.mainStatus === ConsentStatus.Inactive
          ) {
            return -1;
          }
          if (
            a.mainStatus === ConsentStatus.Inactive &&
            b.mainStatus === ConsentStatus.Active
          ) {
            return 1;
          }
          return String(a?.client_name ?? "").localeCompare(
            b?.client_name ?? ""
          );
        })
        .map(client => (
          <ClientCard
            key={client?.client_id}
            client={client}
            onRevokeClient={onRevokeClient}
          />
        ))}
    </div>
  );
}
