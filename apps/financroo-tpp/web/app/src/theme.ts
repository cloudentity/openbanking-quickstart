import { createTheme, Theme } from "@material-ui/core/styles";
import { CreateCSSProperties } from "@material-ui/core/styles/withStyles";

declare module "@material-ui/core/styles/createTheme" {
  interface Theme {
    custom: {
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
    };
  }
  interface ThemeOptions {
    custom?: {
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
    };
  }
}

export const theme: Theme = createTheme({
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
