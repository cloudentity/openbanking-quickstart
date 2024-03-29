import React from "react";
import { Navigate } from "react-router-dom";

import {
  getTokenFromStore,
  isTokenInStore,
  removeAllAuthDataFromStore,
} from "./auth.utils";
import Register from "./Register";

export const logout = () => {
  removeAllAuthDataFromStore();
  window.location.href = `/auth`;
};

const AuthPage = ({ loginFn }) => {
  const HandleLogin = async (login, password) => {
    return new Promise<void>((resolve, reject) => {
      if (login === "test" && password === "p@ssw0rd!") {
        loginFn({
          token:
            "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.XbPfbIHMI6arZ3Y922BhjWgQzWXcXNrz0ogtVhfEd2o",
        });

        window.location.href = "/";
        resolve();
        return;
      }

      reject();
    });
  };

  if (isTokenInStore()) {
    loginFn({ token: getTokenFromStore() });
    return <Navigate to="/" replace />;
  }

  return <Register onLogin={HandleLogin} />;
};

export default AuthPage;
