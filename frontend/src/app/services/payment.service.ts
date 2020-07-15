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

  public pay(ID: string, cartId: string, amount: number) {
    const httpOptions = {
      headers: new HttpHeaders({
        'Content-Type': 'application/json'
      })
    }

    const body = {
      "ID": cartId+cartId,
      "cartId": cartId,
      "Amount": amount
    }

    return this.http.post(`/server/api/v1/payment/pay/${cartId}`, body, httpOptions)
  }

  public applyDiscount(userId: string){
    const httpOptions = {
      headers: new HttpHeaders({
        'Content-Type': 'application/json'
      })
    }

    return this.http.post(`/server/api/v1/coup/${userId}`, httpOptions)
  }
}
