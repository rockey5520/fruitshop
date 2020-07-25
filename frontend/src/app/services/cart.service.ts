import { OrderedFruitModel } from './../models/orderedfruit.mode';
import { CartItem } from './../models/cartitem.model';
import { Injectable, ɵɵresolveBody } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable, BehaviorSubject } from 'rxjs';
import { FruitModel } from '../models/fruit.model';
import { catchError, delay } from 'rxjs/operators';
import { HttpHeaders } from '@angular/common/http';

@Injectable({
  providedIn: 'root'
})
export class CartService {
  public update;


  constructor(private http: HttpClient) {
    this.update = new BehaviorSubject<boolean>(false);
  }
  public getCartByID(ID: number): Observable<CartItem> {
    return this.http.get<CartItem>(`/server/api/v1/cartitem/${ID}`).pipe(delay(400))
    
  }

  public addToCart(ID: Number, fruit: FruitModel, count: number) {
    const httpOptions = {
      headers: new HttpHeaders({
        'Content-Type': 'application/json'
      })
    }

    const body = {
      "cartid": ID,
      "Name": fruit.name,
      "Count": count
    }
    return this.http.post(`/server/api/v1/cartitem`, body, httpOptions)
  }
}
