
import { DiscountModel } from './../models/discount..model';
import { Injectable, ɵɵresolveBody } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { CartModel } from '../models/cart.model';
import { Observable, BehaviorSubject } from 'rxjs';
import { FruitModel } from '../models/fruit.model';
import { HttpHeaders } from '@angular/common/http';
import { catchError, delay } from 'rxjs/operators';

@Injectable({
    providedIn: 'root'
})
export class DiscountService {
    public update;
    

    constructor(private http: HttpClient) {
        this.update = new BehaviorSubject<boolean>(false);
        
    }
    public getDiscountsByID(ID: Number): Observable<DiscountModel[]> {
        return this.http.get<DiscountModel[]>(`/server/discounts/${ID}`).pipe(delay(400));
    }

}
