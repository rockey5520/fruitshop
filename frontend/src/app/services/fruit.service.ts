import { Injectable } from '@angular/core';
import { FruitModel } from './../models/fruit.model';
import { throwError } from 'rxjs';
import { HttpClient, HttpHeaders, HttpErrorResponse } from '@angular/common/http';

import { Observable, of } from 'rxjs';
import { catchError, map, tap } from 'rxjs/operators';


@Injectable({
  providedIn: 'root'

})
export class FruitService {

  constructor(private http: HttpClient) { }

 

  getFruits(): Observable<FruitModel[]> {
    return this.http.get<FruitModel[]>("/server/fruits").pipe(tap(_ => this.log('fetched fruits')),
      catchError(this.handleError<FruitModel[]>('data', []))
    );

  }

  /**
   * Handle Http operation that failed.
   * Let the app continue.
   * @param operation - name of the operation that failed
   * @param result - optional value to return as the observable result
   */
  private handleError<T>(operation = 'operation', result?: T) {
    return (error: any): Observable<T> => {

      // TODO: send the error to remote logging infrastructure
      console.error("error from handleError ",error); // log to console instead

      // TODO: better job of transforming error for user consumption
      this.log(`${operation} failed: ${error.message}`);

      // Let the app keep running by returning an empty result.
      return of(result as T);
    };
  }
  /** Log a HeroService message with the MessageService */
  private log(message: string) {
    console.log("from log",message)
    //this.messageService.add(`HeroService: ${message}`);
  }
  
}


