/* export class User {
    ID: string;
    userId: string;

} */

/* export interface CustomerModel {
    id: number;
    loginid: string;
    firstname: string;
    lastname: string;
}

export interface RootObject {
    customer: CustomerModel;
}
 */

export interface Customer {
    data: Data;
}

export interface Data {
    loginid:   string;
    firstname: string;
    lastname:  string;
    Discounts: null;
    Cart:      Cart;
    ID:        number;
    CreatedAt: Date;
    UpdatedAt: Date;
    DeletedAt: null;
}

export interface Cart {
    id:         number;
    CustomerId: number;
    total:      number;
    CartItem:   null;
    Coupon:     Coupon;
    Payment:    Payment;
}

export interface Coupon {
    id:     number;
    cartid: number;
    name:   string;
    status: string;
}

export interface Payment {
    id:     number;
    CartId: number;
    amount: number;
    string: string;
}

// Converts JSON strings to/from your types
// and asserts the results of JSON.parse at runtime
export class Convert {
    public static toWelcome(json: string): Customer {
        return cast(JSON.parse(json), r("Welcome"));
    }

    public static welcomeToJson(value: Customer): string {
        return JSON.stringify(uncast(value, r("Welcome")), null, 2);
    }
}

function invalidValue(typ: any, val: any): never {
    throw Error(`Invalid value ${JSON.stringify(val)} for type ${JSON.stringify(typ)}`);
}

function jsonToJSProps(typ: any): any {
    if (typ.jsonToJS === undefined) {
        const map: any = {};
        typ.props.forEach((p: any) => map[p.json] = { key: p.js, typ: p.typ });
        typ.jsonToJS = map;
    }
    return typ.jsonToJS;
}

function jsToJSONProps(typ: any): any {
    if (typ.jsToJSON === undefined) {
        const map: any = {};
        typ.props.forEach((p: any) => map[p.js] = { key: p.json, typ: p.typ });
        typ.jsToJSON = map;
    }
    return typ.jsToJSON;
}

function transform(val: any, typ: any, getProps: any): any {
    function transformPrimitive(typ: string, val: any): any {
        if (typeof typ === typeof val) return val;
        return invalidValue(typ, val);
    }

    function transformUnion(typs: any[], val: any): any {
        // val must validate against one typ in typs
        const l = typs.length;
        for (let i = 0; i < l; i++) {
            const typ = typs[i];
            try {
                return transform(val, typ, getProps);
            } catch (_) {}
        }
        return invalidValue(typs, val);
    }

    function transformEnum(cases: string[], val: any): any {
        if (cases.indexOf(val) !== -1) return val;
        return invalidValue(cases, val);
    }

    function transformArray(typ: any, val: any): any {
        // val must be an array with no invalid elements
        if (!Array.isArray(val)) return invalidValue("array", val);
        return val.map(el => transform(el, typ, getProps));
    }

    function transformDate(val: any): any {
        if (val === null) {
            return null;
        }
        const d = new Date(val);
        if (isNaN(d.valueOf())) {
            return invalidValue("Date", val);
        }
        return d;
    }

    function transformObject(props: { [k: string]: any }, additional: any, val: any): any {
        if (val === null || typeof val !== "object" || Array.isArray(val)) {
            return invalidValue("object", val);
        }
        const result: any = {};
        Object.getOwnPropertyNames(props).forEach(key => {
            const prop = props[key];
            const v = Object.prototype.hasOwnProperty.call(val, key) ? val[key] : undefined;
            result[prop.key] = transform(v, prop.typ, getProps);
        });
        Object.getOwnPropertyNames(val).forEach(key => {
            if (!Object.prototype.hasOwnProperty.call(props, key)) {
                result[key] = transform(val[key], additional, getProps);
            }
        });
        return result;
    }

    if (typ === "any") return val;
    if (typ === null) {
        if (val === null) return val;
        return invalidValue(typ, val);
    }
    if (typ === false) return invalidValue(typ, val);
    while (typeof typ === "object" && typ.ref !== undefined) {
        typ = typeMap[typ.ref];
    }
    if (Array.isArray(typ)) return transformEnum(typ, val);
    if (typeof typ === "object") {
        return typ.hasOwnProperty("unionMembers") ? transformUnion(typ.unionMembers, val)
            : typ.hasOwnProperty("arrayItems")    ? transformArray(typ.arrayItems, val)
            : typ.hasOwnProperty("props")         ? transformObject(getProps(typ), typ.additional, val)
            : invalidValue(typ, val);
    }
    // Numbers can be parsed by Date but shouldn't be.
    if (typ === Date && typeof val !== "number") return transformDate(val);
    return transformPrimitive(typ, val);
}

function cast<T>(val: any, typ: any): T {
    return transform(val, typ, jsonToJSProps);
}

function uncast<T>(val: T, typ: any): any {
    return transform(val, typ, jsToJSONProps);
}

function a(typ: any) {
    return { arrayItems: typ };
}

function u(...typs: any[]) {
    return { unionMembers: typs };
}

function o(props: any[], additional: any) {
    return { props, additional };
}

function m(additional: any) {
    return { props: [], additional };
}

function r(name: string) {
    return { ref: name };
}

const typeMap: any = {
    "Welcome": o([
        { json: "data", js: "data", typ: r("Data") },
    ], false),
    "Data": o([
        { json: "loginid", js: "loginid", typ: "" },
        { json: "firstname", js: "firstname", typ: "" },
        { json: "lastname", js: "lastname", typ: "" },
        { json: "Discounts", js: "Discounts", typ: null },
        { json: "Cart", js: "Cart", typ: r("Cart") },
        { json: "ID", js: "ID", typ: 0 },
        { json: "CreatedAt", js: "CreatedAt", typ: Date },
        { json: "UpdatedAt", js: "UpdatedAt", typ: Date },
        { json: "DeletedAt", js: "DeletedAt", typ: null },
    ], false),
    "Cart": o([
        { json: "id", js: "id", typ: 0 },
        { json: "CustomerId", js: "CustomerId", typ: 0 },
        { json: "total", js: "total", typ: 0 },
        { json: "CartItem", js: "CartItem", typ: null },
        { json: "Coupon", js: "Coupon", typ: r("Coupon") },
        { json: "Payment", js: "Payment", typ: r("Payment") },
    ], false),
    "Coupon": o([
        { json: "id", js: "id", typ: 0 },
        { json: "cartid", js: "cartid", typ: 0 },
        { json: "name", js: "name", typ: "" },
        { json: "status", js: "status", typ: "" },
    ], false),
    "Payment": o([
        { json: "id", js: "id", typ: 0 },
        { json: "CartId", js: "CartId", typ: 0 },
        { json: "amount", js: "amount", typ: 0 },
        { json: "string", js: "string", typ: "" },
    ], false),
};