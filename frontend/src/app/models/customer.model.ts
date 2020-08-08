export interface Customer {
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
    totalsavings:              number;
    status:                    string;
    CartItem:                  null;
    Payment:                   Payment;
    AppliedDualItemDiscount:   null;
    AppliedSingleItemDiscount: null;
    AppliedSingleItemCoupon:   null;
}

export interface Payment {
    id:         number;
    customerid: number;
    cartid:     number;
    amount:     number;
    string:     string;
}
