var config = require('../config/'),
    nano   = require('nano')(config.couchdb.url);

module.exports = nano.db.use(config.couchdb.database);
