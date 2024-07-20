'use strict';

var generic = require('./generic.js');
var user = require('./user.js');



exports.all = generic.all;
exports.anyField = generic.anyField;
exports.conditionField = generic.conditionField;
exports.conditionFieldIs = generic.conditionFieldIs;
exports.conditionValue = generic.conditionValue;
exports.conditionsToKeyValue = generic.conditionsToKeyValue;
exports.contains = generic.contains;
exports.equal = generic.equal;
exports.field = generic.field;
exports.fields = generic.fields;
exports.id = generic.id;
exports.orderByField = generic.orderByField;
exports.orderByRandom = generic.orderByRandom;
exports.queryHasFieldCondition = generic.queryHasFieldCondition;
exports.startsWith = generic.startsWith;
exports.userId = generic.userId;
exports.RoleAdmin = user.RoleAdmin;
exports.RoleUser = user.RoleUser;
