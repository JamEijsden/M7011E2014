
function initializeAdd() {
  $('#add_map_canvas').empty();
  var mapCanvas = document.getElementById('add_map_canvas');
  var myLatlng = new google.maps.LatLng(-25.363882,131.044922);
  pos = document.getElementById('geoloc').value.replace('(','').replace(')','');
  var bits = pos.split(/,\s*/);
  latLng = CreateLatLngObject(bits[0], bits[1]);
  var mapOptions = {
    zoom: 17,
    center: latLng,
    mapTypeId: google.maps.MapTypeId.ROADMAP
  }
  map = new google.maps.Map(mapCanvas, mapOptions);
  

  google.maps.event.addListener(map, 'click', function(event) {
    console.log(event.latLng);
    placeMarker(event.latLng);

  });

  $('#addStairsModal').on('shown.bs.modal', function() {
  var currentCenter = map.getCenter();  // Get current center before resizing
  google.maps.event.trigger(map, "resize");
  map.setCenter(currentCenter); // Re-set previous center
  });
}


//google.maps.event.addDomListener(window, 'load', initialize);

    

