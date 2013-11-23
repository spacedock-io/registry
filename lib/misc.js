var misc = exports;

exports.index = function(req, res) {
  res.send('docker registry server');
};

exports.ping = function(req, res) {
  res.send(200);
};
