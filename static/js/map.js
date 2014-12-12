
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
    var myLatlng = new google.maps.LatLng(65.61845752939256, 22.143003344535828);
    var mapOptions = {
      zoom: 15,
      mapTypeId: google.maps.MapTypeId.ROADMAP
    }
    map = new google.maps.Map(mapCanvas, mapOptions);
    getGeoLoc();
    createAllMarkers(locations);
    google.maps.event.addListener(map, 'center_changed', function() {
    //document.getElementById('geoloc').value = map.getCenter();
  });

}

function createAllMarkers(locations){
  if(locations == null){
    console.log('No locations to load');
    return;
  }
  for (var i=0; i < locations.length; i++) {
      pos = locations[i].position.replace('(','').replace(')','');
      var bits = pos.split(/,\s*/);
      latLng = CreateLatLngObject(bits[0], bits[1]);
      locations[i].position = latLng;
      createMarker(locations[i]);
    }
}

function getGeoLoc(){

    if(navigator.geolocation) {
    browserSupportFlag = true;
    navigator.geolocation.getCurrentPosition(function(position) {
      initialLocation = new google.maps.LatLng(position.coords.latitude,position.coords.longitude);
      document.getElementById('geoloc').value = initialLocation;
      map.setCenter(initialLocation);
      console.log(document.getElementById('geoloc').value);
    }, function() {
      handleNoGeolocation(browserSupportFlag);
    });
  }
  // Browser doesn't support Geolocation
  else {
    browserSupportFlag = false;
    handleNoGeolocation(browserSupportFlag);
  }
}

  function handleNoGeolocation(errorFlag) {
    initialLocation = new google.maps.LatLng(65.61845752939256, 22.143003344535828);
    map.setCenter(initialLocation);
    document.getElementById('geoloc').value = initialLocation;
    
}
 

