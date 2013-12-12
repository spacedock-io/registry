var express = require('express'),
    pkg     = require('./package.json'),
    mkdirp  = require('mkdirp'),
    uuid    = require('node-uuid'),
    auth    = require('./middlewares/auth'),
    config  = require('./config/'),
    routes  = require('./routes/');


// Setup the path to save files to
mkdirp(config.diskPath);

// Setup express app
var server = express();

server.use(express.urlencoded());
server.use(express.json());
server.use(express.cookieParser());
server.use(express.session({ secret: uuid.v4() }));

server.use(function(req, res, next) {
  res.set({
    'X-Powered-By': 'SpaceDock',
    'X-Docker-Registry-Version': pkg.version,
    'X-Docker-Registry-Config': config.env
  });
  next();
});

// Setup registry auth
server.use(auth);

// Setup registry routes
routes(server);

// Setup webserver
server.listen(config.port);
