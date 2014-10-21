'use strict';

var wishlistControllers = angular.module('wishlistControllers', []);

wishlistControllers.controller('SignInCtrl', ['$scope', '$location',  'User',
  function($scope, $location, User) {
    User.query(
              function(user){ /* success block*/
                $location.path("/products")
              });

    $scope.signIn = function(user) {
      User.signin({email: user.email}, function(value){ /* success block*/
                      console.log(value)
                      $location.path("/products")
                    });
    };
  }]);

wishlistControllers.controller('ProductListCtrl', ['$scope', '$location', 'User', 'Products', 'Bucket',
  function($scope,  $location, User, Products, Bucket) {
    $scope.user = User.query(
                    function(){ /* success block*/},
                    function(){ /* failure block*/
                      $location.path("/signin")
                    });

    $scope.products = Products.query();
    // derivative index for wish-list products
    $scope.selectedProducts =[];
    Bucket.query(function(values){
          angular.forEach(values, function(value, key) {
            this.push(value['product_id'])
          }, $scope.selectedProducts);
      });

    $scope.signOut = function() {
      User.signout( function(){ /* success block*/
                      $location.path("/signin")
                    });
    };

    $scope.addToBucket = function(product) {
      var bucketItem = new Bucket();
      bucketItem.product_id =  product.id;
      bucketItem.name = product.name;
      bucketItem.description = product.description;
      bucketItem.img = product.img;
      bucketItem.price = product.price;
      bucketItem.price_unit = product.price_unit;
      bucketItem.$save(function(bucketItem){
        // on successfully adding product to wish list modify index
        $scope.selectedProducts.push(bucketItem["product_id"]);
      }, 
      function(){
        // on failure do nothing
      });
    };

    $scope.removeFromBucket = function(product) {
      var bucketItem = new Bucket();
      bucketItem.$remove({'product_id': product.id}, 
        function(bucketItem){
          // modify index on successfully deletion product to wish list
          var indexPosition = $scope.selectedProducts.indexOf(bucketItem["product_id"]);
          console.log(bucketItem)
          $scope.selectedProducts.splice(indexPosition,1)
          console.log($scope.selectedProducts)
        }, 
        function(){
          // on failure do nothing
        });
    };
  }]);

wishlistControllers.controller('BucketListCtrl', ['$scope', '$location', 'User','Bucket',
  function($scope, $location, User, Bucket) {
    $scope.user = User.query(
                    function(){ /* success block*/},
                    function(){ /* failure block*/
                      $location.path("/signin")
                    });

    $scope.bucket = Bucket.query();

    $scope.signOut = function() {
      User.signout( function(){ /* failure block*/
                      $location.path("/signin")
                    });
    };

    $scope.removeFromBucket = function(bucketItem) {
      bucketItem.$remove({'product_id': bucketItem.product_id},
        function(value){
          // on success relode bucket list from server
          $scope.bucket = Bucket.query();
        }, 
        function(){
          // on failure do nothing
        });
    };
  }]);