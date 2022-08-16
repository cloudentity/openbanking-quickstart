import React, { Suspense } from "react";
import { Switch } from "react-router";
import { Route } from "react-router-dom";

import Dashboard from "./Dashboard";
import Investments from "./investments/Investments";
import InvestmentsContribute from "./investments/InvestmentsContribute";
import InvestmentsContributeSuccess from "./investments/InvestmentsContributeSuccess";

export default function AuthenticatedAppBase() {
  return (
    <Suspense>
      <Switch>
        <Route exact path="/" render={() => <Dashboard />} />
        <Route path="/investments" exact render={() => <Investments />} />
        <Route
          path="/investments/contribute"
          exact
          render={() => <InvestmentsContribute />}
        />

        <Route
          path="/investments/contribute/:id/success"
          exact
          render={() => <InvestmentsContributeSuccess />}
        />
      </Switch>
    </Suspense>
  );
}
