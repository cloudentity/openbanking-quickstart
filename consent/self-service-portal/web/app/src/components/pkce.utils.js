import 'fast-text-encoding'; // IE11 and Edge Legacy require TextEncoder polyfill
const crypto = window.crypto || window.msCrypto; // IE11 uses 'msCrypto'

// https://github.com/aaronpk/pkce-vanilla-js

// Generate a secure random string using the browser crypto functions
export const generateRandomString = () => {
  const array = new Uint32Array(28);
  crypto.getRandomValues(array);
  return Array.from(array, dec => ('0' + dec.toString(16)).substr(-2)).join('');
};

// Calculate the SHA256 hash of the input text.
// Returns a promise that resolves to an ArrayBuffer
async function sha256 (plain) {
  const encoder = new TextEncoder();
  const data = encoder.encode(plain);
  if (window.CryptoOperation) {
    // in IE11, window.msCrypto.subtle.digest returns CryptoOperation instead of Promise
    return new Promise((resolve, reject) => {
      try {
        crypto.subtle.digest('SHA-256', data).oncomplete = function (e) {
          return resolve(e && e.target && e.target.result);
        }
      } catch (err) {
        return reject(err);
      }
    });
  } else {
    return crypto.subtle.digest('SHA-256', data);
  }
}

// Base64-urlencodes the input string
function base64urlencode (str) {
  // Convert the ArrayBuffer to string using Uint8 array to conver to what btoa accepts.
  // btoa accepts chars only within ascii 0-255 and base64 encodes them.
  // Then convert the base64 encoded to base64url encoded
  //   (replace + with -, replace / with _, trim trailing =)
  return btoa(String.fromCharCode.apply(null, new Uint8Array(str)))
    .replace(/\+/g, '-').replace(/\//g, '_').replace(/=+$/, '');
}

// Return the base64-urlencoded sha256 hash for the PKCE challenge
export const pkceChallengeFromVerifier = async (v) => {
  const hashed = await sha256(v);
  return base64urlencode(hashed);
};
