var currentMark
var loc;
function createMarker(location) {
  var marker = new google.maps.Marker({
      position: location.position,
      map: map,
      title: location.stairname,
      description: location.description,
      photo: location.photo
  });   
  console.log("Created: " + marker.position);
  loc = location;
  var contentString;
  console.log(loc);
  var infowindow = new google.maps.InfoWindow({
    content: contentString
  });
  
  google.maps.event.addListener(marker, 'click', function () {
          document.getElementById('myModalLabel').innerHTML = this.getTitle();
          document.getElementById('stairPhoto').src = this.photo;
          document.getElementById('stairDesc').innerHTML = this.description;
          $('#modal2').modal('show');

          /*currentMark = marker;
          map.setCenter(currentMark.position)
          infowindow.open(map, marker);
*/

});
google.maps.event.addListener(infowindow,'closeclick',function(){
   map.setCenter(currentMark.position) 
   //removes the marker
   // then, remove the infowindows name from the array
});

}
