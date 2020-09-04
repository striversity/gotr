// https://developer.mozilla.org/en-US/docs/Web/API/Server-sent_events/Using_server-sent_events
// https://www.w3schools.com/html/html5_serversentevents.asp
function onLoaded() {
    var source = new EventSource("/sse/dashboard");
    source.onmessage = function (event) {
        console.log("OnMessage called:");
        console.dir(event);
        var dashboard = JSON.parse(event.data);
        var items = dashboard["inventory"]["items"];
        var booksQuantity = items["book"].quantity;
        var bicyclesQuantity = items["bicycle"].quantity;
        var rcCarsQuantity = items["rccar"].quantity;
            
        document.getElementById("biprice").innerHTML = items["bicycle"].price;
        document.getElementById("biquantity").innerHTML = bicyclesQuantity;
        document.getElementById("bprice").innerHTML = items["book"].price;
        document.getElementById("bquantity").innerHTML = booksQuantity;
        document.getElementById("rccprice").innerHTML = items["rccar"].price;
        document.getElementById("rccquantity").innerHTML = rcCarsQuantity;
        createLine([booksQuantity, bicyclesQuantity, rcCarsQuantity]);
    };

}
function createLine(data) {
    var ctx = document.getElementById('myChart').getContext('2d');
    var chart = new Chart(ctx, {
        // The type of chart we want to create
        type: 'line',

        // The data for our dataset
        data: {
            labels: ["Books", "Bicycles", "RC Cars",],
            datasets: [{
                label: "My First dataset",
                backgroundColor: 'rgb(255, 99, 132)',
                borderColor: 'rgb(255, 99, 132)',
                data: data,
            }]
        },

        // Configuration options go here
        options: {}
    });

}