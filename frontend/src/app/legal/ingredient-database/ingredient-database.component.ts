import { HttpClient } from '@angular/common/http';
import { Component } from '@angular/core';
import { map } from 'rxjs';

@Component({
  selector: 'app-ingredient-database',
  templateUrl: './ingredient-database.component.html',
  styleUrls: ['./ingredient-database.component.css']
})
export class IngredientDatabaseComponent {
  downloadTable() {
    const downloadUrl = "http://localhost:8080/datenbank-nahrungsmittel";

    this.http.get<Blob>(downloadUrl, { responseType: 'blob' as 'json' }).subscribe(blob => {
      const url = window.URL.createObjectURL(blob);
      const a = document.createElement('a');
      a.href = url;
      a.download = 'ingredients.csv';
      document.body.appendChild(a);
      a.click();
      document.body.removeChild(a);
      window.URL.revokeObjectURL(url);
    });
  }

  constructor(private http: HttpClient) { }
}
