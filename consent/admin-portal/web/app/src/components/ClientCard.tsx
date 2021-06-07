import React from "react";
import { makeStyles, Theme } from "@material-ui/core/styles";
import Avatar from "@material-ui/core/Avatar";

const useStyles = makeStyles((theme: Theme) => ({
  card: {
    background: "#FFFFFF",
    padding: "12px 24px",
    boxShadow:
      "0px 1px 1px rgba(0, 0, 0, 0.08), 0px 0px 1px rgba(0, 0, 0, 0.31)",
    borderRadius: 4,
    display: "flex",
    alignItems: "center",
    cursor: "pointer",
    "&:hover": {
      boxShadow:
        "0px 1px 16px -4px rgba(0, 0, 0, 0.25), 0px 0px 1px rgba(0, 0, 0, 0.31)",
    },
  },
  avatar: {
    background: "#FCFCFF",
    border: "1px solid #626576",
    color: "#626576",
    width: 48,
    height: 48,
    marginRight: 12,
    borderRadius: 4,
  },
  name: {
    ...theme.custom.heading3,
  },
}));

export default function ClientCard({
  client,
  onClick,
  onRevokeClient,
  onRevokeConsent,
}) {
  const classes = useStyles();

  return (
    <div className={classes.card} onClick={onClick}>
      <Avatar variant="square" className={classes.avatar}>
        {client?.client_name[0]?.toUpperCase()}
      </Avatar>
      <div className={classes.name}>{client.client_name}</div>
    </div>
    // <Accordion elevation={3} style={{ marginBottom: 24 }}>
    //   <AccordionSummary
    //     classes={{ expandIcon: classes.expandIcon }}
    //     expandIcon={
    //       <Avatar>
    //         <ExpandMoreIcon />
    //       </Avatar>
    //     }
    //     aria-controls="panel1a-content"
    //     id="panel1a-header"
    //     style={{ padding: "24px 32px" }}
    //   >
    //     <div>
    //       <Avatar
    //         variant="square"
    //         style={{
    //           background: "#D5E2E5",
    //           color: "#006580",
    //           width: 44,
    //           height: 44,
    //         }}
    //       >
    //         {client?.client_name[0]?.toUpperCase()}
    //       </Avatar>
    //     </div>
    //     <div style={{ width: "100%" }}>
    //       <div style={{ display: "flex", alignItems: "center" }}>
    //         <div style={{ marginLeft: 24 }}>
    //           <Typography
    //             style={{
    //               color: "#C2C3C6",
    //               textTransform: "uppercase",
    //               fontSize: 12,
    //             }}
    //           >
    //             Client Name
    //           </Typography>
    //           <Typography style={{ fontSize: 16, marginTop: 4 }}>
    //             <strong>{client.client_name}</strong>
    //           </Typography>
    //         </div>
    //       </div>
    //       <div style={{ marginLeft: 24, marginTop: 24 }}>
    //         <Typography
    //           style={{
    //             color: "#C2C3C6",
    //             textTransform: "uppercase",
    //             fontSize: 12,
    //           }}
    //         >
    //           Client ID
    //         </Typography>
    //         <Typography style={{ fontSize: 16, marginTop: 4 }}>
    //           <strong>{client.client_id}</strong>
    //         </Typography>
    //       </div>
    //       {client.consents?.length > 0 && (
    //         <>
    //           <Divider style={{ margin: "24px 0 24px 24px" }} />
    //           <div
    //             style={{
    //               marginLeft: 24,
    //               display: "flex",
    //               justifyContent: "space-between",
    //             }}
    //           >
    //             <Typography>
    //               Revoke all the client consents which access the Financial
    //               information
    //             </Typography>
    //             <Button
    //               className={"revoke-all-button"}
    //               variant={"contained"}
    //               style={{ background: "#DC1B37", color: "#fff" }}
    //               onClick={onRevokeClient(client.client_id)}
    //             >
    //               Revoke all
    //             </Button>
    //           </div>
    //         </>
    //       )}
    //     </div>
    //   </AccordionSummary>
    //   <AccordionDetails
    //     style={{
    //       flexDirection: "column",
    //       background: "#F4F7F8",
    //       padding: 0,
    //     }}
    //   >
    //     {client.consents?.map((consent) => (
    //       <div
    //         id={consent.consent_id}
    //         style={{ paddingBottom: 32, borderLeft: "6px solid #006580" }}
    //         key={consent.consent_id}
    //       >
    //         <div
    //           style={{
    //             background: "#E4EEF0",
    //             color: "#006580",
    //             padding: "24px 0 24px 88px",
    //           }}
    //         >
    //           <Typography display={"block"}>
    //             Consent ID: <strong>{consent.consent_id}</strong>
    //           </Typography>
    //           <Typography display={"block"} style={{ marginTop: 16 }}>
    //             Granted on:{" "}
    //             <strong>{consent.creation_date_time?.split("T")[0]}</strong>
    //           </Typography>
    //           <Typography display={"block"} style={{ marginTop: 16 }}>
    //             Account ID: <strong>{consent.account_ids?.join(", ")}</strong>
    //           </Typography>
    //         </div>
    //         <div style={{ paddingLeft: 88, color: "#006580" }}>
    //           <Typography
    //             style={{
    //               marginTop: 24,
    //               textTransform: "uppercase",
    //             }}
    //           >
    //             <strong>Granted permissions:</strong>
    //           </Typography>
    //           <Divider style={{ margin: "24px 0" }} />
    //           <Grid container spacing={3}>
    //             {consent.permissions?.map((permission) => (
    //               <>
    //                 <Grid
    //                   item
    //                   xs={4}
    //                   style={{
    //                     textAlign: "left",
    //                     wordWrap: "break-word",
    //                   }}
    //                 >
    //                   <Typography>
    //                     <strong>{permission}</strong>
    //                   </Typography>
    //                 </Grid>
    //                 <Grid
    //                   item
    //                   xs={8}
    //                   style={{
    //                     textAlign: "left",
    //                     wordWrap: "break-word",
    //                   }}
    //                 >
    //                   <Typography variant={"caption"}>
    //                     {permissionNameDescriptionMap[permission]}
    //                   </Typography>
    //                 </Grid>
    //               </>
    //             ))}
    //           </Grid>
    //           <Divider style={{ margin: "24px 0" }} />
    //           <Button
    //             variant={"outlined"}
    //             color={"primary"}
    //             className={"revoke-button"}
    //             onClick={onRevokeConsent(consent.consent_id)}
    //           >
    //             Revoke Access
    //           </Button>
    //         </div>
    //       </div>
    //     ))}
    //   </AccordionDetails>
    // </Accordion>
  );
}
