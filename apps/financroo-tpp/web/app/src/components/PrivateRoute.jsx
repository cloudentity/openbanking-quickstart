import React, { Fragment } from "react";
import { Redirect, Route } from "react-router";
import { isTokenInStore } from "./auth.utils";

export default function PrivateRoute({ component: Component, ...rest }) {
  return (
    <Fragment>
      <Route
        {...rest}
        render={props =>
          isTokenInStore() ? (
            <Component {...props} />
          ) : (
            <Redirect
              to={{
                pathname: "/auth",
                state: { from: props.location },
              }}
            />
          )
        }
      />
    </Fragment>
  );
}
