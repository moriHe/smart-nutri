import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Response } from 'api';

@Injectable({
  providedIn: 'root'
})
export class FamilysEndpointsService {
  postFamily(name: string) {
    return this.http.post<Response<string>>('http://localhost:8080/familys', {name})
  }

  constructor(private http: HttpClient) { }
}
