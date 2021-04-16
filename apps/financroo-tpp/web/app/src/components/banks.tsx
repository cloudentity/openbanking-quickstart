import gobank from "../assets/banks/gobank-logo.svg";
import apexfinancial from "../assets/banks/apexfinancial-logo.svg";
import chase from "../assets/banks/chase-logo.svg";
import deutcshebank from "../assets/banks/deutcshebank-logo.svg";
import hsbc from "../assets/banks/hsbc-logo.svg";
import santander from "../assets/banks/santander-logo.svg";
import gobankIcon from "../assets/banks/gobank-icon.svg";
import apexfinancialIcon from "../assets/banks/apexfinancial-icon.svg";
import chaseIcon from "../assets/banks/chase-icon.svg";

export type Bank = {
  value: string;
  disabled: boolean;
  name?: string;
  logo: string;
  icon?: string;
};

export const banks: Bank[] = [
  {
    value: "gobank",
    disabled: false,
    name: "GO Bank",
    logo: gobank,
    icon: gobankIcon,
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
  {
    value: "santander",
    disabled: true,
    logo: santander,
  },
];
