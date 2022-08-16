import React from "react";
import Container from "@mui/material/Container";
import PageContent from "./PageContent";

export default function PageContainer({
  children,
  fixed = true,
  withBackground = false,
  withSubheader = false,
  withOnlySubheader = false,
  style = {},
  containerStyle = {},
}) {
  return (
    <PageContent
      withBackground={withBackground}
      withSubheader={withSubheader}
      withOnlySubheader={withOnlySubheader}
      style={style}
    >
      <Container fixed={fixed} style={containerStyle}>
        {children}
      </Container>
    </PageContent>
  );
}
