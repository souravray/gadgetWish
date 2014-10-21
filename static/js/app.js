'use strict';

var wishlistApp = angular.module('wishlistApp', [
  'ngRoute',
  'wishlistControllers',
  //'wishlistFilters',
  'wishlistServices'
]);

wishlistApp.config(['$routeProvider',
  function($routeProvider) {
    $routeProvider.
      when('/signin', {
        templateUrl: 'partials/signin.html',
        controller: 'SignInCtrl'
      }).
      when('/products', {
        templateUrl: 'partials/product-list.html',
        controller: 'ProductListCtrl'
      }).
      when('/bucket', {
        templateUrl: 'partials/bucket-list.html',
        controller: 'BucketListCtrl'
      }).
      otherwise({
        redirectTo: '/signin'
      });
  }]);