import { CustomerModel } from './../models/customer.model';
import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { BehaviorSubject, Observable } from 'rxjs';
import { map } from 'rxjs/operators';



@Injectable({ providedIn: 'root' })
export class AuthenticationService {
    private currentUserSubject: BehaviorSubject<CustomerModel>;
    public currentUser: Observable<CustomerModel>;

    constructor(private http: HttpClient) {
        this.currentUserSubject = new BehaviorSubject<CustomerModel>(JSON.parse(localStorage.getItem('currentUser')));
        this.currentUser = this.currentUserSubject.asObservable();
    }

    public get currentUserValue(): CustomerModel {
        return this.currentUserSubject.value;
    }

    login(loginid) {
        console.log("success",loginid)
        console.log("localStorage.length", localStorage.length)
        console.log("localStorage.getItem",localStorage.getItem("currentUser"))
        return this.http.get<any>(`/server/api/v1/customers/${loginid}`)
            .pipe(map(user => {
                // store user details and jwt token in local storage to keep user logged in between page refreshes
                localStorage.setItem('currentUser', JSON.stringify(user));
                console.log("currentUser",localStorage.getItem("currentUser"))
                this.currentUserSubject.next(user);
                return user;
            }));
    }

    logout() {
        // remove user from local storage and set current user to null
        localStorage.removeItem('currentUser');
        this.currentUserSubject.next(null);
    }
}