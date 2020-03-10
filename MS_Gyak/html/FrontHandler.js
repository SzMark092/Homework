
//Send a number to the server, and a function code.
function getTableData(dataType) {


    var xhttp = new XMLHttpRequest();

    xhttp.responseType = 'json';
    xhttp.open("GET", "localhost:8080/getTable?typeCode=" + dataType, false);

    xhttp.send(null)

    var jsonInput = JSON.parse(xhttp.response);
    var resultelement = document.getElementById('result').appendChild(document.createTextNode(jsonInput))

}

function createTable(dataType) {

    xhttp.open("POST", "localhost:8080/create?typeCode=" + dataType, false);
    xhttp.send();

}

function callChoosenAction() {

    var actionType = document.getElementById('actionSelector').value
    var dataType = document.getElementById('dataSelector').value

    if (actionType == 1) {
        createTable(dataType)
    } else {
        getTableData(dataType)
    }

}