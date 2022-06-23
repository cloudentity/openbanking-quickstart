import { createTheme, Theme } from "@material-ui/core/styles";
import { CreateCSSProperties } from "@material-ui/core/styles/withStyles";

declare module "@material-ui/core/styles/createTheme" {
  interface Theme {
    custom: {
      heading2: {
        fontWeight: CreateCSSProperties["fontWeight"];
        fontSize: CreateCSSProperties["fontSize"];
        lineHeight: CreateCSSProperties["lineHeight"];
      };
      heading3: {
        fontWeight: CreateCSSProperties["fontWeight"];
        fontSize: CreateCSSProperties["fontSize"];
        lineHeight: CreateCSSProperties["lineHeight"];
      };
      heading6: {
        fontWeight: CreateCSSProperties["fontWeight"];
        fontSize: CreateCSSProperties["fontSize"];
        lineHeight: CreateCSSProperties["lineHeight"];
        color: CreateCSSProperties["color"];
      };
      label: {
        fontWeight: CreateCSSProperties["fontWeight"];
        fontSize: CreateCSSProperties["fontSize"];
        lineHeight: CreateCSSProperties["lineHeight"];
        color: CreateCSSProperties["color"];
      };
      caption: {
        fontWeight: CreateCSSProperties["fontWeight"];
        fontSize: CreateCSSProperties["fontSize"];
        lineHeight: CreateCSSProperties["lineHeight"];
        color: CreateCSSProperties["color"];
      };
      button: {
        fontWeight: CreateCSSProperties["fontWeight"];
        fontSize: CreateCSSProperties["fontSize"];
        lineHeight: CreateCSSProperties["lineHeight"];
        color: CreateCSSProperties["color"];
        textTransform: CreateCSSProperties["textTransform"];
      };
      body1: {
        fontSize: CreateCSSProperties["fontSize"];
        lineHeight: CreateCSSProperties["lineHeight"];
      };
      body2: {
        fontSize: CreateCSSProperties["fontSize"];
        lineHeight: CreateCSSProperties["lineHeight"];
      };
    };
  }
  interface ThemeOptions {
    custom?: {
      heading2?: {
        fontWeight?: CreateCSSProperties["fontWeight"];
        fontSize?: CreateCSSProperties["fontSize"];
        lineHeight?: CreateCSSProperties["lineHeight"];
      };
      heading3?: {
        fontWeight?: CreateCSSProperties["fontWeight"];
        fontSize?: CreateCSSProperties["fontSize"];
        lineHeight?: CreateCSSProperties["lineHeight"];
      };
      heading6?: {
        fontWeight?: CreateCSSProperties["fontWeight"];
        fontSize?: CreateCSSProperties["fontSize"];
        lineHeight?: CreateCSSProperties["lineHeight"];
        color?: CreateCSSProperties["color"];
      };
      label?: {
        fontWeight?: CreateCSSProperties["fontWeight"];
        fontSize?: CreateCSSProperties["fontSize"];
        lineHeight?: CreateCSSProperties["lineHeight"];
        color?: CreateCSSProperties["color"];
      };
      caption?: {
        fontWeight?: CreateCSSProperties["fontWeight"];
        fontSize?: CreateCSSProperties["fontSize"];
        lineHeight?: CreateCSSProperties["lineHeight"];
        color?: CreateCSSProperties["color"];
      };
      button?: {
        fontWeight?: CreateCSSProperties["fontWeight"];
        fontSize?: CreateCSSProperties["fontSize"];
        lineHeight?: CreateCSSProperties["lineHeight"];
        color?: CreateCSSProperties["color"];
        textTransform?: CreateCSSProperties["textTransform"];
      };
      body1?: {
        fontSize?: CreateCSSProperties["fontSize"];
        lineHeight?: CreateCSSProperties["lineHeight"];
      };
      body2?: {
        fontSize?: CreateCSSProperties["fontSize"];
        lineHeight?: CreateCSSProperties["lineHeight"];
      };
    };
  }
}

export const theme: Theme = createTheme({
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
