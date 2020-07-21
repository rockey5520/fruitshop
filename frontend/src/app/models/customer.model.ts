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
    Cart:      Cart;
    ID:        number;
    CreatedAt: Date;
    UpdatedAt: Date;
    DeletedAt: null;
}

export interface Cart {
    ID:                        number;
    CreatedAt:                 Date;
    UpdatedAt:                 Date;
    DeletedAt:                 null;
    CustomerId:                number;
    total:                     number;
    status:                    string;
    CartItem:                  any[];
    Payment:                   Payment;
    AppliedDualItemDiscount:   any[];
    AppliedSingleItemDiscount: any[];
    AppliedSingleItemCoupon:   any[];
}

export interface Payment {
    id:     number;
    CartId: number;
    amount: number;
    string: string;
}