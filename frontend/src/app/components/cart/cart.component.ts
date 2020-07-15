import { FormGroup } from '@angular/forms';
import { Component, OnInit } from '@angular/core';

import { CartModel } from './../../models/cart.model';
import { CartService } from './../../services/cart.service';
import { PaymentService } from './../../services/payment.service';
import { DiscountService } from './../../services/discount.service';
import { AuthenticationService } from './../../services/authentication.service';
import { BehaviorSubject } from 'rxjs';
import { Observable, pipe } from 'rxjs';
import { map, reduce, filter, first } from 'rxjs/operators';
import { async } from 'rxjs/internal/scheduler/async';
import { User } from '../../models/user.model';
import { DiscountModel } from '../../models/discount..model';

import { Inject } from '@angular/core';
import { ViewChild, ElementRef, AfterViewInit } from "@angular/core";
import { DOCUMENT } from '@angular/common';
import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';


@Component({
  selector: 'app-cart',
  templateUrl: './cart.component.html',
  styleUrls: ['./cart.component.scss']
})
export class CartComponent implements OnInit {

  cartList: Observable<Array<CartModel>>;
  discountList: Observable<Array<DiscountModel>>

  displayedColumns: string[] = ['name', 'costPerItem', 'count', 'totalCost'];
  cartdisplayedColumns: string[] = ['name', 'status'];

  total: number;
  currentUser: User;
  cartId: string
  totalSavings: number
  seconds1: number
  formNotComplete: boolean = false;



  constructor(public cartService: CartService, public paymentService: PaymentService, public authenticationService: AuthenticationService,
    public discountService: DiscountService, @Inject(DOCUMENT) document) {
    this.authenticationService.currentUser.subscribe(x => this.currentUser = x);
    document.getElementById('countdown');

  }


  ngOnInit(): void {
    this.updateData()
    this.updateDiscountData();
    this.cartService.update.subscribe((data: boolean) => {
      if (data) {
        this.updateData();
      }
    })

    this.discountService.update.subscribe((data: boolean) => {
      if (data) {
        this.updateDiscountData();
      }
    })


  }

  updateData() {
    this.cartList = this.cartService.getCartByID(this.currentUser.userId)
      .pipe(map(item => item.filter(item => item.count > 0)));

    this.total = 0;
    this.cartList.subscribe((data) => {
      this.cartId == data.map(item => item.cartId)[0];
      this.total = data.map(item => item.totalCost).reduce((a, b) => a + b, 0);
      this.totalSavings = data.map((item) => item.count * item.costPerItem).reduce((a, b) => a + b, 0) - this.total

    })
  }

  updateDiscountData() {

    this.discountList = this.discountService.getDiscountsByID(this.currentUser.userId)
    console.log("discount list", this.discountList)
  }



  pay(): void {
    this.paymentService.pay(this.currentUser.userId, this.currentUser.userId, this.total).subscribe(() => {
      this.cartService.update.next(true)

    })
  }

  applyDiscount(): void {

    this.paymentService.applyDiscount(this.currentUser.userId).subscribe(() => {
      this.cartService.update.next(true)
      this.discountService.update.next(true)
    })
    this.formNotComplete = true;
    setTimeout(() => {
      this.cartService.getCartByID(this.currentUser.userId,).subscribe(() => {
        this.cartService.update.next(true)
        this.discountService.update.next(true)
      })
      this.formNotComplete = false;
    },
      13000);

  }



}



