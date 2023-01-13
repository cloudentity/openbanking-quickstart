export const getCurrency = (currency: any) => {
  const currentCurrency = window.currency || currency;
  switch (currentCurrency) {
    case "USD":
      return "$";
    case "GBP":
      return "£";
    case "EUR":
      return "€";
    case "BRL":
      return "R$";
    default:
      return currentCurrency;
  }
};
