import { OrderedFruitModel, OrderedFruitModelData } from './../../models/orderedfruit.mode';
import { FruitModel } from './../../models/fruit.model';
import { DiscountModel} from './../../models/discount..model';

;
import { Customer } from './../../models/customer.model';


import { Component, OnInit } from '@angular/core';


import { CartService } from './../../services/cart.service';
import { PaymentService } from './../../services/payment.service';
import { DiscountService } from './../../services/discount.service';
import { AuthenticationService } from './../../services/authentication.service';

import { Observable, pipe } from 'rxjs';
import { map, reduce, filter, first } from 'rxjs/operators';
import { async } from 'rxjs/internal/scheduler/async';


import { Inject } from '@angular/core';
import { ViewChild, ElementRef, AfterViewInit } from "@angular/core";
import { DOCUMENT } from '@angular/common';
import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { CartItem } from 'src/app/models/cartitem.model';
import { FruitcardComponent } from '../fruitcard/fruitcard.component';


@Component({
  selector: 'app-cart',
  templateUrl: './cart.component.html',
  styleUrls: ['./cart.component.scss']
})
export class CartComponent implements OnInit {

  cartList: Observable<Array<CartItem>>;
  discountList: Observable<DiscountModel[]>;
  displayedColumns: string[] = ['name', 'costPerItem', 'count', 'totalCost'];
  cartdisplayedColumns: string[] = ['name', 'status'];
  total: number;
  currentUser: Customer;
  cartId: string
  totalSavings: number
  seconds1: number
  formNotComplete: boolean = false;
  myVar : Number
  DiscountCouponFruit: string


  constructor(public cartService: CartService,
    public paymentService: PaymentService,
    public authenticationService: AuthenticationService,
    public discountService: DiscountService,
    @Inject(DOCUMENT) document,
    private httpclient: HttpClient) 
    {
    this.authenticationService.currentUser.subscribe(x => this.currentUser = x);
    document.getElementById('countdown');
  }


  ngOnInit(): void {
    this.updateCartList()
    this.updateDiscountData();
    this.cartService.update.subscribe((data: boolean) => {
      if (data) {
        this.updateCartList();
        this.updateDiscountData();
      }
    })

    this.discountService.update.subscribe((data: boolean) => {
      if (data) {
        this.updateDiscountData();
      }
    })
    this.authenticationService.update.subscribe((data: boolean) => {
      if (data) {
        this.updateUserData();
      }
    })

  }

  updateCartList() {
  
    this.cartList = this.cartService.getCartByID(this.currentUser.Cart.ID)
      .pipe(map(item => item.filter(item => item.count > 0)));

    this.total = 0;
    this.cartList.subscribe((data) => {
      this.total = data.map(item => item.itemtotal).reduce((a, b) => a + b, 0);
      this.totalSavings = data.
      map((item) => item.count * item.costperitem).reduce((a, b) => a + b, 0) - this.total

    })
  }
  updateDiscountData() {
    this.discountList = this.discountService.getDiscountsByID(this.currentUser.Cart.ID)
      .pipe(map(item => item.filter(item => item.status == "APPLIED")));
    }
  
  updateUserData() {  
    this.authenticationService.login(this.currentUser.loginid).subscribe((x)=>{
      this.cartService.update.next(true)
    })
  }

  pay(): void {
    
    this.paymentService.pay(this.currentUser.ID, this.currentUser.Cart.ID, this.total).subscribe(() => {
      this.authenticationService.update.next(true)
      
    },error => {
      console.error(error)
    })
    
  }



  applyDiscount(): void {
  
    this.paymentService.applyDiscount(this.currentUser.Cart.ID, 4).subscribe(() => {
      this.cartService.update.next(true)
      this.discountService.update.next(true)

    })
    this.formNotComplete = true;
    
    setTimeout(() => {
      this.cartService.getCartByID(this.currentUser.Cart.ID,).subscribe(() => {
        this.cartService.update.next(true)
        this.discountService.update.next(true)
      })
      this.formNotComplete = false;
    },
      13000);

  }

  

  



}



