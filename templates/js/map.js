// create a new instance platform
var platform = new H.service.Platform({
    apikey: "UxDC5OOIhOC4BEmxOmwUKAa8q413CJS311fmOzt3UeE"
});
var defaultLayers = platform.createDefaultLayers();

// display the map in html with centered with default latitude and longitude 
var map = new H.Map(document.getElementById('map'),
    defaultLayers.vector.normal.map, {
    center: { lat: 0, lng: 0 },
    zoom: 2,
    pixelRatio: window.devicePixelRatio || 1
});

//Add a resize listener
window.addEventListener('resize', () => map.getViewPort().resize());

//make the map interactive
var behavior = new H.mapevents.Behavior(new H.mapevents.MapEvents(map));

// Create the default UI components
var ui = H.ui.UI.createDefault(map, defaultLayers);

// Read locations and add markers
window.onload = function () {
    // let markers = []; // Create an empty array to store the markers
    let lineString = new H.geo.LineString();
    let locations = document.getElementsByClassName('cities');
    for (let i = 0; i < locations.length; i++) { // loop through locations
        console.log(locations)
        
        fetch('https://geocoder.ls.hereapi.com/6.2/geocode.json?searchtext=' + locations[i].textContent + '&gen=9&apiKey=chDMQRlDRv0KngLYo3sXcOtNBESgx1eEU199e4Z1B7U')
            .then(response => response.json())
            .then(data => {
                let coordinates = { // add marker
                    lat: data.Response.View[0].Result[0].Location.DisplayPosition.Latitude,
                    lng: data.Response.View[0].Result[0].Location.DisplayPosition.Longitude
                }
                let marker = new H.map.Marker({ lat:coordinates.lat, lng:coordinates.lng });
                // marker.setSize(new H.map.Size(20, 20));
                lineString.pushPoint({ lat:coordinates.lat, lng:coordinates.lng });
                map.addObject(marker);
            });
    };
    map.addEventListener('mapviewchangeend', function () {
        // Add the polyline to the map
        map.addObject(new H.map.Polyline(
            lineString, { style: { lineWidth: 4, }}
        ));
    });    
};