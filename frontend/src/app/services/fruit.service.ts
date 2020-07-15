import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { FruitModel } from './../models/fruit.model';
import { Observable } from 'rxjs';


@Injectable({
  providedIn: 'root'

})
export class FruitService {

  constructor(private http: HttpClient) { }



  getFruits(): Observable<FruitModel[]> {
    return this.http.get<FruitModel[]>("/server/api/v1/fruit");
  }
}


