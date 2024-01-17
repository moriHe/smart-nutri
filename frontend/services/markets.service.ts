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

  ShoppingListDisplay = {
    ...this.MarketDisplay,
    "TODAY": "Heute"
  }

  MarketKey = {
    [Markets.NONE]: Markets.NONE,
    [Markets.REWE]: Markets.REWE,
    [Markets.EDEKA]: Markets.EDEKA,
    [Markets.BIO_COMPANY]: Markets.BIO_COMPANY,
    [Markets.WEEKLY_MARKET]: Markets.WEEKLY_MARKET,
    [Markets.ALDI]: Markets.ALDI,
    [Markets.LIDL]: Markets.LIDL

  }

  constructor() { }
}
