'use strict';

var wishlistServices = angular.module('wishlistServices', ['ngResource']);

wishlistServices.factory('User', ['$resource',
  function($resource){
    return $resource('api/user', {}, {
      query: {method:'GET', isArray:false},
      signin: {method:'POST', isArray:false},
      signout: {method: 'DELETE', isArray: false}
    });
}]);

wishlistServices.factory('Products', ['$resource',
  function($resource){
    return $resource('api/products', {}, {
      query: {method:'GET', isArray:true}
    });
}]);

wishlistServices.factory('Bucket', ['$resource',
  function($resource){
    return $resource('api/bucket', {}, {
      query: {method:'GET', isArray:true},
      save: {method: 'POST', isArray: false},
      remove: {method: 'DELETE', isArray: false}
    });
}]);