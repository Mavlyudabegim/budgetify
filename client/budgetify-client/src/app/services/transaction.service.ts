import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable, tap } from 'rxjs';
import { TransactionModel } from '../shared/models/transaction.model';

@Injectable({
  providedIn: 'root',
})
export class TransactionService {
  public transactions: TransactionModel[] = [];
  constructor(private httpClient: HttpClient) {}
  public getTransactions(accountId: string): Observable<any> {
    return this.httpClient
      .get(
        `http://localhost:3000/api/transactions/${accountId}/account-transactions`
      )
      .pipe(
        tap({
          next: (res: any) => {
            this.transactions = res;
          },
        })
      );
  }
}
