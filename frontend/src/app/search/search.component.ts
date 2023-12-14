import { Component } from '@angular/core';
import { TypesenseService } from '../typesense.service';
import { SearchParams } from 'typesense/lib/Typesense/Documents';

@Component({
  selector: 'app-search',
  templateUrl: './search.component.html',
  styleUrls: ['./search.component.css']
})
export class SearchComponent {

  ngOnInit(): void {
   this.typesenseService.search("*").then((res) => console.log(res))
  }

  constructor(private typesenseService: TypesenseService) { }


}
