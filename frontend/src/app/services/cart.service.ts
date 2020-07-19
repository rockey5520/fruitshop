import { OrderedFruitModel } from './../models/orderedfruit.mode';
import { CartItem } from './../models/cartitem.model';
import { Injectable, ɵɵresolveBody } from '@angular/core';
import { HttpClient } from '@angular/common/http';
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
  public getCartByID(ID: String): Observable<CartItem> {
    return this.http.get<CartItem>(`/server/api/v1/cartitem/${ID}`)
    
  }

  public addToCart(ID: String, fruit: FruitModel, count: number) {
    const httpOptions = {
      headers: new HttpHeaders({
        'Content-Type': 'application/json'
      })
    }

    const body = {
      "Name": fruit.name,
      "Count": count
    }

    return this.http.post(`/server/api/v1/cartitem/${ID}`, body, httpOptions)
  }
}
