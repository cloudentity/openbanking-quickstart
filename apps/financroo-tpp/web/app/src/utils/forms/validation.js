import isURL from 'validator/lib/isURL';

const ID_PATTERN = /^[a-z0-9=._-]+$/;

export const validators = {
  length: ({label}) => v => v.trim().length > 0 || `${label} is required`,
  maxLength: ({label, max = 255}) => v => v.trim().length <= max || `${label} maximum length is ${max} characters`,
  validURL: ({label}) => v => !v || isURL(v.trim(), {protocols: ['http', 'https'], require_protocol: true, require_tld: false}) || `${label} is not valid url`,
  validHttpsURL: ({label}) => v => isURL(v.trim(), {protocols: ['https'], require_protocol: true, require_tld: false}) || `${label} is not valid url with https protocol`,
  validID: ({label}) => v => ID_PATTERN.test(v.trim()) || `${label} is not valid ID. Only following characters are allowed a-z0-9=._-`,
  isJSON: ({label}) => v => {
    try {
      JSON.parse(v)
      return true;
    } catch (e) {
      return `${label} is not valid JSON object`;
    }
  },
  notUniq: ({label, options}) => v => options.indexOf(v.trim()) === -1 || `${label} given value already exists`,
};
