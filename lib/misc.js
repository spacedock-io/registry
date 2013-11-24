var config = require('../config/'),
    misc   = exports;

exports.index = function(req, res) {
  res.send('docker registry server (' + config.env + ')');
};

exports.ping = function(req, res) {
  res.send(200);
};
