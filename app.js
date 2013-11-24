var express = require('express'),
    pkg     = require('./package.json'),
    mkdirp  = require('mkdirp'),
    uuid    = require('node-uuid'),
    config  = require('./config/'),
    routes  = require('./routes/');

/*
 * Setup the path to save files to  
 */

mkdirp(config.diskPath);

/*
 * Setup express app
 */

var server = express();

server.use(express.urlencoded());
server.use(express.json());
server.use(express.cookieParser());
server.use(express.session({ secret: uuid.v4() }));

server.use(function(req, res, next) {
  res.header('X-Docker-Registry-Version', pkg.version);
  res.header('X-Docker-Registry-Config', config.env);
  next();
});

server.use(function(req, res, next) {
  req.uuid = uuid.v4(); // TRACK ALL THE PEOPLE
  next();
});

/*
 * Setup routes
 */

routes.hookRoutes(server);

/*
 * Setup webserver
 */

server.listen(config.port);
