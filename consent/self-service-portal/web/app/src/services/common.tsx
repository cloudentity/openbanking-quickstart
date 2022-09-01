import React, {
  createContext,
  FunctionComponent,
  useCallback,
  useState,
} from "react";

interface CommonCtxType {
  error?: string;
  setError(msg: string);
  clearError();
}

export const CommonCtx = createContext<CommonCtxType | undefined>(undefined);

export const CommonProvider: FunctionComponent = props => {
  const [error, setError] = useState<string | undefined>();
  const clearError = useCallback(() => setError(undefined), []);
  return (
    <CommonCtx.Provider value={{ error, setError, clearError }}>
      {props.children}
    </CommonCtx.Provider>
  );
};
