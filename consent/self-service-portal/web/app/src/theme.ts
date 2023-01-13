import { createTheme, Theme } from "@mui/material/styles";

declare module "@mui/material/styles" {
  interface Theme {
    custom: {
      heading2: {
        fontWeight: React.CSSProperties["fontWeight"];
        fontSize: React.CSSProperties["fontSize"];
        lineHeight: React.CSSProperties["lineHeight"];
      };
      heading6: {
        fontWeight: React.CSSProperties["fontWeight"];
        fontSize: React.CSSProperties["fontSize"];
        lineHeight: React.CSSProperties["lineHeight"];
        color: React.CSSProperties["color"];
      };
      label: {
        fontWeight: React.CSSProperties["fontWeight"];
        fontSize: React.CSSProperties["fontSize"];
        lineHeight: React.CSSProperties["lineHeight"];
        color: React.CSSProperties["color"];
      };
      caption: {
        fontWeight: React.CSSProperties["fontWeight"];
        fontSize: React.CSSProperties["fontSize"];
        lineHeight: React.CSSProperties["lineHeight"];
        color: React.CSSProperties["color"];
      };
      button: {
        fontWeight: React.CSSProperties["fontWeight"];
        fontSize: React.CSSProperties["fontSize"];
        lineHeight: React.CSSProperties["lineHeight"];
        color: React.CSSProperties["color"];
        textTransform: React.CSSProperties["textTransform"];
      };
      body2: {
        fontSize: React.CSSProperties["fontSize"];
        lineHeight: React.CSSProperties["lineHeight"];
      };
    };
  }
  interface ThemeOptions extends Theme {}
}

export const theme: Theme = createTheme({
  custom: {
    heading2: {
      fontWeight: "normal",
      fontSize: 28,
      lineHeight: "40px",
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
    body2: {
      fontSize: 14,
      lineHeight: "24px",
    },
  },
  palette: {
    primary: {
      main: "#007FFF",
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
