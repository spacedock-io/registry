var misc = exports;

exports.index = function(req, res) {
  res.send('docker registry server (' + (process.env['SPACEDOCK_REGISTRY_ENV'] || 'staging') + ')');
};

exports.ping = function(req, res) {
  res.send(200);
};
