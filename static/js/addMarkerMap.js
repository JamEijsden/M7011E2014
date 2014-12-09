
function initializeAdd() {
  $('#add_map_canvas').empty();
  var mapCanvas = document.getElementById('add_map_canvas');
  var myLatlng = new google.maps.LatLng(-25.363882,131.044922);
  var mapOptions = {
    zoom: 17,
    center: myLatlng,
    mapTypeId: google.maps.MapTypeId.ROADMAP
  }
  map = new google.maps.Map(mapCanvas, mapOptions);


     if(navigator.geolocation) {
    browserSupportFlag = true;
    navigator.geolocation.getCurrentPosition(function(position) {
      initialLocation = new google.maps.LatLng(position.coords.latitude,position.coords.longitude);
      map.setCenter(initialLocation);
    }, function() {
      handleNoGeolocation(browserSupportFlag);
    });
  }
  // Browser doesn't support Geolocation
  else {
    browserSupportFlag = false;
    handleNoGeolocation(browserSupportFlag);
  }

  function handleNoGeolocation(errorFlag) {
    if (errorFlag == true) {
      alert("Geolocation service failed.");
      initialLocation = newyork;
    } else {
      alert("Your browser doesn't support geolocation. We've placed you in Siberia.");
      initialLocation = siberia;
    }
    map.setCenter(initialLocation);
  }

  

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

    

