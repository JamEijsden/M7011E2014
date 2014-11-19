var markerPlaced = 0



function updateDrag(location) {
  var input = $('#loc');
  input.val(location);

}

function placeMarker(location) {
  if(markerPlaced == 0){
    var marker = new google.maps.Marker({
                position:location,
                map: map,
                draggable:true,
                animation: google.maps.Animation.DROP,
                title:"HEJ",
              icon: "http://maps.google.com/mapfiles/ms/micons/blue.png"
            });
    var input = $('#loc');
    input.val(location);
  }
    
    markerPlaced = 1
    google.maps.event.addListener(marker, 'dragend', function(event){
    updateDrag(marker.getPosition())
    });
    
    
}