import { createTheme } from "@mui/material/styles";

declare module "@mui/material/styles" {
  interface Theme {
    custom: {
      heading6: {
        fontWeight: string;
        fontSize: number;
        lineHeight: string;
        color: string;
      };
      label: {
        fontWeight: string;
        fontSize: number;
        lineHeight: string;
        color: string;
      };
      caption: {
        fontWeight: string;
        fontSize: number;
        lineHeight: string;
        color: string;
      };
      button: {
        fontWeight: string;
        fontSize: number;
        lineHeight: string;
        color: string;
      };
    };
  }
  interface ThemeOptions extends Theme {}
}

const custom = {
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
};

const palette = {
  primary: {
    main: "#36C6AF",
  },
  secondary: {
    main: "#1F2D48",
  },
};

const overrides = {
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
};

export const theme = createTheme({
  custom,
  palette,
  components: overrides,
  breakpoints: {
    values: {
      xs: 0,
      sm: 480,
      md: 768,
      lg: 1280,
      xl: 1920,
    },
  },
  unstable_sx: {} as any,
});
