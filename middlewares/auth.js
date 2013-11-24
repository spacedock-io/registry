module.exports = function parseAuthToken(req, res, next) {
  if(req.headers.authorization) {
    var auth = req.headers.authorization.match(/^Token signature=(\w+),repository="(\w+)\/(\w+)",access=(\w+)$/);

    if(typeof auth === 'null') {
      return res.send(400);
    }

    req.user = {
      signature: auth[1],
      namespace: auth[2],
      repo     : auth[3],
      access   : auth[4]
    };

    return next();
  }
  else { return next(); }
};
