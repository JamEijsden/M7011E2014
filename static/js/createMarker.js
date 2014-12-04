var currentMark

function createMarker(location) {
  var marker = new google.maps.Marker({
      position: location,
      map: map,
      title: ""+location
  });   
  console.log("Created: " + marker.position);
 
  contentString = "<h4>"+marker.title+"</h4>"+
  "<img src='http://media.bymk.se/2014/06/katten_jarvis.jpg' alt='Smileyface' height='300' width='400'><input type='button' class='btn btn-primary' onclick='getJson();' value=Add Location'>";

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
