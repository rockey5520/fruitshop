import { Customer } from './../models/customer.model';

import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { map } from 'rxjs/operators';
import { BehaviorSubject, Observable } from 'rxjs';



@Injectable({ providedIn: 'root' })
export class UserService {

    private currentUserSubject: BehaviorSubject<Customer>;
    public currentUser: Observable<Customer>;



    constructor(private http: HttpClient) { 
        this.currentUserSubject = new BehaviorSubject<Customer>(JSON.parse(localStorage.getItem('currentUser')));
        this.currentUser = this.currentUserSubject.asObservable();
    }

      
    register(customer: Customer) {        
        return this.http.post(`/server/customers`, customer).pipe(map(customer => {
            // store user details and jwt token in local storage to keep user logged in between page refreshes
            localStorage.setItem('currentUser', JSON.stringify(customer));
            return customer;
        }));
    }

    getUser(loginId : string){
        return this.http.get(`/server/customers/${loginId}`)
    }
 
}