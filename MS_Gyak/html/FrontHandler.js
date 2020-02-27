
//Send a number to the server, and a function code.
function getTableData(dataType) {


    var xhttp = new XMLHttpRequest();

    xhttp.responseType = 'json';
    xhttp.open("GET", "localhost:8080/getTable?tableType=" + dataType, false);
    xhttp.send();

    var jsonInput = JSON.parse(xhttp.response);

    var resultelement = document.getElementById('result').appendChild(document.createTextNode(jsonInput))

}

function createTable(dataType) {

    xhttp.responseType = 'json';
    xhttp.open("GET", "localhost:8080/create?tableType=" + dataType, false);
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