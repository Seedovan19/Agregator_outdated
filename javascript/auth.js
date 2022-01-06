document.querySelector("#submit").onclick = function(){ 
    var object = {
        "email": document.getElementById("email").value,
        "password": document.getElementById("password").value
    };
    
    (async () => {
    let response = await fetch("/auth/sign-in", {
        method: 'POST',
        body: JSON.stringify(object) // <-- Post parameters        
    })
    let result = await response.json();
    if (result.token != null) {
        sessionStorage.setItem('token', result.token);
        // const request = new XMLHttpRequest();
        // request.open('GET', "/");
        // request.setRequestHeader('Authorization', sessionStorage.token);
        // request.send();
    }
    else {
        alert("Неправильный email или пароль")
    }
    })()

};