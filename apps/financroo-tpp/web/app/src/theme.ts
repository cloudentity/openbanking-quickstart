import { createMuiTheme, Theme } from "@material-ui/core";

declare module "@material-ui/core/styles/createMuiTheme" {
  interface Theme {
    custom: {
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
    };
  }
  interface ThemeOptions {
    custom?: {
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
    };
  }
}

export const theme: Theme = createMuiTheme({
  custom: {
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
  },
  palette: {
    primary: {
      main: "#36C6AF",
    },
    secondary: {
      main: "#1F2D48",
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
