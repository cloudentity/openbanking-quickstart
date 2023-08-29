import gobank from "../assets/banks/gobank-logo.svg";
import apexfinancial from "../assets/banks/apexfinancial-logo.svg";
import chase from "../assets/banks/chase-logo.svg";
import deutcshebank from "../assets/banks/deutcshebank-logo.svg";
import hsbc from "../assets/banks/hsbc-logo.svg";
import hyperscalebank from "../assets/banks/hyperscalebank-logo.svg";
import hyperscalebankIcon from "../assets/banks/hyperscalebank-icon.svg";
import gobankIcon from "../assets/banks/gobank-icon.svg";
import apexfinancialIcon from "../assets/banks/apexfinancial-icon.svg";
import chaseIcon from "../assets/banks/chase-icon.svg";
import unknownBankLogo from "../assets/banks/unknown-bank-logo.svg";
import unknownBankIcon from "../assets/banks/unknown-bank.svg";
import requestAccessPermissionsUK from "./request-uk-access-permissions.json";
import requestAccessPermissionsBR from "./request-br-access-permissions.json";
import { AvailableBank } from "./types";

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
    disabled: true,
    name: "GO Bank",
    logo: gobank,
    icon: gobankIcon,
    permissions:
      window.spec === "obuk"
        ? requestAccessPermissionsUK.permissions
        : requestAccessPermissionsBR.permissions,
  },
  {
    value: "hyperscalebank",
    disabled: true,
    name: "Hyperscale Bank",
    logo: hyperscalebank,
    icon: hyperscalebankIcon,
    permissions:
      window.spec === "obuk"
        ? requestAccessPermissionsUK.permissions
        : requestAccessPermissionsBR.permissions,
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

export const getUnknownBankConfig = (bank: AvailableBank) => ({
  value: bank.id,
  name: bank.name ?? "Unknown Bank",
  disabled: false,
  logo: bank.logo_url ?? unknownBankLogo,
  icon: bank.icon_url ?? unknownBankIcon,
});
