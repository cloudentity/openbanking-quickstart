import React, { createContext, ReactNode, useCallback, useState } from "react";

interface CommonCtxType {
  error?: string;
  setError(msg: string);
  clearError();
}

export const CommonCtx = createContext<CommonCtxType | undefined>(undefined);

export const CommonProvider = ({ children }: { children: ReactNode }) => {
  const [error, setError] = useState<string | undefined>();
  const clearError = useCallback(() => setError(undefined), []);
  return (
    <CommonCtx.Provider value={{ error, setError, clearError }}>
      {children}
    </CommonCtx.Provider>
  );
};
