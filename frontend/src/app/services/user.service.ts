import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';

import { User } from '../models/user.model';

@Injectable({ providedIn: 'root' })
export class UserService {
    constructor(private http: HttpClient) { }

      
    register(user: User) {
        console.log(user.userId)
        console.log(user)
        return this.http.post(`/server/api/v1/user/${user.userId}`, user);
    }

    getUser(userId : string){
        return this.http.get(`/server/api/v1/user/${userId}`)
    }
 
}