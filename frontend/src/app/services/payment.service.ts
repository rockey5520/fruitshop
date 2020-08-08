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

  public pay(id:number, cartID: Number,amount: number) {
    const httpOptions = {
      headers: new HttpHeaders({
        'Content-Type': 'application/json'
      })
    }

    const body = {
      "customerid":id,
      "cartid":cartID,
      "amount": amount
  }

    return this.http.post(`/server/pay`, body, httpOptions)
  }

  public applyDiscount(cart_id: number, fruit_id:number){
    return this.http.get(`/server/orangecoupon/${cart_id}/${fruit_id}`)
  }
}
