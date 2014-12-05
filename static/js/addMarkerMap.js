var map

function initializeAdd() {
  var mapCanvas = document.getElementById('add_map_canvas');
  var myLatlng = new google.maps.LatLng(-25.363882,131.044922);
  var mapOptions = {
    zoom: 4,
    center: myLatlng,
    mapTypeId: google.maps.MapTypeId.ROADMAP
  }
  map = new google.maps.Map(mapCanvas, mapOptions);

  

  google.maps.event.addListener(map, 'click', function(event) {
    placeMarker(event.latLng);
  });

  $('#addStairsModal').on('shown.bs.modal', function() {
  var currentCenter = map.getCenter();  // Get current center before resizing
  google.maps.event.trigger(map, "resize");
  map.setCenter(currentCenter); // Re-set previous center
  });
}


//google.maps.event.addDomListener(window, 'load', initialize);

    

