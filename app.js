var express = require('express'),
    routes  = require('./routes/');

var server = express();

server.use(express.urlencoded());
server.use(express.json());
server.use(express.cookieParser());
server.use(express.session({ secret: require('node-uuid').v4() }));

routes.hookRoutes(server);

server.listen(8080);
