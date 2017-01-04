angular.module('myApp', []).controller('rankController', ['$scope', '$http', function($scope, $http) {

    $scope.init = init;

    function init() {
        $http.get('http://52.78.203.214/wall')
            .success(function (data) {
                $scope.ranks = data;
            })
            .error(function (errors, status) {
                alert('获取排名数据失败!');
            });
    }

}]);