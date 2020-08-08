import { Customer } from './../models/customer.model';
import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { BehaviorSubject, Observable } from 'rxjs';
import { map } from 'rxjs/operators';



@Injectable({ providedIn: 'root' })
export class AuthenticationService {
    private currentUserSubject: BehaviorSubject<Customer>;
    public currentUser: Observable<Customer>;
    public update;

    constructor(private http: HttpClient) {
        this.currentUserSubject = new BehaviorSubject<Customer>(JSON.parse(localStorage.getItem('currentUser')));
        this.currentUser = this.currentUserSubject.asObservable();
        this.update = new BehaviorSubject<Boolean>(false)
    }

    public get currentUserValue(): Customer {
        return this.currentUserSubject.value;
    }

    login(loginid){
        return this.http.get<any>(`/server/customers/${loginid}`)
            .pipe(map(user => {
                // store user details and jwt token in local storage to keep user logged in between page refreshes
                localStorage.setItem('currentUser', JSON.stringify(user));
                this.currentUserSubject.next(user);
                return user;
            }));
    }

    login2(loginid){
        return this.http.get<any>(`/server/customers/${loginid}`)
    }

    logout() {
        // remove user from local storage and set current user to null
        localStorage.removeItem('currentUser');
        this.currentUserSubject.next(null);
    }
}