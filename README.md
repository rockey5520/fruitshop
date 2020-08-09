<p align="center">
    <a href=""">
        <img height=600 src="https://res.cloudinary.com/rockey5520/image/upload/v1596896057/fruitstore/fruit-market-1534355_m2c0av.jpg">
    </a>
    <br><u>A Fruit store built using Angular 10 for font end and Golang for backend api</u>
</p>




**Overview :** This project is to build an Fruit store with both front end and back end parts of the application. Angular 10 is used to build front end part and Golang is used for building RESTFul API serving data from in-memory DB.

**Observation:** When tried to attempt the same exact backend with Java, Go outperformed in performance and time taken for each api request. i have let the logging enabled while application running to have a little information about where the program control is going for each call and track the time take for each call to optimize further more

**Project Rules**

**Task**

Develop an online eCommerce store selling fruit, which contains the following features:

- Simple sign-up and login form.

- Browse the following products

- - Apples
  - Bananas
  - Pears
  - Oranges

- Add items to your cart

- - Adjust quantity.
  - Delete items from the cart.
  - Apply coupons.

- Checkout your cart

- - Mocked purchase
  - An address does not need to be entered.

**Requirements**

Architecture diagrams.

- Single-page frontend app.
- Backend RESTful web service written in GoLang.
- Users must be able to return to their cart after closing the browser, and see the previous items that were added.

**Cart Rules**

 If 7 or more apples are added to the cart, a 10% discount is applied to all apples.

- For each set of 4 pears and 2 bananas, a 30% discount is applied, to each set.

- - These sets must be added to their own cart item entry.
  - If pears or bananas already exist in the cart, this discount must be recalculated when new pears or bananas are added.

- A coupon code can be used to get a 30% discount on oranges, if applied to the cart, otherwise oranges are full price.

- - Can only be applied once
  - Has an configurable expiry timeout (10 seconds for testing purposes) once generated.

- The following totals must be shown:

- - Total price.
  - Total savings.

**Several ways to run the application**

- **Option 1** : **Deployed service**

  - I have deployed this application on [Google cloud run](https://cloud.google.com/run) and to access the service enter the following URL into browser(incognito mode) and register as a new user to  login to the website. Website URL : https://fruitshop-p5kl2n3awa-nn.a.run.app
  
  Container in Google cloud run takes only **69 milliseconds to start( which includes both frontend and backend )**
  - ![](https://res.cloudinary.com/rockey5520/image/upload/v1596896322/fruitstore/couldrun-fruitshop_ro5mtr.jpg)

- **Option 2 : Running prebuild docker image on your computer**

  - You can download the docker image I have prebuilt to your computer and run below command to the start the application and navigate to URL  http://localhost:8080 in your browser(incognito mode).

    - ```shell
      dockepull docker.pkg.github.com/rockey5520/fruitshop/fruitshop:5
      docker run -p 8080:8080 -it docker.pkg.github.com/rockey5520/fruitshop/fruitshop:5
      ```

- **Option 3 : Build and Run the image from the source code** (lengthier process)

  - ```shell
    git clone https://github.com/rockey5520/fruitshop.git
    cd fruitshop
    docker build -t fruitshop .
    docker run -p 8080:8080 -it fruitshop
    ```


**Running unit tests**

```shell
cd fruitshop/server
go test -v ./...
```



**Endpoints**

| Method | Path                                      | Description                                                  |
| ------ | ----------------------------------------- | ------------------------------------------------------------ |
| POST   | /server/customers                         | Saves new customer                                           |
| GET    | /server/customers/{loginid}               | Returns customer record                                      |
| GET    | /server/fruits                            | Fetched fruits available in the inventory                    |
| POST   | /server/cartitem                          | Adds new fruit entry to the cart                             |
| PUT    | /server/cartitem                          | Updates quantity of the fruit item in the cart if available |
| DELETE | /server/cartitem/{cart_id}/{fruitname}    | Deletes fruit entry in the cart                              |
| GET    | /server/cartitems/{cart_id}               | Fetched all items in a given cart                            |
| GET    | /server/cart/{cart_id}                    | Fectes cart details                                          |
| GET    | /server/discounts/{cart_id}               | Fetches all applied discounts a specific cart                |
| GET    | /server/orangecoupon/{cart_id}/{fruit_id} | Applies RANGE30 coupon to the cart and gets updated total and discounts |
| POST   | /server/pay                               | Takes the payment for the cart                               |



**Registration/Login Page**

To see the application, launch incognito mode in a browser you like and go to `http://localhost:8080`and this should present you a login form as below

![](https://res.cloudinary.com/rockey5520/image/upload/v1594849128/fruitstore/loginform_tzxsg1.jpg)

If you have not created an account earlier you can use the login id you choose (preferred you name without spaces or mobile number) and click login. But if you try to register with an id exists in database form with throw error saying `userid already exists`

**Shopping cart home page**

Upon login you should see a shopping cart something like below

![](https://res.cloudinary.com/rockey5520/image/upload/v1594849520/fruitstore/home_page_uugebm.jpg)



**Cart table and Discount table sample**

![](https://res.cloudinary.com/rockey5520/image/upload/v1594851723/fruitstore/discounts_applied_zvonn5.jpg)



Here you can use self explanatory descriptions to add fruits to carts and discounts earlier mentioned will be applied automatically but for 30% discount on oranges , one need to click on `ORANGE 30 Discount coupn`  to apply which is valid for only 10 seconds post the time discount will removed from the cart. 

In Angular Observable is used to link the components so that changes are applied across other components when there is a change to one. Discount coupon table is not in requirement but I left it there so it sits as a nice help to check if the discounts and coupons active for a particular user.



**Tech Stack**

- Angular 10
- Golang 1.14.6
- SqlLite ( code is written for Postgres, Mysql as well to switch to a different database)
- GORM ( ORM library)
- Gorilla MUX (HTTP router)
- Docker

**Interesting Parts of the application**

- **Docker builds**
  
  - I have used docker multistage builds for this implementation which helped docker image size a so small that everything just fitted in **26MB** which is amazing when you want you application to be rapidly sent across for deployments
- **Angular + Go**
  
  - Instead of hosting frontend as a standalone application along side Go server and establishing connectivity between those, I found many advantages by mounting the production distribution folder of angular application and used Golang HTTP Fileserver to mount and redirected calls to index.html () 
- **Blazing fast**
  - Application is blazing fast due the fact that 
    - There is no extra network call to frontend application
    
    - Go application is built and loaded to image as a binary (binary images are slightly faster than go run)
    
    - Smaller Docker images(26mb) from faster pull and **speedy application start (around 250 milli seconds**)
    
    - ![](https://res.cloudinary.com/rockey5520/image/upload/v1596920226/fruitstore/Application_start_time_pwicrl.jpg)
    
  
- **open for extension closed for modification** 
  
  - I have ensured the application is open for extending new functionalities but closed for modification, This is particularly resembles in the way discount rules are applied. Discount Rules are stored in database and programmed the logic to apply discounts by fetching conditions from DB. This gives a flexibility to add/update/delete discount rules without rebuilding the application. Every time there is a change to one of discount rule we will only update in the database as meta information and backend application will apply with the existing logic built for the new discounts as well as existing discount without modifying the core logic.



**Entity Relationship diagram**

![](https://res.cloudinary.com/rockey5520/image/upload/v1597008008/fruitstore/Entity_relation_ship_diagram_dorxfr.jpg)



**Sequence diagrams**

- **Registration Journey**

  ![](https://res.cloudinary.com/rockey5520/image/upload/v1597003673/fruitstore/Registration-Journey-SequenceDiagram_ascmlw.png)

- **Login Journey**

  ![](https://res.cloudinary.com/rockey5520/image/upload/v1597003673/fruitstore/Login-Journey-SequenceDiagram_rneb4n.png)



- **Add Fruit Journey**

  ![](https://res.cloudinary.com/rockey5520/image/upload/v1597003673/fruitstore/Add_Fruit-Journey-SequenceDiagram_tmsoge.png)

- **Update Fruit**

  ![](https://res.cloudinary.com/rockey5520/image/upload/v1597003673/fruitstore/Update_Fruit_Quantity-Journey-SequenceDiagram_yczg05.png)

- **Delete Fruit**

  ![](https://res.cloudinary.com/rockey5520/image/upload/v1597003673/fruitstore/Delete_Fruit_sequence_diagram_tv5yz5.png)



Wondering why branch named `main` instead of `master` ? Reason is after the following tweet from github ceo https://twitter.com/natfriedman/status/1271253144442253312 i am trying to use less racially charged words. Not that this will resolve every problem related to racial injustice,  However i believe it's always good to have the know-how and the tools to be able to change the names we use. That goes for branch names as well. :)
