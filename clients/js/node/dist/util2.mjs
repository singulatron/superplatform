import axios from 'axios';

/******************************************************************************
Copyright (c) Microsoft Corporation.

Permission to use, copy, modify, and/or distribute this software for any
purpose with or without fee is hereby granted.

THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES WITH
REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF MERCHANTABILITY
AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR ANY SPECIAL, DIRECT,
INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES WHATSOEVER RESULTING FROM
LOSS OF USE, DATA OR PROFITS, WHETHER IN AN ACTION OF CONTRACT, NEGLIGENCE OR
OTHER TORTIOUS ACTION, ARISING OUT OF OR IN CONNECTION WITH THE USE OR
PERFORMANCE OF THIS SOFTWARE.
***************************************************************************** */
/* global Reflect, Promise, SuppressedError, Symbol */


function __awaiter(thisArg, _arguments, P, generator) {
    function adopt(value) { return value instanceof P ? value : new P(function (resolve) { resolve(value); }); }
    return new (P || (P = Promise))(function (resolve, reject) {
        function fulfilled(value) { try { step(generator.next(value)); } catch (e) { reject(e); } }
        function rejected(value) { try { step(generator["throw"](value)); } catch (e) { reject(e); } }
        function step(result) { result.done ? resolve(result.value) : adopt(result.value).then(fulfilled, rejected); }
        step((generator = generator.apply(thisArg, _arguments || [])).next());
    });
}

typeof SuppressedError === "function" ? SuppressedError : function (error, suppressed, message) {
    var e = new Error(message);
    return e.name = "SuppressedError", e.error = error, e.suppressed = suppressed, e;
};

function call(address_1, apiKey_1, endpoint_1, data_1) {
    return __awaiter(this, arguments, void 0, function* (address, apiKey, endpoint, data, method = "POST") {
        const url = `${address}${endpoint}`;
        const headers = {
            Authorization: `Bearer ${apiKey}`,
        };
        const config = {
            url,
            method,
            headers,
            data,
        };
        try {
            const response = yield axios(config);
            return response.data;
        }
        catch (error) {
            if (axios.isAxiosError(error)) {
                console.error("Error:", error.response ? error.response.data : error.message);
            }
            else {
                console.error("Unexpected Error:", error);
            }
            throw error;
        }
    });
}
function uuid() {
    return (generateSegment(8) +
        "-" +
        generateSegment(4) +
        "-" +
        generateSegment(4) +
        "-" +
        generateSegment(4) +
        "-" +
        generateSegment(12));
}
function generateSegment(length) {
    return Array.from({ length: length }, () => Math.floor(Math.random() * 16).toString(16)).join("");
}

export { __awaiter as _, call as c, uuid as u };
