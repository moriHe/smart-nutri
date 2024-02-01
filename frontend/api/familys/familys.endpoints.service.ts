import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Response } from 'api';
import { environment } from 'src/environments/environment.development';

@Injectable({
  providedIn: 'root'
})
export class FamilysEndpointsService {
  postFamily(name: string) {
    return this.http.post<Response<string>>(`${environment.backendBaseUrl}/familys`, {name})
  }

  constructor(private http: HttpClient) { }
}
