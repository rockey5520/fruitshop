/* export class User {
    ID: string;
    userId: string;

} */

export interface CustomerModel {
    id: number;
    loginid: string;
    firstname: string;
    lastname: string;
}

export interface RootObject {
    customer: CustomerModel;
}

