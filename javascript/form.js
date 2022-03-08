document.querySelector("#submit").onclick = function(){ 
    var object = {
        "adress": document.getElementById("adress").value,
        "name": document.getElementById("name").value,
        "square": parseInt(document.getElementById("square").value),
        "class": document.getElementById("class").value, //TODO: класс не видит
        "age_of_construction": document.getElementById("age_of_construction").value,
        "shelf_cost": parseInt(document.getElementById("shelf_cost").value), // TODO: не видит
        "floor_cost": parseInt(document.getElementById("floor_cost").value), // TODO: не видит
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