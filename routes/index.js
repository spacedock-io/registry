var fs = require('fs'),
    path = require('path');

var files = fs.readdirSync(__dirname).filter(function (file) {
  return file.match(/\.js$/i) !== null && file !== path.basename(module.id);
});

function flatten(routes, start, parent) {
  Object.keys(routes).forEach(function (key) {
    var route = routes[key];
    if (typeof route === 'function') return;
    else if (typeof route === 'object' && !Array.isArray(route)) {
      if (start) {
        parent[start + key] = flatten(routes[key], start + key, parent);
        delete routes[key];
        delete parent[start];
      }
      else return routes[key] = flatten(routes[key], key, routes);
    }
  });
  return routes;
}

module.exports = function (server) {
  files.forEach(function (file) {
    var routes = flatten(require(path.join(__dirname, file)));
    Object.keys(routes).forEach(function (route) {
      Object.keys(routes[route]).forEach(function (verb) {
        server[verb.toLowerCase()].call(server, route, routes[route][verb]);
      });
    });
  });
};
