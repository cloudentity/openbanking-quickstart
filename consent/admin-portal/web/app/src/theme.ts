import {
  unstable_createMuiStrictModeTheme as createMuiTheme,
  Theme,
} from "@material-ui/core";

declare module "@material-ui/core/styles/createMuiTheme" {
  interface Theme {
    custom: {
      heading2: {
        fontWeight: React.CSSProperties["fontWeight"];
        fontSize: React.CSSProperties["fontSize"];
        lineHeight: React.CSSProperties["lineHeight"];
      };
      heading3: {
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
      body1: {
        fontSize: React.CSSProperties["fontSize"];
        lineHeight: React.CSSProperties["lineHeight"];
      };
      body2: {
        fontSize: React.CSSProperties["fontSize"];
        lineHeight: React.CSSProperties["lineHeight"];
      };
    };
  }
  interface ThemeOptions {
    custom?: {
      heading2?: {
        fontWeight?: React.CSSProperties["fontWeight"];
        fontSize?: React.CSSProperties["fontSize"];
        lineHeight?: React.CSSProperties["lineHeight"];
      };
      heading3?: {
        fontWeight?: React.CSSProperties["fontWeight"];
        fontSize?: React.CSSProperties["fontSize"];
        lineHeight?: React.CSSProperties["lineHeight"];
      };
      heading6?: {
        fontWeight?: React.CSSProperties["fontWeight"];
        fontSize?: React.CSSProperties["fontSize"];
        lineHeight?: React.CSSProperties["lineHeight"];
        color?: React.CSSProperties["color"];
      };
      label?: {
        fontWeight?: React.CSSProperties["fontWeight"];
        fontSize?: React.CSSProperties["fontSize"];
        lineHeight?: React.CSSProperties["lineHeight"];
        color?: React.CSSProperties["color"];
      };
      caption?: {
        fontWeight?: React.CSSProperties["fontWeight"];
        fontSize?: React.CSSProperties["fontSize"];
        lineHeight?: React.CSSProperties["lineHeight"];
        color?: React.CSSProperties["color"];
      };
      button?: {
        fontWeight?: React.CSSProperties["fontWeight"];
        fontSize?: React.CSSProperties["fontSize"];
        lineHeight?: React.CSSProperties["lineHeight"];
        color?: React.CSSProperties["color"];
        textTransform?: React.CSSProperties["textTransform"];
      };
      body1?: {
        fontSize?: React.CSSProperties["fontSize"];
        lineHeight?: React.CSSProperties["lineHeight"];
      };
      body2?: {
        fontSize?: React.CSSProperties["fontSize"];
        lineHeight?: React.CSSProperties["lineHeight"];
      };
    };
  }
}

export const theme: Theme = createMuiTheme({
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
      main: "#007FFF",
    },
    secondary: {
      main: "#434656",
    },
  },
  overrides: {
    MuiTableRow: {
      root: {
        "&$selected": {
          backgroundColor: "rgba(54, 198, 175, 0.08)",
          "&:hover": {
            backgroundColor: "rgba(54, 198, 175, 0.2)",
          },
        },
      },
    },
    MuiTableCell: {
      root: {
        borderBottom: "none",
      },
    },
  },
});
