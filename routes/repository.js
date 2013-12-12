/*
 *  User repository routes
 */

var repos = require('../lib/repos'),
    tags = require('../lib/tags');

module.exports = {
  '/v1': {
    '/repository': {
      '/:ns/:repo': {
        // delete: repos.delete,

        '/tags': {
          // get: tags.getAll,

          '/:tag': {
            // get: tags.get,
            // delete: tags.delete,
            // put: tags.put
          }
        }
      }
    }
  }
};
