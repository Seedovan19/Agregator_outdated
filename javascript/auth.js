document.querySelector("#submit").onclick = function(){ 
    var object = {
        "email": document.getElementById("email").value,
        "password": document.getElementById("password").value
    };

    sendRequest('POST', '/auth/sign-in', object, true)
    .then(data => {localStorage.token = data.token ;console.log(data.token)})
    .catch(err => console.log(err))
};


function sendRequest(method, url, params = null, asyncProc) {
    return new Promise ( (resolve, reject) => {
        const xhr = new XMLHttpRequest()
    xhr.open(method, url, !(!asyncProc))
    xhr.responseType = 'json'
    xhr.setRequestHeader('Content-Type', 'application/json')
    
    xhr.onload  = () => {
        if(xhr.status >= 400) {
            reject(xhr.response)
        } else {
            resolve(xhr.response)
        }
    }

    xhr.onerror = () => {
        reject(xhr.response)
    }

    xhr.send(JSON.stringify(params))
    })
}