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
export class CartService {
  public update;


  constructor(private http: HttpClient) {
    this.update = new BehaviorSubject<boolean>(false);
  }
  public getCartByID(ID: String): Observable<CartModel[]> {

    return this.http.get<CartModel[]>(`/server/api/v1/cart/${ID}`);
  }

  public addToCart(ID: String, fruit: FruitModel, count: number) {
    const httpOptions = {
      headers: new HttpHeaders({
        'Content-Type': 'application/json'
      })
    }

    const body = {
      "cartId": ID,
      "Name": fruit.name,
      "Count": count
    }

    return this.http.post(`/server/api/v1/cart/add/${ID}`, body, httpOptions)
  }
}
