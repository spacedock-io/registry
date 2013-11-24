module.exports = function parseAuthToken(req, res, next) {
  if(!req.headers.authorization) return next();

  var auth = req.headers.authorization.match(/^Token signature=(\w+),repository="(\w+)\/(\w+)",access=(\w+)$/);

  if(!auth) {
    return res.send(400);
  }

  req.auth = {
    signature: auth[1],
    namespace: auth[2],
    repo     : auth[3],
    access   : auth[4]
  };

  next();
};
