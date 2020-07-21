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
export class PaymentService {
  public update;


  constructor(private http: HttpClient) {
    this.update = new BehaviorSubject<boolean>(false);
  }

  public pay( loginid:String, cartID: Number,amount: number) {
    const httpOptions = {
      headers: new HttpHeaders({
        'Content-Type': 'application/json'
      })
    }

    const body = {
      "customerid":loginid,
      "CartID":cartID,
      "amount": 19.0
  }

    return this.http.post(`/server/api/v1/pay`, body, httpOptions)
  }

  public applyDiscount(userId: number){
    return this.http.get(`/server/api/v1/orangecoupon/${userId}`)
  }
}
