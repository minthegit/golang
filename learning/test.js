var fn = function() {
  this.run = function() {
    return {"run"};
  }
}


var obj = new fn();
console.log(obj.run());
var str = "asdasdasd";
var k = obj.run();
console.log("result : " + k);
obj.run("asd");
