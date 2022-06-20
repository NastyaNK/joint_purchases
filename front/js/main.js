// Дополнение строк: форматирование
String.prototype.format = function() {
    var args = arguments;
    return this.replace(/{(\d+)}/g, function(match, number) {
        return typeof args[number] != 'undefined' ? args[number] : match;
    });
};
// Дополнение строк: шаблонизатор
String.prototype.template = function(params) {
    const names = Object.keys(params);
    const vals = Object.values(params);
    return new Function(...names, `return \`${this}\`;`)(...vals);
};
// Дополнение строк: замена всех совпадений в строке
String.prototype.replaceAll = function(search, replacement) {
    var target = this;
    return target.replace(new RegExp(search, 'g'), replacement);
};

// Дополнение массивов: удаление всех совпадений по значению
Array.prototype.removeByValue = function(item) {
    for (let i = this.length; i--;)
        if (this[i] === item) this.splice(i, 1);
    return this;
};
// Дополнение массивов: удаление всех совпадений по индексу
Array.prototype.remove = function(index) {
    if (index >= 0 && index < this.length)
        this.splice(index, 1);
    return this;
};

// Выполнение нескольких функций после загрузки окна
let on_loads = [];
let on_loaded = false;
window.onload = (e) => {
    on_loaded = true;
    on_loads.forEach((func) => func())
};

add_on_loads = (obj) => {
    if (on_loaded)
        obj();
    else
        on_loads.push(obj);
};


add_switcher = (buttons, contents) => {
    let btns = buttons.children;
    let ctns = contents.children;
    for (let i = 0; i < btns.length; i++) {
        btns[i].onclick = () => {
            for (let j = 0; j < ctns.length; j++) {
                if (i == j) {
                    btns[j].classList.add("active");
                } else {
                    btns[j].classList.remove("active");
                }
                ctns[j].hidden = i != j;
            }
        }
    }
    btns[0].onclick();
}

// function clear() {
//     localStorage.removeItem("open");
// }

// clear();