type DateType = string | number | Date;

const parse = (date: DateType) => new Date(date);

const formatDateTime = (date: DateType) => parse(date).toLocaleString();

export const date = {
  formatDateTime
};
