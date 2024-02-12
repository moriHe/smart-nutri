import { Injectable } from '@angular/core';
import { Observable, map } from 'rxjs';
import { Response } from 'api';
import { environment } from 'src/environments/environment.development';
import { HttpClient } from '@angular/common/http';

@Injectable({
  providedIn: 'root'
})
export class FamilysService {
  createFamily(name: string): Observable<string> {
    return this.http.post<Response<string>>(`${environment.backendBaseUrl}/familys`, {name})
      .pipe(
        map((response: Response<string>) => response.data))
  }
  constructor(private http: HttpClient) {}
}
