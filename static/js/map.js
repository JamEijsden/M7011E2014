var map
var currentMark

 

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

function createMarker(location) {
  var marker = new google.maps.Marker({
      position: location,
      map: map,
      title: ""+location
  });   
 
  contentString = "<h4>"+marker.title+"</h4><img src='http://media.bymk.se/2014/06/katten_jarvis.jpg' alt='Smiley face' height='300' width='400'>";
  var infowindow = new google.maps.InfoWindow({
    content: contentString
  });

  google.maps.event.addListener(marker, 'click', function () {
          currentMark = marker;
          map.setCenter(currentMark.position)
          infowindow.open(map, marker);

});
google.maps.event.addListener(infowindow,'closeclick',function(){
   map.setCenter(currentMark.position) 
   //removes the marker
   // then, remove the infowindows name from the array
});

}
google.maps.event.addDomListener(window, 'load', initialize);

    

