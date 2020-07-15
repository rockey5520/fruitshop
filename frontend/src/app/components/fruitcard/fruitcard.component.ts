import { Component, OnInit, Input } from '@angular/core';
import { FruitModel } from 'src/app/models/fruit.model';
import { CartService } from './../../services/cart.service';
import { AuthenticationService } from './../../services/authentication.service';
import { DiscountService } from './../../services/discount.service';
import { User } from '../../models/user.model';

@Component({
  selector: 'app-fruitcard',
  templateUrl: './fruitcard.component.html',
  styleUrls: ['./fruitcard.component.scss']
})
export class FruitcardComponent implements OnInit {

  @Input()
  fruit: FruitModel
  currentUser: User;

  count: number;

  constructor(private cartService: CartService, public authenticationService: AuthenticationService, public discountService: DiscountService) {
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
    this.cartService.addToCart(this.currentUser.userId, this.fruit, this.count).subscribe(() => {
      this.cartService.update.next(true)
      this.discountService.update.next(true)
    })
  }
}
