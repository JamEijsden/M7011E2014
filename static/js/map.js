var map
function initialize(map) {
  var mapCanvas = document.getElementById('map_canvas');
  var mapOptions = {
    center: new google.maps.LatLng(44.5403, -78.5463),
    zoom: 8,
    mapTypeId: google.maps.MapTypeId.ROADMAP
  }
  map = new google.maps.Map(mapCanvas, mapOptions)
}
    google.maps.event.addDomListener(window, 'load', initialize);

    google.maps.event.addListener(map, 'click', function(event) {
 		placeMarker(event.latLng);
	});

function placeMarker(location) {
    var marker = new google.maps.Marker({
                position:location,
                map: map,
                draggable:true,
                animation: google.maps.Animation.DROP,
                title:"Stair!",
              icon: "http://maps.google.com/mapfiles/ms/micons/blue.png"
            });
}