export function field(fieldName) {
    return {
        field: fieldName,
    };
}
export function fields(fieldNames) {
    return {
        oneOf: fieldNames,
    };
}
export function anyField() {
    return {
        any: true,
    };
}
export function conditionFieldIs(condition, fieldName) {
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
export function conditionField(condition) {
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
export function conditionsToKeyValue(conditions) {
    if (!conditions) {
        return {};
    }
    const object = {};
    for (const condition of conditions) {
        object[conditionField(condition)] = conditionValue(condition);
    }
    return object;
}
export function conditionValue(condition) {
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
export function queryHasFieldCondition(query, fieldName) {
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
export function equal(selector, value) {
    return {
        equal: {
            selector,
            value,
        },
    };
}
export function startsWith(selector, value) {
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
export function contains(selector, value) {
    return {
        contains: {
            selector,
            value,
        },
    };
}
export function all() {
    return {
        all: {},
    };
}
export function id(id) {
    return equal(field("id"), id);
}
export function userId(id) {
    return equal(field("userId"), id);
}
