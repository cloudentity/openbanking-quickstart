import React from 'react';
import ClientCard from "./ClientCard";

export default function ClientsList({clients, onRevokeClient, onRevokeConsent}) {
    return (
        <div style={{marginTop: 32}}>
            {clients
                .sort((a, b) => ("" + a.client_name).localeCompare(b.client_name))
                .map(client => (
                    <ClientCard client={client} onRevokeClient={onRevokeClient} onRevokeConsent={onRevokeConsent}/>
                ))}
        </div>
    );
}
