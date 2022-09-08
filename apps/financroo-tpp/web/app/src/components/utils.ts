export const getCurrency = (currency: any) => {
  switch (currency) {
    case "USD":
      return "$";
    case "GBP":
      return "£";
    case "EUR":
      return "€";
    case "BRL":
      return "R$";
    default:
      return currency;
  }
};
