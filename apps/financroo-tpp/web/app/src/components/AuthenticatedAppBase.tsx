import React, { Suspense } from "react";
import { Route, Routes } from "react-router-dom";

import Dashboard from "./Dashboard";
import Investments from "./investments/Investments";
import InvestmentsContribute from "./investments/InvestmentsContribute";
import InvestmentsContributeSuccess from "./investments/InvestmentsContributeSuccess";

export default function AuthenticatedAppBase() {
  return (
    <Suspense>
      <Routes>
        <Route path="/" element={<Dashboard />} />
        <Route path="/investments" element={<Investments />} />
        <Route
          path="/investments/contribute"
          element={<InvestmentsContribute />}
        />
        <Route
          path="/investments/contribute/:id/success"
          element={<InvestmentsContributeSuccess />}
        />
      </Routes>
    </Suspense>
  );
}
