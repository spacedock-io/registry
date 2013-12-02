var db     = require('../db/'),
    logger = require('../log/'),
    tags   = exports;

tags.setProperties = function(req, res) {
  var docId = req.params.ns + '/' + req.params.repo;
  logger.info('tags.setProperties', req.auth);

  if(!req.body) return res.send(400);

  db.get(docId, function(err, body) {
    body.access = req.body.access;
    db.insert(docId, body, function(err) {
      res.send(err ? 500 : 200);
    });
  });
};

tags.getProperties = function(req, res) {
  logger.info('tags.getProperties', req.auth);

  db.get(req.params.ns + '/' + req.params.repo, function(err, body) {
    if(err) return res.send(500);

    res.json(200, {
      access: body.access
    });
  });
};

tags.getTags = function(req, res) {

};

tags.getTag = function(req, res) {

};

tags.putTag = function(req, res) {

};

tags.deleteTag = function(req, res) {

};
