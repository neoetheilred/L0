import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root'
})
export class OrderService {

  constructor(private http: HttpClient) { }

  getOrder(id: string) {
    return this.http.get<any>(`http://localhost:8000/order/${id}`)
  }

  getAvailableIds() {
    return this.http.get<string[]>(`http://localhost:8000/available`)
  }
}
