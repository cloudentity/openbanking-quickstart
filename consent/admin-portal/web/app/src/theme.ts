import { createTheme } from "@mui/material/styles";
import { CSSProperties } from "react";

declare module "@mui/material/styles" {
  interface Theme {
    custom: {
      heading2: {
        fontWeight: string | number;
        fontSize: number;
        lineHeight: string;
      };
      heading3: {
        fontWeight: string | number;
        fontSize: number;
        lineHeight: string;
      };
      heading6: {
        fontWeight: string | number;
        fontSize: number;
        lineHeight: string;
        color: string;
      };
      label: {
        fontWeight: string | number;
        fontSize: number;
        lineHeight: string;
        color: string;
      };
      caption: {
        fontWeight: string | number;
        fontSize: number;
        lineHeight: string;
        color: string;
      };
      button: {
        fontWeight: string | number;
        fontSize: number;
        lineHeight: string;
        color: string;
        textTransform: CSSProperties["textTransform"];
      };
      body1: {
        fontSize: number;
        lineHeight: string;
      };
      body2: {
        fontSize: number;
        lineHeight: string;
      };
    };
  }
  interface ThemeOptions extends Theme {}
}
export const theme = createTheme({
  custom: {
    heading2: {
      fontWeight: "normal",
      fontSize: 28,
      lineHeight: "40px",
    },
    heading3: {
      fontWeight: 500,
      fontSize: 20,
      lineHeight: "32px",
    },
    heading6: {
      fontWeight: "bold",
      fontSize: 12,
      lineHeight: "16px",
      color: "#626576",
    },
    label: {
      fontWeight: "bold",
      fontSize: 12,
      lineHeight: "24px",
      color: "#212533",
    },
    caption: {
      fontWeight: "normal",
      fontSize: 12,
      lineHeight: "22px",
      color: "#626576",
    },
    button: {
      fontWeight: "normal",
      fontSize: 16,
      lineHeight: "24px",
      color: "white",
      textTransform: "none",
    },
    body1: {
      fontSize: 16,
      lineHeight: "24px",
    },
    body2: {
      fontSize: 14,
      lineHeight: "24px",
    },
  },
  palette: {
    primary: {
      main: "#DC1B37",
    },
    secondary: {
      main: "#434656",
    },
  },
  components: {
    MuiTableRow: {
      styleOverrides: {
        root: {
          "&$selected": {
            backgroundColor: "rgba(54, 198, 175, 0.08)",
            "&:hover": {
              backgroundColor: "rgba(54, 198, 175, 0.2)",
            },
          },
        },
      },
    },
    MuiTableCell: {
      styleOverrides: {
        root: {
          borderBottom: "none",
        },
      },
    },
  },
  unstable_sx: {} as any,
});
