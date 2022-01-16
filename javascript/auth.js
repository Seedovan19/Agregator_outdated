document.querySelector("#submit").onclick = function(){ 
    var object = {
        "email": document.getElementById("email").value,
        "password": document.getElementById("password").value
    };

    sendRequest('POST', '/auth/sign-in', object)
    .then(data => localStorage.setItem('token', JSON.stringify(data.token)))
    .catch(err => console.log(err))
    


    // new Promise ( (resolve, reject) => {
    // const xhr = new XMLHttpRequest()
    // xhr.open('GET', '/')
    // xhr.setRequestHeader('Authorization', localStorage.token)
    // xhr.send()
    // xhr.onload  = () => {
    //     if(xhr.status >= 400) {
    //         reject(xhr.response)
    //     }
    //     else {
    //         resolve(xhr.response)
    //     }
    // }

    // xhr.onerror = () => {
    //     reject(xhr.response)
    // }
    // }).then(result => html(result))
    // .catch(err => console.log(err))
};


function sendRequest(method, url, params = null) {
    return new Promise ( (resolve, reject) => {
        const xhr = new XMLHttpRequest()
    xhr.open(method, url)
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