function register() {
    var name = document.getElementById("form_register_name").value;
    var email = document.getElementById("form_register_email").value;
    var pwd = document.getElementById("form_register_password").value;
    var role = document.getElementById("form_register_role").value;
    send("/auth/register", { 'email': email, 'name': name, 'password': pwd, 'role': role }, (status, result) => {
        var fieldNameElement = document.getElementById('result_register');
        if (status != 201) {
            fieldNameElement.innerHTML = "Попробуйте еще раз";
            return
        }
        if (status == 201) {

            document.getElementById("form_login_email").value = email;
            document.getElementById("form_login_password").value = pwd;
            login();
            return
        }
    });
}

function login() {
    var email = document.getElementById("form_login_email").value;
    var pwd = document.getElementById("form_login_password").value;
    send("/auth/login", { 'email': email, 'password': pwd }, (status, result) => {
        var fieldNameElement = document.getElementById('login_error');
        if (status != 200) {
            fieldNameElement.innerHTML = "Попробуйте еще раз";
            return
        }
        localStorage.setItem("userID", result);
        window.location = '/mvp/profile.html';
    });
}