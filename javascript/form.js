document.querySelector("#submit").onclick = function(){ 
    var object = {
        "adress": document.getElementById("adress").value,
        "name": document.getElementById("name").value,
        "square": parseInt(document.getElementById("square").value),
        "class": document.getElementById("class").value, 
        "shelf_cost": parseInt(document.getElementById("shelf_cost").value), 
        "floor_cost": parseInt(document.getElementById("floor_cost").value), 
        "description": document.getElementById("description").value
    };
    
    fetch("/api/warehouses/add_warehouse", {
        method: 'POST',
        headers: {
            Authorization: localStorage.token
        },
        body: JSON.stringify(object) // <-- Post parameters        
    })
};