var misc = require('../lib/misc');

var routes = [
  [ 'get', '/'        , misc.index ],

  [ 'get', '/_ping'   , misc.ping  ],
  [ 'get', '/v1/_ping', misc.ping  ]
];

exports.hookRoutes = function(server) {
  routes.forEach(function(what) {
    server[what.shift()].apply(server, what);
  });
};
