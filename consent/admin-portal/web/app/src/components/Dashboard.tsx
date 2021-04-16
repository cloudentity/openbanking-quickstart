import React, {useEffect, useState} from "react";
import PageToolbar from "./PageToolbar";
import MenuIcon from "@material-ui/icons/Menu";
import IconButton from "@material-ui/core/IconButton";
import Progress from "./Progress";
import Tab from "@material-ui/core/Tab";
import Hidden from "@material-ui/core/Hidden";
import Tabs from "@material-ui/core/Tabs";
import {Button, Container, Theme, Typography} from "@material-ui/core";
import {logout} from "./AuthPage";
import {makeStyles} from "@material-ui/core/styles";
import {api} from "../api/api";
import ConfirmationDialog from "./ConfirmationDialog";
import {Route} from "react-router-dom";
import noAccountEmptyState from "./no-accounts-empty-state.svg";
import ClientsList from "./ClientsList";

const useStyles = makeStyles((theme: Theme) => ({
    indicator: {
        backgroundColor: '#fff',
    }
}));

export default function Dashboard({authorizationServerURL, authorizationServerId, tenantId}) {
    const [isProgress, setProgress] = useState(true);
    const [clients, setClients] = useState<any>([]);
    const [revokeDialog, setRevokeDialog] = useState<any>(null);
    const classes = useStyles();

    useEffect(() => {
        setProgress(true);
        api.getClients()
            .then(res => setClients(res.clients || []))
            .catch(err => console.log(err))
            .finally(() => setProgress(false));
    }, []);

    const handleRevokeConsent = id => () => {
        setProgress(true);
        api.deleteConsent({id})
            .then(api.getClients)
            .then(res => setClients(res.clients || []))
            .catch(err => console.log(err))
            .finally(() => setProgress(false));
    }

    const handleRevokeClient = id => () => {
        setProgress(true);
        api.deleteClient({id})
            .then(api.getClients)
            .then(res => setClients(res.clients || []))
            .catch(err => console.log(err))
            .finally(() => setProgress(false));
    }

    return (
        <div>
            <PageToolbar>
                <Hidden mdUp>
                    <IconButton edge="start" color="inherit" aria-label="menu">
                        <MenuIcon/>
                    </IconButton>
                </Hidden>
                <Hidden smDown>
                    <Tabs
                        value={'clients'}
                        onChange={() => {
                        }}
                        classes={{
                            indicator: classes.indicator
                        }}
                        aria-label="menu tabs"
                        style={{height: 64, marginLeft: 32}}
                    >
                        <Tab label="Clients" value={'clients'}
                             style={{height: 64, textTransform: "none", minWidth: "unset"}}/>
                    </Tabs>
                </Hidden>
                <div style={{flex: 1}}/>
                <Button
                    style={{color: '#fff'}}
                    onClick={() => logout(authorizationServerURL, tenantId, authorizationServerId)}>Logout</Button>
            </PageToolbar>
            <div style={{position: "relative"}}>
                {isProgress && <Progress/>}
                {!isProgress && (
                    <>
                        {clients.length === 0 && (
                            <div style={{textAlign: "center", marginTop: 128}}>
                                <Typography variant={"h3"} style={{color: "#626576"}}>No authorized 3rd party
                                    Applications</Typography>
                                <img src={noAccountEmptyState} style={{marginTop: 64}} alt={"empty state"}/>
                            </div>
                        )}
                        {clients.length > 0 && (
                            <>
                                <div style={{
                                    background: "#F7F7F7",
                                    height: 148,
                                    display: 'flex',
                                    alignItems: 'center'
                                }}>
                                    <Container>
                                        <Typography variant={"h3"} color={"primary"}>
                                            Authorized 3rd party Applications
                                        </Typography>
                                    </Container>
                                </div>
                                <Container>
                                    <Route
                                        exact
                                        path="/"
                                        render={() => <ClientsList
                                            clients={clients}
                                            onRevokeClient={handleRevokeClient}
                                            onRevokeConsent={handleRevokeConsent}
                                        />}
                                    />
                                </Container>
                            </>
                        )}
                    </>
                )}

            </div>
            {revokeDialog && (
                <ConfirmationDialog
                    title="Revoke application access"
                    content="Are you sure you want to revoke access to your accounts for this application?"
                    confirmText="Yes, Revoke Access"
                    warningItems={[
                        'Your connected application will not be able to access your bank accounts.',
                        'You will have to authorize the application again if you would like to use it with your bank in the future.'
                    ]}
                    onCancel={() => setRevokeDialog(null)}
                    onConfirm={() => {
                        handleRevokeConsent(revokeDialog.consent_id);
                        setRevokeDialog(null);
                    }}
                />
            )}
        </div>
    )
};
