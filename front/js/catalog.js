async function loadProducts(container, on_load = null, search = "") {
    if (search != "")
        search = "/" + search
    send("/product/list" + search, null, async(status, result) => {
        if (status != 200) {
            return
        }
        if (result == null) {
            container.innerHTML = await template("search-empty", { 'text': "Ничего не найдено" });
            return
        }
        for (let i = 0; i < result.length; i++) {
            result[i]['StartTime'] = result[i]['StartTime'].replace("T", " ").replace("Z", "")
            result[i]['EndTime'] = result[i]['EndTime'].replace("T", " ").replace("Z", "")
            result[i]['ImageOne'] = result[i]['Image'].split(";")[0]
            result[i]['percent'] = Math.floor(100 * Math.random())
        }
        let content = await multi_template("product", result);
        container.innerHTML = content;
        if (on_load != null) on_load(result, container);
    });
}

async function loadProduct(container, productId, on_load = null) {
    send("/product/one/" + productId, null, async(status, result) => {
        if (status != 200) {
            return
        }
        if (result == null) {
            container.innerHTML = await template("search-empty", { 'text': "Ничего не найдено" });
            return
        }
        result['StartTime'] = result['StartTime'].replace("T", " ").replace("Z", "")
        result['EndTime'] = result['EndTime'].replace("T", " ").replace("Z", "")
        let content = await template("full-product", result);
        container.innerHTML = content;
        if (on_load != null) on_load(result, container);
    });
}

function settingProductItem(item) {
    let tobasket = item.querySelector("[tobasket]");
    if (tobasket != null) {
        tobasket.onclick = (event) => {
            event.stopPropagation();
            send("/basket/add", { 'ProductID': item.getAttribute("product-id") - 0, "UserID": localStorage.getItem("userID") - 0, 'Count': 4 })
        }
    }
}

function listenClick(element) {
    localStorage.setItem("open", element.getAttribute("product-id"));
    window.open("/mvp/product.html", '_self')
}

function loadAll(result, container) {
    Object.values(container.children).forEach(element => {
        element.onclick = (event) => {
            event.stopPropagation();
            loadProduct(fullWrapper, element.getAttribute("product-id"), (result2, container2) => {
                let prod = container2.children[0];
                prod.onclick = (event) => {
                    event.stopPropagation();
                }
                let buttons = prod.querySelector("#buttons");
                let contents = prod.querySelector("#contents");
                let items = prod.querySelector("#items");
                let image = prod.querySelector("#big-image");
                add_switcher(buttons, contents);

                let xSlider = new ImageSlider(items, image, "slider-image");
                let imgs = result2.Image.split(";");
                imgs.forEach((img) => xSlider.add_image("/mvp/img/products/" + img));
                settingProductItem(container2.children[0]);

                catalog.classList.toggle("blurring", true);
                fullWrapper.hidden = false;
            });
        }
    });
}

let search, catalog, fullWrapper;
add_on_loads(() => {
    search = document.querySelector(".search"); // Поисковая строка
    catalog = document.getElementById("catalog"); // Контейнер для всех товаров
    fullWrapper = document.getElementById("full-wrapper"); // Для открытия одного
    if (catalog != null) {
        catalog.classList.toggle("blurring", false);
        fullWrapper.hidden = true;
        fullWrapper.onclick = () => {
            fullWrapper.hidden = true;
            catalog.classList.toggle("blurring", false);
        }
        search.oninput = () => {
            loadProducts(catalog, loadAll, search.value);
        }
        loadProducts(catalog, loadAll);
    }
});