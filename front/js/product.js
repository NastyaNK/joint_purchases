
var host = "http://localhost:3333"
var productID = localStorage.getItem("open");

function httpGet(url){
    var req = new XMLHttpRequest();
    req.responseType = 'json';
    req.open('GET', url, true);
    req.onload  = function() {
       var jsonResponse = req.response;
       if (req.status != 200){
           return
       }

        addElement(jsonResponse.ID, jsonResponse.Name, jsonResponse.Description)
        
    };
    req.send(null);
}

httpGet(host+"/product/one/"+productID);


function addElement(id, name, description) {
    var newDiv = document.createElement("div");
    newDiv.innerHTML = '<div class="list_product">'+
    '<img src="img/arhidei/'+name+'.jpg" alt="" class="img_product">'+
    '<p class="name_product">'+name+'</p>'+
    '<p class="description_product">'+description+'</p>'+
    '<button  onclick="checkClick(this)" id="'+id+'">Добавить к заказу</button>'+
    '<p id="response"></p>'+
    '</div>';

    var my_div = document.getElementById("body_product");
    my_div.append(newDiv)
}


function checkClick(elem){
    let req = new XMLHttpRequest();
    req.open("POST", host+"/order/buy");

    req.setRequestHeader("Accept", "application/json");
    req.setRequestHeader("Content-Type", "application/json");

    req.onload  = function() {
        var fieldNameElement = document.getElementById('response');

        if (req.status != 201){
            fieldNameElement.innerHTML = "Попробуйте еще раз";
            return
        }
        fieldNameElement.innerHTML = "Ваш заказ прнят";
     };

    var userID = localStorage.getItem("userID");
    let data = `{
        "productID":`+elem.id+`,`+
        `"userID":`+userID+
    `}`;

    req.send(data);
}

