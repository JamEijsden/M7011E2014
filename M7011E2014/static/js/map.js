var map


 

function initialize() {
  var mapCanvas = document.getElementById('map_canvas');
  var myLatlng = new google.maps.LatLng(-25.363882,131.044922);
  var mapOptions = {
    zoom: 4,
    center: myLatlng,
    mapTypeId: google.maps.MapTypeId.ROADMAP
  }
  map = new google.maps.Map(mapCanvas, mapOptions);
  createMarker(myLatlng);
  

  google.maps.event.addListener(map, 'click', function(event) {
    createMarker(event.latLng);
  });

}


google.maps.event.addDomListener(window, 'load', initialize);

    

