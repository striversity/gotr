// https://www.w3schools.com/js/js_ajax_intro.asp
// https://www.w3schools.com/js/js_ajax_http_send.asp

function addUser(e) {
    e.preventDefault();

    var xhttp = new XMLHttpRequest();
    
    xhttp.onreadystatechange = function () {
        if (this.readyState == 4 && this.status == 200) {
            getUsers();
        }
    };
    
    console.dir(document.getElementById("username").value);
    var username = document.getElementById("username").value;
    var email = document.getElementById("email").value;
    var password = document.getElementById("password").value;

    var data = { username: username, email: email, password: password };

    xhttp.open("POST", "/api/users/", true);
    xhttp.send(JSON.stringify(data));
}

function getUsers() {
    var xhttp = new XMLHttpRequest();

    xhttp.onreadystatechange = function () {
        if (this.readyState == 4 && this.status == 200) {
            var users = JSON.parse(this.responseText);
            console.dir(users);

            userCountElement = document.getElementById("user-count");
            userCountElement.innerHTML = users.length
        }
    };

    xhttp.open("GET", "/api/users/", true);
    xhttp.send();
}