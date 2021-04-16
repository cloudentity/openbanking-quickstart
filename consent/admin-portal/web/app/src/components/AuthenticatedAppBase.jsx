import React, {Suspense} from 'react';
import {Switch} from 'react-router';
import {Route} from "react-router-dom";
import {useSilentAuthentication} from './useSilentAuthentication';
import Container from "@material-ui/core/Container";
import Dashboard from "./Dashboard";

export default function AuthenticatedAppBase({authorizationServerURL, authorizationServerId, tenantId, clientId, scopes, userinfo = {}}) {
    useSilentAuthentication(authorizationServerURL, authorizationServerId, tenantId, clientId, scopes);

    return (
        <div style={{marginTop: 64}}>
            <Suspense>
                <Switch>
                    <Route path={"/"} render={() =>
                        <Dashboard
                            authorizationServerURL={authorizationServerURL}
                            authorizationServerId={authorizationServerId}
                            tenantId={tenantId}
                            userinfo={userinfo}
                        />}/>
                </Switch>
            </Suspense>
        </div>
    )
}
