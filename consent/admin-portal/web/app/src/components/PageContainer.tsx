import React from "react";
import Container from "@material-ui/core/Container";
import PageContent from "./PageContent";

export default function PageContainer ({children, fixed = true, withBackground = false, style = {}}) {

  return (
    <PageContent withBackground={withBackground} style={style}>
      <Container fixed={fixed}>
        {children}
      </Container>
    </PageContent>
  )
};
