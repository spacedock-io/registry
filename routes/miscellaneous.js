/*
 *  Miscellaneous routes
 */

var misc = require('../lib/misc');

module.exports = {
  '/': {
    get: misc.index
  },
  '/_ping': {
    get: misc.ping
  },
  '/v1': {
    '/_ping': {
      get: misc.ping
    }
  }
};
