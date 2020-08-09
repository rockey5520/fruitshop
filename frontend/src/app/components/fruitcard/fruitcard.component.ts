import { CartComponent } from './../cart/cart.component';
import { Customer } from './../../models/customer.model';
import { Component, OnInit, Input } from '@angular/core';
import { FruitModel } from 'src/app/models/fruit.model';
import { CartService } from './../../services/cart.service';
import { AuthenticationService } from './../../services/authentication.service';
import { DiscountService } from './../../services/discount.service';
import { FruitService } from 'src/app/services/fruit.service';
import { Observable, pipe } from 'rxjs';

import { map } from 'rxjs/operators';
import { CartItem } from 'src/app/models/cartitem.model';


@Component({
  selector: 'app-fruitcard',
  templateUrl: './fruitcard.component.html',
  styleUrls: ['./fruitcard.component.scss']
})
export class FruitcardComponent implements OnInit {

  @Input()
  fruit: FruitModel
  currentUser: Customer;
  cartList: Observable<CartItem[]>;

  count: number;

  constructor(private cartService: CartService, 
    public authenticationService: AuthenticationService, 
    public discountService: DiscountService,
    public fruitService: FruitService,
    ) {
    this.authenticationService.currentUser.subscribe(x => {
      this.currentUser = x
    });

  }

  

  ngOnInit(): void {
    this.count = 0;
  }

  addToCount(): void {
    this.count++;
  }

  subToCount(): void {
    this.count--;
  }

  addToCart(): void {
    console.log(this.count)
    if(this.count == 0){
      this.count++;
      this.cartService.addToCart(this.currentUser.Cart.ID, this.fruit, this.count).subscribe(() => {
        this.cartService.update.next(true)
        this.discountService.update.next(true)
      })
    }else {
      this.count++;
      this.cartService.subtractFromCart(this.currentUser.Cart.ID, this.fruit, this.count).subscribe(() => {
        this.cartService.update.next(true)
        this.discountService.update.next(true)
      })
    }
    console.log(this.count)
  }

  reduceFromCart(): void{
    this.count--;
    if(this.count == 0 ){
      this.cartService.deleteFromCart(this.currentUser.Cart.ID, this.fruit, this.count).subscribe(() => {
        this.cartService.update.next(true)
        this.discountService.update.next(true)
      })
    }else{
      this.cartService.subtractFromCart(this.currentUser.Cart.ID, this.fruit, this.count).subscribe(() => {
        this.cartService.update.next(true)
        this.discountService.update.next(true)
      })
    }

  }

  deleteFromCart(): void{
    this.count = 0;
    this.cartService.deleteFromCart(this.currentUser.Cart.ID, this.fruit, this.count).subscribe(() => {
      this.cartService.update.next(true)
      this.discountService.update.next(true)
    })
  }

}
