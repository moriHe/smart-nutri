import { Injectable } from '@angular/core';
import Typesense from 'typesense';
import { SearchParams, SearchParamsWithPreset, SearchResponse, SearchResponseHit } from 'typesense/lib/Typesense/Documents';

export type Result = {
  ingredientId: number,
  name: string,
  category: string
}

@Injectable({
  providedIn: 'root'
})

// If not working, check if in_memory collection has been deleted due to laptop restart
export class TypesenseService {
  client = new Typesense.Client({
    'nodes': [{
      'host': 'localhost',
      'port': 8108,
      'protocol': 'http'
    }],
    'apiKey': 'xyz',
    'connectionTimeoutSeconds': 2
  })

  async search(q: string) {
    const searchParams: SearchParams = {
      q,
      'per_page': 20,
      query_by: "name",
    }
    
    const data = await this.client.collections('ingredients')
      .documents()
      .search(searchParams)
    
    return data.hits as SearchResponseHit<Result>[]
  }

  constructor() { }
}