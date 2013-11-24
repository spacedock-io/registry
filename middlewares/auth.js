module.exports = function parseAuthToken(req, res, next) {
  if(req.headers.authorization) {
    var auth    = req.headers.authorization.split(' '),
        details = auth[1].split(',');

    if(auth[0].toLowerCase() !== 'token') {
      return res.send(400);
    }

    req.user = {
      signature: details[1].split('=').pop(),
      namespace: details[2].split('=').pop().split('/').shift(),
      repo     : details[2].split('=').pop().split('/').pop(),
      access   : details[3].split('=').pop()
    };

    return next();
  }
  else { return next(); }
};
