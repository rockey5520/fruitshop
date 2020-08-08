export interface FruitModel {
    id: number;
    name: string;
    price: number;
}

export interface RootObject {
    data: FruitModel[];
} 