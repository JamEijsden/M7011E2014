;
function initialize() {
 // loadPage('map');
 var data;
 getLocations(data ,0, 0);  
}




//google.maps.event.addDomListener(window, 'load', initialize);
function CreateLatLngObject(Latitude, Longitude) {
  var latlng = new google.maps.LatLng(parseFloat(Latitude), parseFloat(Longitude));
  return latlng;
}
    
function loadMap(locations){
  var mapCanvas = document.getElementById('main_map_canvas');
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
  






    createAllMarkers(locations);
    
    
}

function createAllMarkers(locations){
  for (var i=0; i < locations.length; i++) {
      pos = locations[i].position.replace('(','').replace(')','');
      var bits = pos.split(/,\s*/);
      latLng = CreateLatLngObject(bits[0], bits[1]);
      locations[i].position = latLng;
      createMarker(locations[i]);
    }
}