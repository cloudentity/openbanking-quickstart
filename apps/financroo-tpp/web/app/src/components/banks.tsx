import gobank from "../assets/banks/gobank-logo.svg";
import apexfinancial from "../assets/banks/apexfinancial-logo.svg";
import chase from "../assets/banks/chase-logo.svg";
import deutcshebank from "../assets/banks/deutcshebank-logo.svg";
import hsbc from "../assets/banks/hsbc-logo.svg";
import santander from "../assets/banks/santander-logo.svg";
import santanderIcon from "../assets/banks/santander-icon.svg";
import gobankIcon from "../assets/banks/gobank-icon.svg";
import apexfinancialIcon from "../assets/banks/apexfinancial-icon.svg";
import chaseIcon from "../assets/banks/chase-icon.svg";
import requestAccessPermissionsUK from "./request-uk-access-permissions.json";
import requestAccessPermissionsBR from "./request-br-access-permissions.json";

export type Permission = {
  title: string;
  value: string;
  description?: string;
};

export type Bank = {
  value: string;
  disabled: boolean;
  name?: string;
  logo: string;
  icon?: string;
  permissions?: Permission[];
};

// TODO: AUT-5813
export const banks: Bank[] = [
  {
    value: "gobank",
    disabled: false,
    name: "GO Bank",
    logo: gobank,
    icon: gobankIcon,
    permissions:
      window.spec === "obuk"
        ? requestAccessPermissionsUK.permissions
        : requestAccessPermissionsBR.permissions,
  },
  {
    value: "santander",
    name: "Santander",
    disabled: true,
    logo: santander,
    icon: santanderIcon,
  },
  {
    value: "apexfinancial",
    disabled: true,
    name: "Apex Financial",
    logo: apexfinancial,
    icon: apexfinancialIcon,
  },
  {
    value: "chase",
    disabled: true,
    logo: chase,
    icon: chaseIcon,
  },
  {
    value: "deutcshebank",
    disabled: true,
    logo: deutcshebank,
  },
  {
    value: "hsbc",
    disabled: true,
    logo: hsbc,
  },
];
