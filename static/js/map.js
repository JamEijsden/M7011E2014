var map
function initialize() {
  var mapCanvas = document.getElementById('map_canvas');
  var mapOptions = {
    center: new google.maps.LatLng(44.5403, -78.5463),
    zoom: 8,
    mapTypeId: google.maps.MapTypeId.ROADMAP
  }
  if (navigator.geolocation) {
     navigator.geolocation.getCurrentPosition(function (position) {
         initialLocation = new google.maps.LatLng(position.coords.latitude, position.coords.longitude);
         map.setCenter(initialLocation);
     });
 }
  map = new google.maps.Map(mapCanvas, mapOptions)

  google.maps.event.addListener(map, 'click', function(event) {
    placeMarker(event.latLng);
  });
}
    google.maps.event.addDomListener(window, 'load', initialize);

    

