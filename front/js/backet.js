var host = "http://localhost:3333"
var userID = localStorage.getItem("userID");

async function getBacket(container, on_load = null) {
    send("/basket/" + userID, null, async(status, res) => {
        if (status != 200) {
            return
        }
        if (res == null) {
            container.innerHTML = await template("search-empty", { 'text': "Ничего не найдено" });
            return
        }
        let content = "";
        for (let i = 0; i < res.length; i++) {
            await send("/product/one/" + res[i].ProductID, null, async(status, result) => {
                if (status != 200) {
                    return
                }
                if (result == null) {
                    container.innerHTML = await template("search-empty", { 'text': "Ничего не найдено" });
                    return
                }
                result['StartTime'] = result['StartTime'].replace("T", " ").replace("Z", "")
                result['EndTime'] = result['EndTime'].replace("T", " ").replace("Z", "")
                content += await template("basket-product", result);
            });
        }
        container.innerHTML = content;
        if (on_load != null) on_load(res, container);
    });
}

add_on_loads(() => {

    if (userID != null)
        getBacket(document.getElementById('basket-products'));
})

function addElement(id, name, count) {
    var newDiv = document.createElement("div");
    newDiv.innerHTML = '<div class="list_product">' +
        '<img src="img/products/' + name + '.jpg" alt="" class="cart-img">' +
        '<p class="name_product">' + name + '</p>' +
        '<p id="' + id + '" class="product_count_id">' + count + '</p>' +
        '</div>';

    var my_div = document.getElementById("backet");
    my_div.append(newDiv)
}


function addBacket(elem, id, count = 1) {
    if (userID != null)
        send("/basket/add", { 'productID': id, 'userID': userID - 0, 'count': count }, (status, result) => {
            getBacket(document.getElementById('basket-products'));
        });
}

function buyBacket() {
    backet = document.getElementsByClassName("product_count_id")

    data = "["
    for (let i = 0; i < backet.length; i++) {
        data += `{"productID":` + backet[i].id + `,"userID":` + userID + `,"count":` + backet[i].outerText
        if (i != backet.length - 1) {
            data += "},"
        } else {
            data += "}"
        }
    }
    data += "]"

    let req = new XMLHttpRequest();
    req.open("POST", host + "/order/buy");

    req.setRequestHeader("Accept", "application/json");
    req.setRequestHeader("Content-Type", "application/json");

    req.onload = function() {
        if (req.status == 201) {
            drop = document.getElementsByClassName("list_product")
            len = drop.length
            for (let i = 0; i < len; i++) {
                drop[0].remove();
            }
        }
    };

    req.send(data);
}