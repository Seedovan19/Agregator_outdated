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
    }
    else {
        alert("Неправильный email или пароль")
    }
    })() 

};