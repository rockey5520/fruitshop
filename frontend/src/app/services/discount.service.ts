import { DiscountModel } from './../models/discount..model';
import { Injectable, ɵɵresolveBody } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { CartModel } from '../models/cart.model';
import { Observable, BehaviorSubject } from 'rxjs';
import { FruitModel } from '../models/fruit.model';
import { catchError } from 'rxjs/operators';
import { HttpHeaders } from '@angular/common/http';

@Injectable({
    providedIn: 'root'
})
export class DiscountService {
    public update;

    constructor(private http: HttpClient) {
        this.update = new BehaviorSubject<boolean>(false);
    }
    public getDiscountsByID(ID: String): Observable<DiscountModel[]> {
        console.log("discount ID", ID)
        return this.http.get<DiscountModel[]>(`/server/api/v1/discount/${ID}`);
    }

}
