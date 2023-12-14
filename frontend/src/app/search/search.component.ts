import { Component } from '@angular/core';
import { Result, TypesenseService } from '../typesense.service';
import { SearchResponseHit } from 'typesense/lib/Typesense/Documents';





@Component({
  selector: 'app-search',
  templateUrl: './search.component.html',
  styleUrls: ['./search.component.css']
})
export class SearchComponent {
  results: SearchResponseHit<Result>[] = []
  ngOnInit(): void {
   this.typesenseService.search("*").then((res) => {
    if (res) {
      console.log(res)
      this.results = res
    }
   })
  }

  constructor(private typesenseService: TypesenseService) { }


}
