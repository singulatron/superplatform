'use strict';

function field(fieldName) {
    return {
        field: fieldName,
    };
}
function fields(fieldNames) {
    return {
        oneOf: fieldNames,
    };
}
function anyField() {
    return {
        any: true,
    };
}
function conditionFieldIs(condition, fieldName) {
    var _a, _b, _c;
    if (condition.equal &&
        (condition.equal.selector.field == fieldName ||
            ((_a = condition.equal.selector.oneOf) === null || _a === void 0 ? void 0 : _a.includes(fieldName)))) {
        return true;
    }
    if (condition.contains &&
        (condition.contains.selector.field == fieldName ||
            ((_b = condition.contains.selector.oneOf) === null || _b === void 0 ? void 0 : _b.includes(fieldName)))) {
        return true;
    }
    if (condition.startsWith &&
        (condition.startsWith.selector.field == fieldName ||
            ((_c = condition.startsWith.selector.oneOf) === null || _c === void 0 ? void 0 : _c.includes(fieldName)))) {
        return true;
    }
    return false;
}
function conditionField(condition) {
    var _a, _b, _c;
    if (condition.equal) {
        return (condition.equal.selector.field ||
            ((_a = condition.equal.selector.oneOf) === null || _a === void 0 ? void 0 : _a.join(",")) ||
            "");
    }
    if (condition.contains) {
        return (condition.contains.selector.field ||
            ((_b = condition.contains.selector.oneOf) === null || _b === void 0 ? void 0 : _b.join(",")) ||
            "");
    }
    if (condition.startsWith) {
        return (condition.startsWith.selector.field ||
            ((_c = condition.startsWith.selector.oneOf) === null || _c === void 0 ? void 0 : _c.join(",")) ||
            "");
    }
    return "";
}
function conditionsToKeyValue(conditions) {
    if (!conditions) {
        return {};
    }
    const object = {};
    for (const condition of conditions) {
        object[conditionField(condition)] = conditionValue(condition);
    }
    return object;
}
function conditionValue(condition) {
    if (condition.equal) {
        return condition.equal.value;
    }
    if (condition.contains) {
        return condition.contains.value;
    }
    if (condition.startsWith) {
        return condition.startsWith.value;
    }
    return "";
}
function queryHasFieldCondition(query, fieldName) {
    if (!query.conditions) {
        return false;
    }
    for (const condition of query.conditions) {
        if (conditionFieldIs(condition, fieldName)) {
            return true;
        }
    }
    return false;
}
function equal(selector, value) {
    return {
        equal: {
            selector,
            value,
        },
    };
}
function startsWith(selector, value) {
    return {
        startsWith: {
            selector,
            value,
        },
    };
}
/* contains creates a Condition for the given fields specifed by the selector
 * eg. 'field1:~something' can be acquired by contains(field("field1"), "something")
 * 'field1,field2:~something' can be acquired by contains(fields("field1", "field2"), "something")
 */
function contains(selector, value) {
    return {
        contains: {
            selector,
            value,
        },
    };
}
function all() {
    return {
        all: {},
    };
}
function id(id) {
    return equal(field("id"), id);
}
function userId(id) {
    return equal(field("userId"), id);
}

exports.all = all;
exports.anyField = anyField;
exports.conditionField = conditionField;
exports.conditionFieldIs = conditionFieldIs;
exports.conditionValue = conditionValue;
exports.conditionsToKeyValue = conditionsToKeyValue;
exports.contains = contains;
exports.equal = equal;
exports.field = field;
exports.fields = fields;
exports.id = id;
exports.queryHasFieldCondition = queryHasFieldCondition;
exports.startsWith = startsWith;
exports.userId = userId;
