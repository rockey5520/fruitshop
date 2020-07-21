import { map } from 'rxjs/operators';
import { Component, OnInit } from '@angular/core';
import { FruitModel } from './../../models/fruit.model';
import { FruitService } from './../../services/fruit.service';
import { Observable } from 'rxjs';

@Component({
  selector: 'app-shop',
  templateUrl: './shop.component.html',
  styleUrls: ['./shop.component.scss']
})
export class ShopComponent implements OnInit {

  //fruitList: Observable<Array<FruitModel>>
  fruitList: FruitModel[] = []



  constructor(private fruitService: FruitService) {

  }

  ngOnInit(): void {
    this.getFruits()
  }

  getFruits(): void {
   this.fruitService.getFruits().subscribe(
      fruits => this.fruitList = fruits
    );
  }

}
