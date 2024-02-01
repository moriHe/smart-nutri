import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Response } from 'api';
import { map } from 'rxjs';
import { environment } from 'src/environments/environment.development';

@Injectable({
  providedIn: 'root'
})
export class InvitationsService {

  getInvitationLink() {
    return this.http.get<Response<string>>(`${environment.backendBaseUrl}/invitations/link`)
      .pipe(map((response) => response.data))
  }

  acceptInvitation(token: string) {
    return this.http.get<Response<string>>(`${environment.backendBaseUrl}/invitations/accept?token=${token}`)
      .pipe(map((response) => response.data))
  }

  constructor(private http: HttpClient) { }
}
