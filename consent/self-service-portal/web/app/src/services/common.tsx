import React, {createContext, FunctionComponent, useCallback, useState} from "react";

interface CommonCtx {
  error?: string;
  setError(msg: string);
  clearError();
}

export const CommonCtx = createContext<CommonCtx | undefined>(undefined);

export const CommonProvider: FunctionComponent = (props) => {
  const [error, setError] = useState<string | undefined>();
  const clearError = useCallback(() => setError(undefined), [])
  return <CommonCtx.Provider value={{error, setError, clearError}}>{props.children}</CommonCtx.Provider>
}
