var currentMark

function createMarker(location) {
  var marker = new google.maps.Marker({
      position: location.position,
      map: map,
      title: location.stairname
  });   
  console.log("Created: " + marker.position);
 
  contentString = "<div class='container' style='height:100%;width:100%;margin:0;text-align:center;'>"+
  "<h4>"+marker.title+"</h4>"+
    "<div class='row' style='float:left; width:50%;'>"+
      "<img src='"+location.photo+"' alt='Smileyface' height='70%' width='80%'</div>"+
    "</div>"+
    "<div class='row' style='float:right; margin:5px; text-align:left;width:50%;'>"+
     "<p style='font-size:16px;'>"+location.description+"</p>"+
    "</div>"+
  "</div>";

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
