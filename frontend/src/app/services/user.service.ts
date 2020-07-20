import { Data } from './../models/customer.model';

import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { map } from 'rxjs/operators';
import { BehaviorSubject, Observable } from 'rxjs';



@Injectable({ providedIn: 'root' })
export class UserService {

    private currentUserSubject: BehaviorSubject<Data>;
    public currentUser: Observable<Data>;



    constructor(private http: HttpClient) { 
        this.currentUserSubject = new BehaviorSubject<Data>(JSON.parse(localStorage.getItem('currentUser')));
        this.currentUser = this.currentUserSubject.asObservable();
    }

      
    register(customer: Data) {
        console.log("customer id ",customer.loginid)
        console.log("customer firstname ",customer.firstname)
        
        return this.http.post(`/server/api/v1/customers`, customer).pipe(map(customer => {
            // store user details and jwt token in local storage to keep user logged in between page refreshes
            localStorage.setItem('currentUser', JSON.stringify(customer));
            console.log("currentUser",localStorage.getItem("currentUser"))
            //this.currentUserSubject.next(customer)
            return customer;
        }));
    }

    getUser(loginId : string){
        return this.http.get(`/server/api/v1/customers/${loginId}`)
    }
 
}