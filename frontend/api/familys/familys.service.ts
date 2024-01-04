import { Injectable } from '@angular/core';
import { Observable, map } from 'rxjs';
import { FamilysEndpointsService } from './familys.endpoints.service';
import { Response } from 'api';

@Injectable({
  providedIn: 'root'
})
export class FamilysService {
  createFamily(name: string): Observable<string> {
    return this.familysEndpoint.postFamily(name).pipe(map((response: Response<string>) => response.data))
  }
  constructor(private familysEndpoint: FamilysEndpointsService) { }
}
