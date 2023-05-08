import { Component, OnInit } from '@angular/core';
import { OrderService } from '../order.service';
import { FormControl } from '@angular/forms';

@Component({
  selector: 'app-order-display',
  templateUrl: './order-display.component.html',
  styleUrls: ['./order-display.component.css']
})
export class OrderDisplayComponent implements OnInit {
  constructor(private orderService: OrderService) { }

  order = undefined
  orderId = new FormControl("")
  availableIds: string[] = []

  ngOnInit(): void {
    console.log("init")
    this.orderService.getOrder("aaaaa").subscribe(data => {
      this.order = data
    })
    this.orderService.getAvailableIds().subscribe(data => {
      this.availableIds = data
    })
  }

  getOrderById(id: string) {
    this.orderService.getOrder(id).subscribe(data => {
      this.order = data
    })
  }

  getOrder() {
    this.orderService.getOrder(this.orderId.value!).subscribe(data => {
      this.order = data
    })
  }

  stringify() {
    console.log(JSON.stringify(this.order, undefined, 2))
    return JSON.stringify(this.order, undefined, 2)
  }
}
