var misc   = require('../lib/misc'),
    images = require('../lib/images'),
    tags   = require('../lib/tags'),
    repos  = require('../lib/repos');

var routes = [
  [ 'get', '/'                                   , misc.index         ],

  /*
   * Registry routes
   */

  [ 'get', '/_ping'                              , misc.ping          ],
  [ 'get', '/v1/_ping'                           , misc.ping          ]

//[ 'get', '/v1/images/:id/layer'                , images.getLayer    ],
//[ 'put', '/v1/images/:id/layer'                , images.putLayer    ],
//[ 'get', '/v1/images/:id/json'                 , images.getJson     ],
//[ 'put', '/v1/images/:id/json'                 , images.putJson     ],
//[ 'get', '/v1/images/:id/ancestry'             , images.getAncestry ],
//
//[ 'get', '/v1/repositories/:ns/:repo/tags'     , tags.getAll        ],
//[ 'get', '/v1/repositories/:ns/:repo/tags/:tag', tags.get           ],
//[ 'del', '/v1/repositories/:ns/:repo/tags/:tag', tags.delete        ],
//[ 'put', '/v1/repositories/:ns/:repo/tags/:tag', tags.put           ],
//
//[ 'del', '/v1/repositories/:ns/:repo'          , repos.delete       ]

  /*
   * Index routes
   */
];

exports.hookRoutes = function(server) {
  routes.forEach(function(what) {
    server[what.shift()].apply(server, what);
  });
};
