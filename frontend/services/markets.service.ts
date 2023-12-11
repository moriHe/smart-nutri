import { Injectable } from '@angular/core';
import { Markets } from 'api/recipes/recipes.interface';

@Injectable({
  providedIn: 'root'
})
export class MarketsService {
  
  MarketDisplay = {
    [Markets.NONE]: "-",
    [Markets.REWE]: "Rewe",
    [Markets.EDEKA]: "Edeka",
    [Markets.BIO_COMPANY]: "Bio Company",
    [Markets.WEEKLY_MARKET]: "Wochenmarkt",
    [Markets.ALDI]: "Aldi",
    [Markets.LIDL]: "Lidl"
  }

  constructor() { }
}