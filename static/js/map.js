var map

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
      zoom: 4,
      center: myLatlng,
      mapTypeId: google.maps.MapTypeId.ROADMAP
    }
    map = new google.maps.Map(mapCanvas, mapOptions);
  
    createAllMarkers(locations);
    
    

    google.maps.event.addListener(map, 'click', function(event) {
      createMarker(event.latLng);
    });
}

function createAllMarkers(locations){
  for (var i=0; i < locations.length; i++) {
      console.log(locations[i]);
      pos = locations[i].position.replace('(','').replace(')','');
      var bits = pos.split(/,\s*/);
      console.log(bits[0]+", "+bits[1]);
      latLng = CreateLatLngObject(bits[0], bits[1]);
      locations[i].position = latLng;
      getLocations(locations[i], 1, locations[i].id);
    }
}