/*
 *  User images routes
 */

var images = require('../lib/images');

module.exports = {
  '/v1': {
    '/images': {
      '/:id': {
        '/ancestry': {
          // get: images.getAncestry
        },
        '/json': {
          // get: images.getJSON,
          // put: images.putJSON
        },
        '/layer': {
          // get: images.getLayer,
          // put: images.putLayer
        }
      }
    }
  }
};
