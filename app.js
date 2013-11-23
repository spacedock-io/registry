var express = require('express'),
    pkg     = require('./package.json'),
    routes  = require('./routes/');

var server = express();

server.use(express.urlencoded());
server.use(express.json());
server.use(express.cookieParser());
server.use(express.session({ secret: require('node-uuid').v4() }));

server.use(function(req, res, next) {
  res.header('X-Docker-Registry-Version', pkg.version);
  res.header('X-Docker-Registry-Config', (process.env['SPACEDOCK_REGISTRY_ENV'] || 'staging'));
  next();
});

routes.hookRoutes(server);

server.listen(8080);
